package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	min = *numbers[0]
	max = *numbers[0]
	for _, angka := range numbers {
		if *angka > max {
			max = *angka
		}
		if *angka < min {
			min = *angka
		}
	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai Minimum pada susunan angka tersebut:", min)
	fmt.Println("Nilai Maximum pada susunan angka tersebut:", max)
}
