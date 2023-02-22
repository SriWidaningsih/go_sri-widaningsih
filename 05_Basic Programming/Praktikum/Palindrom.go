package main

import (
	"fmt"
	"strings"
)

//Cek kata atau kalimat palindrom

func main() {

	var KataAwal string = "Gilang"
	var reverse string = ""
	var a = len(KataAwal)

	fmt.Println("Apakah kata", KataAwal, "adalah palindrom?")

	for i := a - 1; i >= 0; i-- {
		reverse = reverse + string(KataAwal[i])
	}
	if strings.ToLower(KataAwal) == strings.ToLower(reverse) {
		fmt.Println("Ya, Kata", KataAwal, "adalah palindrom")
	} else {
		fmt.Println("Tidak, Kata", KataAwal, "Bukan palindrom")
	}

}
