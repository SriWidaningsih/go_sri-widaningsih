package main

import "fmt"

func BilBiner(number int) {
	fmt.Print("[")
	for i := 0; i <= number; i++ {
		fmt.Printf("%b", i)
		if i != number {
			fmt.Print(",")
		}
	}
	fmt.Print("]")
}
func main() {
	BilBiner(2)
	fmt.Println()
	BilBiner(3)
	fmt.Println()
}
