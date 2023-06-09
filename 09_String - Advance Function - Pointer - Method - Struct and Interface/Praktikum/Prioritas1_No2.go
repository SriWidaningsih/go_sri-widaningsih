package main

import "fmt"

type Siswa struct {
	Nama  []string
	Nilai []int
}

func (s Siswa) averageNilai() int {
	totalNilai := 0
	for _, nilai := range s.Nilai {
		totalNilai += nilai
	}
	averageNilai := totalNilai / len(s.Nilai)
	return averageNilai
}

func (s Siswa) findMinNilai() (string, int) {
	minNilai := s.Nilai[0]
	minName := s.Nama[0]
	for i := 1; i < len(s.Nilai); i++ {
		if s.Nilai[i] < minNilai {
			minNilai = s.Nilai[i]
			minName = s.Nama[i]
		}
	}
	return minName, minNilai
}

func (s Siswa) findMaxNilai() (string, int) {
	maxNilai := s.Nilai[0]
	maxName := s.Nama[0]
	for i := 1; i < len(s.Nilai); i++ {
		if s.Nilai[i] > maxNilai {
			maxNilai = s.Nilai[i]
			maxName = s.Nama[i]
		}
	}
	return maxName, maxNilai
}
func main() {
	KelasGolang := Siswa{
		Nama:  []string{"Rizky", "Kobar", "Ismail", "Umam", "Yopan"},
		Nilai: []int{80, 75, 70, 75, 60},
	}
	averageNilai := KelasGolang.averageNilai()
	minName, minNilai := KelasGolang.findMinNilai()
	maxName, maxNilai := KelasGolang.findMaxNilai()
	fmt.Printf("Nilai rata-rata Siswa: %d\n", averageNilai)
	fmt.Printf("Nilai terendah dari Siswa: %s %d\n", minName, minNilai)
	fmt.Printf("Nilai tertinggi dari Siswa: %s %d\n", maxName, maxNilai)
}
