## Sri Widaningsih


#### Resume Intro Echo Golang

1. ECHO
- Echo adalah sebuah framework web untuk bahasa pemrograman Go (Golang) yang dirancang untuk mempermudah pengembangan aplikasi web dan RESTful API. Echo menyediakan fitur-fitur seperti middleware, routing, penggunaan template, validasi input, dan lainnya untuk mempercepat pengembangan aplikasi web.

2. Download Echo
```js
go get github.com/labstack/echo/v4
```

- sehingga dapat digunakan sebagai berikut :
```js
package main

import (
    "github.com/labstack/echo/v4"
    "net/http"
)

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    e.Logger.Fatal(e.Start(":8080"))
}

```
- disini akan dihubungkan dengan postman memasukan nilai
```js
localhost:8080/
```
- sehingga nantinya akan muncul output berupa
```js
Hello, World!
```

3. Data rendering
- Dalam framework Echo pada bahasa pemrograman Go, Echo menyediakan beberapa renderer yang dapat digunakan untuk melakukan rendering data, seperti HTML template renderer, JSON renderer, XML renderer, dan Plain text renderer. Renderer dapat diatur dalam file konfigurasi dan dihubungkan dengan route tertentu untuk melakukan rendering data sesuai dengan format yang diinginkan.

4. Retrieve data
5. Binding Data

#### Untuk Tugas sudah terlampir pada folder screenshot dan script yang digunakan ada pada folder praktikum