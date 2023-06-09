package main

import "fmt"

func Mapping(slice []string) map[string]int {
	HasilAkhir := make(map[string]int)
	for _, slices := range slice {
		HasilAkhir[slices]++

	}
	return HasilAkhir

}

func main() {
	fmt.Println("Maka berikut kata beserta jumlah perulangan nya :")
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
	fmt.Println(Mapping([]string{"taehyung", "jungkook", "suga", "taehyung", "taehyung"}))

}
