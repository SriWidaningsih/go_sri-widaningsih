## Sri Widaningsih


#### Resume Configuration Management and CLI

1. Configuration Management
- Configuration Management (Manajemen Konfigurasi) adalah praktik untuk mengelola dan mengontrol versi dari konfigurasi perangkat lunak atau sistem. Ini melibatkan penyimpanan dan manajemen semua elemen yang diperlukan untuk menjalankan perangkat lunak atau sistem, termasuk kode sumber, konfigurasi, data, dan dokumentasi.
- Contoh penerapan Configuration Management adalah menggunakan Git untuk mengelola kode sumber sebuah proyek.

``` js
# Inisialisasi direktori proyek sebagai repository Git
git init

# Menambahkan file ke staging area
git add <file>

# Membuat commit dengan pesan deskripsi perubahan
git commit -m "deskripsi perubahan"

# Push commit ke repository remote
git push origin master

```

2. Command Line Interface (CLI)
- Command Line Interface (CLI) adalah antarmuka pengguna yang memungkinkan pengguna untuk berinteraksi dengan komputer melalui baris perintah yang ditulis secara manual. CLI sering digunakan oleh pengembang dan administrator sistem untuk melakukan tugas-tugas yang berulang atau kompleks secara otomatis.
- Contohnya

``` js
User@DESKTOP-1SD571G MINGW64 ~/Downloads/tugasfix
$ ls
 automate*  'wida 2023-03-26'/


```

3. Shell Unix
- Shell Unix adalah program antarmuka pengguna (user interface) pada sistem operasi Unix, yang menyediakan cara untuk pengguna berinteraksi dengan sistem operasi melalui baris perintah (command-line interface).
- Contohnya

``` js
# membuat folder
User@DESKTOP-1SD571G MINGW64 ~/Downloads
$ mkdir Contoh

# pindah folder
User@DESKTOP-1SD571G MINGW64 ~/Downloads
$ cd Contoh

User@DESKTOP-1SD571G MINGW64 ~/Downloads/Contoh
$

# membuat file
User@DESKTOP-1SD571G MINGW64 ~/Downloads/Contoh
$ touch contoh_file.txt

# periksa koneksi
User@DESKTOP-1SD571G MINGW64 ~/Downloads/Contoh
$ ping google.com

Pinging forcesafesearch.google.com [2001:4860:4802:32::78] with 32 bytes of data:
Reply from 2001:4860:4802:32::78: time=21ms
Reply from 2001:4860:4802:32::78: time=23ms
Reply from 2001:4860:4802:32::78: time=23ms
Reply from 2001:4860:4802:32::78: time=22ms

Ping statistics for 2001:4860:4802:32::78:
    Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
    Minimum = 21ms, Maximum = 23ms, Average = 22ms

# mencari teks dalam file
User@DESKTOP-1SD571G MINGW64 ~/Downloads/Contoh
$ grep "wida" contoh_file.txt
Halo nama saya wida!!!

```


#### Untuk Tugas sudah terlampir pada folder screenshot dan script automate yang digunakan ada pada folder praktikum