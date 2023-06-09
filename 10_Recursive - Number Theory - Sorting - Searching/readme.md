## Sri Widaningsih


### Resume Recursive - Number Theory - Sorting – Searching
1. recursive (rekursif)
- recursion adalah keadaan di mana sebuah fungsi menyelesaikan masalah dengan memanggil dirinya sendiri berulang kali. Banyak masalah yang lebih mudah dipecahkan dan mempersingkat kode ketika menggunakan pendekatan recursive. 
- Contoh pada Bilangan mencari faktorial

``` js
public int factorial(int n) {
  if (n == 1) {
    return 1;
  } else {
  return n*factorial(n-1);
  }

}
```


2. Number theory (teori bilangan)
- Number theory adalah cabang matematika yang mempelajari bilangan bulat. Ada banyak topik dalam bidang teori bilangan seperti bilangan prima, dll.
- Contoh mencari nilai FPB


``` js
public int FPB(int a, int b) {
   if (b == 0) {
     return a; 
   }
   return gcd(b, a % b);
}
```
```
FPB dari bilangan bulat a dan b, dilambangkan FPB(a, b), adalah
bilangan bulat terbesar yang membagi a dan b.
```

3. searching (pencarian)
- Searching adalah proses menemukan posisi nilai tertentu dalam daftar nilai. 
- Contoh Linier search - O(n)


``` js
public int linierSearch(List<Integer> elements, int x) {
    for (int i = 0; i < elements.toArray().length; i++) {
        if (elements.get(i) == x) {
           return 1;
        }
    }
    return -1;
}
```
4. Sorting (Penyortiran), Sorting adalah proses menyusun data dalam urutan tertentu

#### Untuk Tugas sudah terlampir pada folder Praktikum, dan juga folder screenshot