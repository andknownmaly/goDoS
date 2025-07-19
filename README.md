# Simple Go DoS Tester

## Deskripsi

**Simple Go DoS Tester** adalah program sederhana berbasis Go untuk melakukan uji ketahanan HTTP pada sebuah target URL.
Program ini mengirimkan permintaan HTTP GET secara terus-menerus menggunakan multi-threading (goroutine), kemudian menampilkan statistik jumlah request yang berhasil (`Sukses`) dan gagal (`Gagal`).

⚠️ **Catatan Penting:**
Program ini ditujukan untuk **pengujian yang legal dan etis saja**. Jangan gunakan untuk menyerang sistem tanpa izin, karena itu melanggar hukum.

## Fitur

* Permintaan HTTP GET paralel menggunakan goroutine
* Monitoring real-time jumlah permintaan sukses dan gagal
* Konfigurasi jumlah thread (goroutine) saat runtime

## Cara Penggunaan

### Prasyarat

* Go 1.20 atau versi lebih baru

### Build

```bash
go build -o dos-tester goDoS.go
```

### Jalankan Program

```bash
./dos-tester
```

### Input

* **Target URL** – URL target yang ingin diuji (contoh: `https://example.com`)
* **Jumlah Threads** – Berapa banyak goroutine yang digunakan untuk melakukan request secara paralel

### Contoh Output

```
Target URL: https://example.com
Jumlah Threads: 100
Memulai DoS Tester ke https://example.com dengan 100 threads
[Monitor] Sukses: 500 | Gagal: 30
[Monitor] Sukses: 1200 | Gagal: 85
```

## Struktur Kode

* `attack()` – Fungsi yang mengirim permintaan GET ke target dan menghitung keberhasilan/gagal
* `monitor()` – Fungsi yang menampilkan statistik setiap 5 detik
* `main()` – Fungsi utama untuk membaca input, memulai thread, dan memanggil monitor

## Disclaimer

Program ini hanya untuk **educational purpose dan pengujian infrastruktur milik sendiri atau yang telah diberikan izin eksplisit**.

Segala bentuk penyalahgunaan adalah tanggung jawab pengguna sepenuhnya.
