## Sri Widaningsih


#### Resume ORM and Code Structure (MVC)

1. ORM atau Object Relational Mapping
- ORM atau Object Relational Mapping adalah teknologi pemetaan objek ke dalam tabel dalam database relasional. Dalam ORM, objek dalam bahasa pemrograman seperti Java atau Python dipetakan ke dalam tabel dan kolom dalam database relasional.
- GO RM pada Golang merujuk pada Go Object Relational Mapping, yang merupakan sebuah teknologi pemetaan objek ke dalam database relasional pada bahasa pemrograman Go. GO RM memungkinkan pengembang untuk memetakan objek Go ke dalam tabel dan kolom dalam database relasional, serta melakukan operasi database seperti membuat, membaca, memperbarui, dan menghapus data dengan mudah.
- contoh
```js
type Person struct {
  gorm.Model
  ID   uint64
  Name string
  Age  int32
  Sex  string
}

var person Person
person.ID = 1
person.Name = "Alex"
person.Age = 22
person.Sex = "Male"
```

- Tabel nya

| ID     | Name    | Age  |Sex  | 
| ------ | ----------- |-----| ----|
| 1   | Wida | 21 | Female |
| 2 | Gilang| 20 | Male |
| 3 | Nesya| 10 | Female |

2. DATABASE Migration
- Database migration adalah sebuah teknik untuk mengelola evolusi struktur database secara otomatis dan terkontrol melalui kode, sehingga memudahkan pengembangan aplikasi yang membutuhkan perubahan skema database. Dengan menggunakan database migration, pengembang dapat mengubah skema database dengan aman, melakukan rollback, dan berbagi perubahan skema database dengan tim pengembang yang lain.
- untuk melakukan koneksi ke MYSQL digunaka GORM
```js
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

3. Code Structure
- Model-View-Controller (MVC)
- Model-View-Controller (MVC) adalah sebuah pola desain arsitektur perangkat lunak yang terdiri dari tiga komponen utama yaitu Model, View, dan Controller. Setiap komponen memiliki peran dan tanggung jawab yang terpisah, sehingga memudahkan dalam mengembangkan dan memelihara aplikasi.
- contohnya yaitu
```js
User@DESKTOP-1SD571G MINGW64 /c/Program Files/Go/src/project
$ tree
.
|-- config
|   `-- config.go
|-- controllers
|   |-- blogController.go
|   |-- bookController.go
|   `-- userController.go
|-- go.mod
|-- go.sum
|-- lib
|   `-- database
|       |-- blog.go
|       |-- book.go
|       `-- user.go
|-- main.go
|-- models
|   |-- blog.go
|   |-- book.go
|   `-- user.go
|-- routes
|   `-- routes.go
`-- usecase
    |-- blog.go
    |-- book.go
    `-- user.go
```


#### Untuk Tugas sudah terlampir pada folder screenshot dan script yang digunakan ada pada folder praktikum