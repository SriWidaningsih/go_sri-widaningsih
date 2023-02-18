package main

import "fmt"

func main() {
	nilai := -6

	if nilai >= 80 {
		fmt.Println("Nilai A")
	} else if nilai >= 65 {
		fmt.Println("Nilai B")
	} else if nilai >= 50 {
		fmt.Println("Nilai C")
	} else if nilai >= 35 {
		fmt.Println("Nilai D")
	} else if nilai >= 0 {
		fmt.Println("Nilai E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}
