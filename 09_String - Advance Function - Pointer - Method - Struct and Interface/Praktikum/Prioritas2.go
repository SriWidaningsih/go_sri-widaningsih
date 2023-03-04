package main

import "fmt"

func caesar(offset int, input string) string {
	var Hasil string
	for _, karakter := range input {
		if karakter == ' ' {
			Hasil += " "
			continue
		}
		kodeKarakter := int(karakter)
		KodeTerbaru := (kodeKarakter-97+offset)%26 + 97
		KarakterBaru := string(KodeTerbaru)
		Hasil += KarakterBaru
	}
	return Hasil
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1, "wida"))
	fmt.Println(caesar(25, "xjeb"))
}
