package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	Penghitung := make(map[rune]int)
	for _, keadaan := range angka {
		Penghitung[keadaan]++
	}

	var Hasilnya []int
	for keadaan, perhitungan := range Penghitung {
		if perhitungan == 1 {
			number, hasil := strconv.Atoi(string(keadaan))
			if hasil == nil {
				Hasilnya = append(Hasilnya, number)
			}
		}
	}
	return Hasilnya

}

func main() {
	fmt.Println("Berikut angka-angka yang hanya muncul satu kali :")
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
	fmt.Println(munculSekali("29012002"))
}
