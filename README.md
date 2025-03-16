# Dokumentasi Tugas TEFA Golang Ke-5

## Setup Proyek
```
go get github.com/dgrijalva/jwt-go@v3.2.0
go get github.com/gin-gonic/gin@v1.10.0
go mod tidy
```

## Menjalankan Aplikasi
```
go run cmd/main.go
```

## Link URL untuk register (membuat akun)
```http://localhost:8080/register```
Setelah sukes membuat akun halaman register akan redirect ke halaman login.

## Link URL untuk login (masuk akun)
```http://localhost:8080/login```
Jika belum buat akun atau mencoba masuk dengan akun yang non existent halaman login akan reset form atau menampilkan `Invalid Credentials`.
Jika akun ada dan password, nama, dan email benar maka halaman login akan redirect ke halaman profile

## Link URL melihat profile
```http://localhost:8080/profile```
Jika sukses dalam login, halaman profile akan menampilkan nama yang anda gunakan dalam login dan tanggal akun anda dibuat.
Halaman profile akan redirect ke halaman login jika belum login.
