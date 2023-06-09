package main

import "fmt"

type Car struct {
	typeMobil string
	fuelIn    int
}

func (e *Car) EstimasiSampai() float64 {
	return float64(e.fuelIn) / (1.5)
}

func main() {
	Tipe := Car{
		typeMobil: "Sport",
		fuelIn:    60,
	}
	fmt.Println("Maka mobil", Tipe.typeMobil, "jika diisi", Tipe.fuelIn, "L bahan bakarnya, akan berjalan sejauh:")

	fmt.Println(Tipe.EstimasiSampai(), "mill")
}
