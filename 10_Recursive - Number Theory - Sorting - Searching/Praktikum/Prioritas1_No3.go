package main

import "fmt"

func isPrime(number int, i int) bool {
	if i == 1 {
		return true
	}

	if number%i == 0 {
		return false
	}

	return isPrime(number, i-1)
}

func getNilaiPrime(numbers int, nilai int) int {
	if isPrime(nilai, nilai-1) {
		if numbers == 1 {
			return nilai
		}
		return getNilaiPrime(numbers-1, nilai+1)
	}
	return getNilaiPrime(numbers, nilai+1)

}

func primeX(numb int) int {
	if numb == 1 {
		return 2
	}
	for i := 3; numb >= 1; i += 2 {
		if isPrime(i, i-1) {
			numb--
			if numb == 1 {
				return i
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) //29
}
