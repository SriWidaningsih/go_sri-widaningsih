## Sri Widaningsih


#### Resume Intro RESTful API

1. API (Application Programming Interface)
- API adalah antarmuka yang memungkinkan berbagai aplikasi dan sistem berkomunikasi dan berinteraksi satu sama lain.

2. REST (Representational State Transfer)
- REST adalah sebuah arsitektur yang digunakan untuk mengembangkan layanan web. RESTful API menggunakan prinsip HTTP untuk memungkinkan komunikasi antara client dan server. RESTful API menggunakan HTTP verbs seperti GET, POST, PUT, dan DELETE untuk membaca, membuat, mengubah, dan menghapus data.

3. JSON (JavaScript Object Notation)
- JSON adalah format pertukaran data ringan yang sering digunakan dalam pengembangan web. JSON mudah dibaca dan ditulis oleh manusia dan mudah diuraikan dan dibangun oleh komputer. JSON digunakan untuk mengirim dan menerima data antara server dan client dalam RESTful API.

4. HTTP Response Code
- HTTP Response Code adalah status yang dikirim oleh server sebagai respons terhadap permintaan client. Beberapa kode status HTTP yang umum digunakan antara lain:
```js
200 OK                   : permintaan berhasil
201 Created              : permintaan berhasil dan sumber daya baru telah dibuat
400 Bad Request          : permintaan tidak valid
401 Unauthorized         : akses tidak diizinkan
404 Not Found            : sumber daya yang diminta tidak ditemukan
500 Internal Server Error: server mengalami kesalahan internal
```


5. Package net/http (golang)
Package net/http adalah paket standar Go untuk membangun server HTTP dan klien. Paket ini menyediakan fungsi dan struktur data untuk mengontrol permintaan HTTP dan respons dari server.

```js
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8080", nil)
}


```

#### Untuk Tugas sudah terlampir pada folder screenshot dan script yang digunakan ada pada folder praktikum