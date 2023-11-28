## Note
Aplikasi ini seharusnya dijalankan menggunakan Docker compose namun ada masalah ketika build aplikasi menggunakan docker compose di sisi React (Next JS). Belum sempat diperbaiki karena keterbatasan waktu. Sehingga harus dijalankan secara manual

## Folder Aplikasi
- /backend -> Aplikasi backend
- /web -> Aplikasi frontend

## Mengubah Env Backend
- Menduplikasi file env contoh
```
cp .env.example .env
```
env backend bisa diubah disesuaikan sesuai local machine

## Menjalankan Backend
- Melakukan inisiai project golang
```
go mod download
```
- Melakukan migrasi DB
```
go run main.go migration migrate
```
- Menjalankan aplikasi backend
```
go run main.go serve ---port=5000
```

## Mengubah env frontend
- Menduplikasi file env contoh
```
cp .env.example .env.local
```
env frontend bisa diubah di file .env.local disesuaikan sesuai local machine

## Menjalankan aplikasi frontend
- Melakukan instalasi dependency
```
yarn install
```
- Menjalankan Aplikasi
```
yarn dev
```
