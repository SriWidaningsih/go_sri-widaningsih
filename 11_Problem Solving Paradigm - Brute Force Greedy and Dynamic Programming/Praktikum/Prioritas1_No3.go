package main

import "fmt"

func fibonacciTopDown(numb int, listFibonacci map[int]int) int {
	if numb < 2 {
		return numb
	}
	if val, ok := listFibonacci[numb]; ok {
		return val
	}
	listFibonacci[numb] = fibonacciTopDown(numb-1, listFibonacci) + fibonacciTopDown(numb-2, listFibonacci)
	return listFibonacci[numb]
}

func main() {
	list := make(map[int]int)
	fmt.Printf("Bilangan Fibonacci ke-%d adalah %d\n", 0, fibonacciTopDown(0, list))
	fmt.Printf("Bilangan Fibonacci ke-%d adalah %d\n", 1, fibonacciTopDown(1, list))
	fmt.Printf("Bilangan Fibonacci ke-%d adalah %d\n", 2, fibonacciTopDown(2, list))
	fmt.Printf("Bilangan Fibonacci ke-%d adalah %d\n", 3, fibonacciTopDown(3, list))
	fmt.Printf("Bilangan Fibonacci ke-%d adalah %d\n", 4, fibonacciTopDown(4, list))
	fmt.Printf("Bilangan Fibonacci ke-%d adalah %d\n", 5, fibonacciTopDown(5, list))
}
