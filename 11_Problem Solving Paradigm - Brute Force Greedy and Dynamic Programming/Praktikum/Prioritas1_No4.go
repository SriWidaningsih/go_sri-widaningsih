package main

import "fmt"

func simpleEquation(a, b, c int) {
	var x, y, z int
	for x = 1; x <= 10000; x++ {
		for y = 1; y <= 10000; y++ {
			for z = 1; z <= 10000; z++ {
				if x != y && y != z && z != x && x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
					fmt.Println(x, y, z)
					return
				}
			}
		}
	}
	fmt.Println("No Solution")

}

func main() {
	simpleEquation(1, 2, 3)  // no solution
	simpleEquation(6, 6, 14) // 1 2 3
}
