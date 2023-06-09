package main

import "fmt"

func PairSum(arr []int, target int) []int {
	A := make(map[int]int)
	for K, number := range arr {
		komp := target - number
		if L, Hasil := A[komp]; Hasil {
			return []int{L, K}
		}
		A[number] = K
	}
	return nil
}

func main() {
	fmt.Println("Nomor dari indeks dibawah ini, jika dijumlahkan akan menghasilkan target")
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
	fmt.Println(PairSum([]int{1, 2, 5, 7, 9}, 16))
}
