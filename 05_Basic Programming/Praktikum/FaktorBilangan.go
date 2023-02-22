package main

import "fmt"

//Mencari faktor bilangan

func main() {
	n := 30

	fmt.Println("Faktor dari", n, ":")
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)

		}
	}
}
