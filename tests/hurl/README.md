# Hurl E2E — API Gateway (localhost:5000, GraphQL `/query`)

Suite e2e untuk `service/apigateway` (GraphQL :5000, endpoint `/query`).
Semua file memakai GraphQL mutation/query dengan pola response
`{ status, message, data }` ala API gateway ini. Register auto-verify di
environment non-production (`APP_ENV != production/kubernetes`), sehingga
register → login langsung berhasil.

## Prasyarat

- Stack berjalan: infra compose + service lokal (lihat `run_e2e.sh`).
- Database sudah di-migrate dan role `ROLE_ADMIN` tersedia (registerUser
  butuh role ini; role seeder tidak memuatnya — runner meng-insert-nya).
- Email test memakai `{{uuid}}` sehingga aman dijalankan berulang (data
  unik per run).

## Menjalankan

Semua file memakai variabel `baseUrl` dan `uuid`:

```bash
hurl --variable baseUrl=http://localhost:5000 --variable uuid=$(uuidgen) --test tests/hurl/
```

Satu file saja:

```bash
hurl --variable baseUrl=http://localhost:5000 --variable uuid=$(uuidgen) --test tests/hurl/auth.hurl
```

## Isi suite

| File | Cakupan |
|---|---|
| `auth.hurl` | register → login → me → 401 → refresh-token → forgot-password → reset-password (invalid) → verify-code (invalid) |
| `user.hurl` | lifecycle user + active/trashed + restore-all/permanent-all + invalid id |
| `role.hurl` | lifecycle role + active/trashed/find-by-user + restore-all/permanent-all |
| `card.hurl` | lifecycle card + dashboard + semua stats balance/topup/withdraw/transaction/transfer |
| `merchant.hurl` | lifecycle merchant + stats + query by api_key/user_id + restore-all/permanent-all |
| `merchantdocument.hurl` | lifecycle merchant document + update-status + restore-all/permanent-all |
| `saldo.hurl` | lifecycle saldo + stats balance |
| `topup.hurl` | card + saldo → lifecycle topup + restore-all/permanent-all |
| `transaction.hurl` | payer + merchant (card & saldo) → lifecycle transaction + queries |
| `transfer.hurl` | sender + receiver (card & saldo) → lifecycle transfer + queries |
| `withdraw.hurl` | card + saldo → lifecycle withdraw + restore-all/permanent-all |
| `stats.hurl` | semua query stats per resource (menggunakan data yang dibuat file ini) |

Semua file **self-contained** (register user sendiri di awal), sehingga bisa
dijalankan independen.

## E2E otomatis (docker compose infra + service lokal)

Jalankan langsung:

```bash
bash tests/hurl/run_e2e.sh
```

Alur:

1. Infra up via `docker compose -f deployments/local/docker-compose.yml`
   (postgres, pgbouncer, kafka, redis per-service — service Go TIDAK
   di-container, dijalankan lokal dengan `deployments/local/local.env`).
2. Tunggu postgres/redis/kafka → reset DB → migrate → insert `ROLE_ADMIN`.
3. Build service lokal → start service lokal (`APP_ENV=development`).
4. Jalankan setiap `*.hurl` satu per satu (dengan jeda antar file agar tidak
   kena rate limiter gateway).

## Catatan

- Financial domain (topup, withdraw, transfer, transaction) memakai SQL
  atomik yang butuh saldo row di card terkait — file terkait membuat
  card + saldo sendiri sebelum operasi.
- Transaction: card settlement merchant harus beda dari card payer (aturan
  SQL `$1 <> $3`) — file memakai dua user berbeda.
- Status negatif: GraphQL error dibawa lewat `$.errors` (HTTP tetap 200)
  kecuali auth middleware yang mengembalikan `401`.
