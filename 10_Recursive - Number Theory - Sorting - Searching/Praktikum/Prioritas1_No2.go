package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	if len(items) == 0 {
		return []pair{}
	}
	Itemnya := items[0]
	Nilai := 1
	for i := 1; i < len(items); i++ {
		if items[i] == Itemnya {
			Nilai++
		}
	}
	ItemBaru := []string{}
	for i := 1; i < len(items); i++ {
		if items[i] != Itemnya {
			ItemBaru = append(ItemBaru, items[i])
		}
	}
	Hasil := MostAppearItem(ItemBaru)
	index := 0
	for i := 0; i < len(Hasil); i++ {
		if Hasil[i].count < Nilai {
			index = i + 1
		} else {
			break
		}
	}
	Pair2 := pair{name: Itemnya, count: Nilai}
	Hasil = append(Hasil[:index], append([]pair{Pair2}, Hasil[index:]...)...)

	return Hasil
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C->1 D->2 B->3 A->4

}
