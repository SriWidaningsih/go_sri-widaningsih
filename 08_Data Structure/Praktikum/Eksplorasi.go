package main

import (
	"fmt"
	"math"
)

func main() {
	var JumlahDiagonal1, JumlahDiagonal2 float64
	MatriksPersegi := [3][3]float64{
		{1, 1, 1},
		{3, 3, 3},
		{5, 5, 7},
	}
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			if a == b {
				JumlahDiagonal1 += MatriksPersegi[a][b]
			}
			if a+b == 2 {
				JumlahDiagonal2 += MatriksPersegi[a][b]
			}
		}
	}
	fmt.Println("Jumlah diagonal kiri ke kanan: ", JumlahDiagonal1)
	fmt.Println("Jumlah diagonal kanan ke kiri: ", JumlahDiagonal2)

	different := math.Abs(JumlahDiagonal1 - JumlahDiagonal2)

	fmt.Println("Maka selisih absolut antara jumlah diagonal: ", different)
}
