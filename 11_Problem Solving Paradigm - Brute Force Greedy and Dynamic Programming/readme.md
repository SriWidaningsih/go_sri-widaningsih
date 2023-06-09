## Sri Widaningsih


#### Resume Problem Solving Paradigm - Brute Force Greedy and Dynamic Programming
Problem solving paradigm adalah pendekatan yang biasa digunakan untuk memecahkan masalah: Complete Search (alias Brute Force). Divide and Conquer, pendekatan Greedy, dan Dynamic Programming. Setiap masalah harus diselesaikan dengan pendekatan yang sesuai.
1. Complete Search
- Complete Search (juga dikenal sebagai) Bruteforce adalah metode untuk menyelesaikan masalah dengan menelusuri seluruh ruang pencarian untuk mendapatkan solusi yang dibutuhkan.
- Contoh Find Max and Min

``` 
Diberikan array A yang berisi n≤ 10.000
PENYELESAIANNYA :
findMaxMin([10, 7, 3, 5, 8, 2, 9]) // 10 2
```


2. Divide and Conquer
- Divide & Conquer (D&C) adalah paradigma pemecahan masalah di mana sebuah masalah dibuat lebih sederhana dengan 'membaginya' menjadi bagian-bagian yang lebih kecil dan kemudian mengambil alih setiap bagian.

| Nama | Deskripsi |
| ------ | ----------- |
| Divide   | membagi masalah yang besar menjadi masalah yang lebih kecil |
| Conquer | ketika masalah sudah cukup kecil untuk diselesaikan, langsung selesaikan |
| combine    | jika dibutuhkan maka perlu menggabungkan solusi dan masalah-masalah yang lebih kecil menjadi solusi untuk masalah yang besar |

- Contoh Binary Search


``` 
binarySearch([1, 1, 3, 5, 5, 6, 7], 3) // 2 
binarySearch([12, 15, 15, 19, 24, 31, 53, 59, 60], 53) // 6 
binarySearch([12, 15, 15, 19, 24, 31, 53, 59, 60], 180) //-1
```

3. Greedy
- Sebuah algoritma dikatakan greedy jika ia membuat pilihan yang optimal secara lokal pada setiap langkah dengan harapan pada akhirnya mencapai solusi yang optimal secara global. Dalam beberapa kasus, greedy bekerja dengan baik - solusinya singkat dan berjalan dengan efisien.
- Contoh digunakan pada kasus berikut

``` 
Diberikan jumlah target V sen dan daftar pecahan n koin,
misalnya kita memiliki coinValue[i] (dalam sen) untuk
jenis koin i = [0..n-1], berapakah jumlah minimum koin yang
harus kita gunakan untuk merepresentasikan jumlah V?
Asumsikan bahwa kita memiliki persediaan koin yang tidak
terbatas untuk semua jenis koin. coinValue = {10, 25, 5. 1 }
```
```
CoinChange(42) // 25 18 5 11
```
4. Dynamic Programming

#### Untuk Tugas sudah terlampir pada folder Praktikum, dan juga folder screenshot