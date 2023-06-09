package main

import "fmt"

//Membuat FizzBuzz
//Sri Widaningsih
func main() {
	for a := 1; a <= 100; a++ {
		if a%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if a%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(a)
	}
}
