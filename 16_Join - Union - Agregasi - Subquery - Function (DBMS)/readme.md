## Sri Widaningsih


#### Resume Join - Union - Agregasi - Subquery - Function (DBMS)

1. Join
- Join adalah operasi yang digunakan untuk menggabungkan dua atau lebih tabel dalam database berdasarkan kunci primernya. Join biasanya digunakan untuk mengambil data dari dua tabel atau lebih dan menggabungkannya menjadi satu hasil. 
- Contohnya

``` js
SELECT * FROM table1 JOIN table2 ON table1.key = table2.key;
```

2. Union
- Union adalah operasi yang digunakan untuk menggabungkan hasil dari dua atau lebih query menjadi satu hasil. Union hanya dapat digunakan jika kedua query menghasilkan hasil dengan jumlah kolom yang sama dan tipe data yang sama. 
- Contohnya

``` js
SELECT col1 FROM table1
UNION
SELECT col1 FROM table2;
```

3. Agregasi
- Agregasi adalah operasi yang digunakan untuk menghitung statistik pada data yang disimpan dalam tabel. Agregasi biasanya digunakan untuk menghitung jumlah, rata-rata, nilai maksimum, atau nilai minimum dari kolom tertentu dalam tabel.
- Contohnya

``` js
SELECT COUNT(*) FROM table1;
SELECT AVG(col1) FROM table1;
SELECT MAX(col1) FROM table1;
```

4. Subquery
- Subquery adalah query yang disisipkan dalam query lain untuk mengambil data dari tabel. Subquery biasanya digunakan untuk membatasi hasil query atau digunakan sebagai kriteria pencarian dalam WHERE clause.
- contohnya
``` js
SELECT col1 FROM table1 WHERE col1 IN (SELECT col1 FROM table2);
);
```


5. Function
- Function adalah blok kode yang dapat digunakan untuk melakukan operasi pada data yang disimpan dalam tabel. Function biasanya digunakan untuk melakukan operasi kompleks pada data, seperti pengolahan teks atau konversi tipe data.

- contohnya
``` js
SELECT UPPER(col1) FROM table1;
SELECT DATE_FORMAT(col1, '%Y-%m-%d') FROM table1;
```
#### Untuk Tugas sudah terlampir pada folder screenshot