package main

import "fmt"

func Romawi(num int) string {
	Nilai := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	Simbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var Hasil string
	for i := 0; i < len(Nilai); i++ {
		for num >= Nilai[i] {
			num -= Nilai[i]
			Hasil += Simbol[i]
		}
	}
	return Hasil
}

func main() {
	fmt.Println(Romawi(4))    // IV
	fmt.Println(Romawi(9))    // IX
	fmt.Println(Romawi(23))   // XXIII
	fmt.Println(Romawi(2021)) // MMXXI
	fmt.Println(Romawi(1646)) // MDCXLVI
}
