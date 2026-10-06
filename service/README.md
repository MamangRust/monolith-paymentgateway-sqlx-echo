## Service

### Api Gateway
``ApiGateway`` adalah layanan API yang menyediakan berbagai endpoint untuk mengakses layanan lainnya melalui protokol gRPC.

### Auth GRPC Server
``AuthService`` adalah layanan otentikasi berbasis gRPC yang menyediakan berbagai endpoint terkait proses otentikasi pengguna seperti:

- Registrasi pengguna
- Login
- Reset password
- Verifikasi kode OTP
- Refresh token
- Mendapatkan data user yang lewat GetMe


### Card GRPC Server
``CardService`` adalah layanan berbasis gRPC yang menangani seluruh operasi terkait data kartu pengguna dalam sistem, seperti pengelolaan kartu, statistik saldo dan transaksi, serta manajemen data aktif dan terhapus (soft delete).

#### Fitur Utama
1. CRUD Kartu: Buat, ubah, dan hapus kartu pengguna dengan sistem soft-delete.
2. Manajemen Data: Mendukung FindByUserId, FindByCardNumber, FindByActive, dan FindByTrashed.
3. Dashboard Analytics: Menampilkan statistik jumlah kartu aktif, terhapus, saldo, dan transaksi.
5. Analitik Transaksi: Statistik berdasarkan bulan/tahun untuk topup, withdraw, transfer (pengirim & penerima).
6. Granular Analytics by Card Number: Statistik individual berdasarkan nomor kartu.


### Email(Worker Service)
Layanan ini bertanggung jawab untuk mengirim email secara asynchronous berdasarkan event yang diterima dari berbagai topik Kafka. Ini adalah komponen worker/service yang berjalan secara terus-menerus dan mendengarkan pesan dari Kafka, lalu mengirim email menggunakan SMTP.
#### Fitur Utama
- Mendengarkan berbagai topik Kafka terkait email.
- Mengirimkan email menggunakan SMTP (saat ini menggunakan Ethereal Email untuk testing).
- Mendaftarkan endpoint /metrics menggunakan Prometheus untuk monitoring.
- Struktur handler yang modular dan dapat dikembangkan untuk berbagai jenis email.


### Merchant GRPC Server
``MerchantService`` adalah layanan gRPC yang menangani seluruh kebutuhan manajemen merchant, termasuk data dasar merchant, transaksi, statistik bulanan/tahunan, serta fitur penghapusan dan pemulihan data (soft delete & restore).


### Topup GRPC Server
``TopupService`` merupakan layanan gRPC yang menangani operasi top-up termasuk pencatatan, pelacakan, statistik, dan penghapusan data.


### Transaction GRPC Server
``TransactionService`` adalah layanan gRPC yang menangani proses pencatatan dan pengelolaan transaksi pengguna, termasuk statistik bulanan/tahunan serta fitur soft delete & restore.


### Transfer GRPC Server
``TransferService`` adalah layanan gRPC yang menangani operasi transfer termasuk pencatatan, pelacakan, statistik, dan penghapusan data.


### User GRPC Server
```UserService`` adalah layanan gRPC yang menangani operasi pengguna termasuk crud,  serta fitur penghapusan dan pemulihan data (soft delete & restore).

### Withdraw GRPC Server
``WithdrawService`` adalah layanan gRPC yang menangani operasi withdraw termasuk pencatatan, pelacakan, statistik, dan penghapusan data.
