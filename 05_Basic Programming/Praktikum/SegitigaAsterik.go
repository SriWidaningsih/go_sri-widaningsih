package main

import "fmt"

//membuat segitiga asterik
//Sri Widaningsih

func main() {
	n := 5
	a := 0

	for i := 1; i <= n; i++ {
		a = 0
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for {
			fmt.Print("*")
			a++
			if a == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}
}
