package main

import "fmt"

func AngFibonacci(angka int) int {
	if angka <= 1 {
		return angka
	}
	return AngFibonacci(angka-1) + AngFibonacci(angka-2)

}

func main() {
	fmt.Println(AngFibonacci(0))  // 0
	fmt.Println(AngFibonacci(2))  // 1
	fmt.Println(AngFibonacci(9))  // 34
	fmt.Println(AngFibonacci(10)) // 55
	fmt.Println(AngFibonacci(12)) // 144
	fmt.Println(AngFibonacci(20))
}
