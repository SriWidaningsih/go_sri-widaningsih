package main

import (
	"fmt"
	"sort"
)

func MaximumBuyProduct(uang int, HargaProduk []int) int {
	sort.Ints(HargaProduk)
	TotalPerhitungan := 0
	for _, Harga := range HargaProduk {
		if uang-Harga < 0 {
			break
		}
		TotalPerhitungan++
		uang -= Harga
	}
	return TotalPerhitungan
}

func main() {
	fmt.Println(MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}))      // 3
	fmt.Println(MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})) // 4
	fmt.Println(MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000}))   // 4
	fmt.Println(MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000}))           // 1
	fmt.Println(MaximumBuyProduct(0, []int{10000, 30000}))                        // 0
}
