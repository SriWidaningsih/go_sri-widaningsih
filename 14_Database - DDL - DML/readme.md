## Sri Widaningsih


#### Resume Database - DDL - DML
Database adalah kumpulan data yang terorganisir secara sistematis dan dapat diakses, dikelola, dan diperbarui dengan mudah. Database digunakan untuk menyimpan informasi yang terstruktur dalam format tabel, yang terdiri dari baris dan kolom yang terhubung satu sama lain dengan cara tertentu.
1. One to one (1:1)
- Dalam hubungan one to one, satu baris dalam tabel pertama berkorespondensi dengan satu baris dalam tabel kedua.
- Contohnya adalah hubungan antara tabel "karyawan" dan "alamat karyawan", di mana setiap karyawan hanya memiliki satu alamat.

Tabel Karyawan
 
| ID        | Nama        | Umur |
| -------   | ----------- | ---- |
| 1         |  Wida       |  21  |
| 2         |  Gilang     |  21  |


Tabel Alamat 
 
| ID        | Alam        | Kota |
| -------   | ----------- | ---- |
| 1         |  Jl. BudiAsih       |  Tangerang  |
| 2         |  Gang Leci     |  Bali  |

2. One to many (1:N)
- Dalam hubungan one to many, satu baris dalam tabel pertama berkorespondensi dengan banyak baris dalam tabel kedua.
- Contohnya adalah hubungan antara tabel "departemen" dan "karyawan", di mana setiap departemen memiliki banyak karyawan.

Tabel Departemen
 
| ID        | Bidang        |
| -------   | ----------- |
| 1         |  IT       | 
| 2         |  Keuangan     |


Tabel Karyawan 
 
| ID        | Nama        | Umur | Departemen ID |
| -------   | ----------- | ---- | ------------- |
| 1         |  Wida       |  21  | 1 |
| 2         |  Gilang     |  21  | 1|

3. many to many (N:M)
- Dalam hubungan many to many, banyak baris dalam tabel pertama berkorespondensi dengan banyak baris dalam tabel kedua.
- Contohnya adalah hubungan antara tabel "mahasiswa" dan "mata kuliah", di mana setiap mahasiswa dapat mengambil banyak mata kuliah, dan setiap mata kuliah dapat diambil oleh banyak mahasiswa.

Tabel Mahasiswa
 
| ID        | Nama        | Umur | Jurusan |
| -------   | ----------- | ---- | ------------- |
| 1         |  Wida       |  21  | Fisika |
| 2         |  Gilang     |  21  | Manajemen|


Tabel Mata Kuliah 
 
| ID        | Nama Mata Kuliah        | Dosen Pengajar | 
| -------   | ----------- | ---- |
| 1         |  SPAI       |  Dr. Budi  | 
| 2         |  MRPI     |  Prof. Hasan  | 
4. DDL
- DDL (Data Definition Language) adalah bahasa pemrograman yang digunakan untuk mendefinisikan, mengubah, dan menghapus struktur database, seperti tabel, indeks, kunci, dan prosedur penyimpanan.
- contohnya
``` js
CREATE TABLE karyawan (
id INT PRIMARY KEY,
nama VARCHAR(50),
umur INT
);

ALTER TABLE karyawan
ADD COLUMN alamat VARCHAR(100);

DROP TABLE karyawan;

TRUNCATE TABLE karyawan;

RENAME TABLE karyawan TO pegawai;
```


5. DML
- DML (Data Manipulation Language) adalah bahasa pemrograman yang digunakan untuk memanipulasi data dalam database, seperti menambah, mengubah, dan menghapus data dalam tabel. 

- contohnya
``` js
INSERT INTO mahasiswa (id, nama, umur) VALUES (1, 'Andi', 20);

UPDATE mahasiswa SET umur = 21 WHERE id = 1;

DELETE FROM mahasiswa WHERE id = 1;
```
#### Untuk Tugas sudah terlampir pada folder screenshot