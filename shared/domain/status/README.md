# Status Lifecycle (State Machine) & Idempotency Key

Package `shared/domain/status` adalah **single source of truth** untuk lifecycle
status record finansial (topup, transfer, withdraw, transaction) dan
dokumentasi format idempotency key. Service memanggil
`status.ValidateTransition(domain, from, to)` — transisi invalid ditolak dengan
**400 Bad Request** sebelum mutasi saldo/ledger apa pun terjadi. Jangan
mengimplementasikan cek transisi inline di service; tambahkan transisi baru di
package ini dulu.

## Aturan Umum

- Status disimpan sebagai string lowercase pendek (`VARCHAR(20)`, default
  `'pending'`).
- **Transisi same-state hanya diizinkan bila eksplisit di tabel**:
  `success → success` untuk re-assertion idempotent di flow update. Operasi
  lifecycle yang **memutasi saldo** (transaction Capture/Void/Refund) sengaja
  **tidak punya same-state entry** — panggilan ganda ditolak 400 (mencegah
  double-credit merchant, double-release hold, double-refund).
- **`failed` boleh dicapai dari status terminal** (`success`) **hanya oleh
  kompensasi sistem** (`markAsFailed` di service) — client tidak pernah bisa
  meminta status `failed` melalui tabel ini.
- Status terminal tidak dapat mundur: record yang sudah `success`/`voided`/
  `refunded` tidak bisa kembali ke `pending`; pembalikan uang dilakukan lewat
  record/flow baru (refund, void), bukan lewat flip status.
- `card_billing` didefinisikan sebagai referensi; wiring ke card service masih
  backlog.

## State Machine per Domain

### Topup / Transfer / Withdraw (pola sama)

```
            +-----------------------+
            v                       |
   pending ----> success            |
     |   \         |                |
     |    \        +--> failed -----+ (kompensasi sistem)
     v     \            |
   failed  <------------+ (retry reopen)
     |
     +----> success (retry completes)
```

| From | To (diizinkan) | Catatan |
|---|---|---|
| `pending` | `success`, `failed` | settlement selesai / gagal |
| `failed` | `pending`, `success` | retry dibuka ulang / selesai |
| `success` | `success`, `failed` | `success` = re-confirm update; `failed` = kompensasi sistem |

Tidak diizinkan: `success → pending`, dan transisi antar domain lainnya.

### Transaction

```
   pending ---> authorized ---> captured ---> refunded
      |  \          |   \          |
      |   \         |    \         +---> failed (kompensasi sistem)
      |    +--------+     +-------------> voided (dari pending/authorized)
      |    \
      +-----> success (update flow) ---> failed (kompensasi sistem)
      |
      +-----> failed (kompensasi sistem)
```

| From | To (diizinkan) | Catatan |
|---|---|---|
| `pending` | `authorized`, `voided`, `success`, `failed` | create default `pending` |
| `authorized` | `captured`, `voided`, `failed` | hold ditempatkan |
| `captured` | `refunded`, `failed` | settlement ke merchant |
| `success` | `success`, `failed` | flow update; `failed` = kompensasi |
| `failed` | `pending`, `success` | retry |

Tidak diizinkan: `captured → authorized` (tidak bisa un-capture),
`voided → *` / `refunded → *` (terminal), `authorized → refunded` (harus
capture dulu), `success → captured`.

### Card Billing (referensi, belum di-wire)

| From | To (diizinkan) |
|---|---|
| `pending` | `unpaid`, `paid`, `completed`, `failed` |
| `unpaid` | `paid`, `completed`, `failed` |
| `paid` | `completed` |
| `completed` | `failed` (dibuka ulang untuk koreksi) |
| `failed` | `unpaid`, `paid` |

## Format Idempotency Key

Keputusan A.1 (Fase A):

| Aspek | Keputusan |
|---|---|
| **Format** | `{domain}:{operation}:{owner}:{uuid}` — contoh `topup:create:user:123e4567-e89b-12d3-a456-426614174000` |
| **Panjang** | `≤ 64` karakter (validasi `validate:"omitempty,max=64"`) |
| **Storage** | Kolom DB `idempotency_key VARCHAR(64) NOT NULL DEFAULT ''` + **partial unique index** `WHERE idempotency_key <> ''` |
| **TTL** | Tidak ada — key **unik permanen** seumur hidup record (soft-delete tidak membebaskan key, mencegah double-create setelah trash) |
| **Kosong** | Key kosong = request **tidak** idempotent (perilaku lama tetap berjalan, non-breaking) |
| **Replay** | Key sama → kembalikan **response asli** (fast-path `Get*ByIdempotencyKey`); race insert → fallback `23505` → refetch record |
| **Side-effect** | Saat replay, side-effect (email, update card metadata) **tidak diulang** — ditandai `Replayed` di repository |
| **Owner scope** | `{owner}` memisahkan key antar pemilik (mis. `user:<id>` untuk topup/transfer/withdraw, `merchant:<id>` untuk transaction) agar key tak bentrok antar akun |

Migration: `20260805000000_add_idempotency_keys_to_financial_commands.sql`
(tabel `topups`, `transfers`, `withdraws`, `transactions`).

## Kontrak Error

- `ValidateTransition(domain, from, "")` → 400 "status is required".
- Transisi invalid → 400 dengan pesan: `invalid status transition for topup:
  "success" -> "pending" (allowed from "success": failed)`.
- Kompensasi sistem (`markAsFailed`) memanggil repo status langsung dan di-log
  bila gagal — ini pengecualian sistem, bukan jalur client.
