package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	Beda := make(map[string]bool)

	HasilAkhir := []string{}

	for _, nilainya := range arrayA {

		if _, fix := Beda[nilainya]; !fix {
			Beda[nilainya] = true
			HasilAkhir = append(HasilAkhir, nilainya)
		}
	}

	for _, nilainya := range arrayB {

		if _, ok := Beda[nilainya]; !ok {
			Beda[nilainya] = true
			HasilAkhir = append(HasilAkhir, nilainya)
		}
	}

	return HasilAkhir

}
func main() {
	fmt.Println("Kumpulan nama-nama pada array tersebut (tidak ada pengulangan) adalah:")

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))

}
