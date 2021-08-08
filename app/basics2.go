package main

import (
	"fmt"
	"math"
)

func main() {
	// ---Defer 12/14
	defer fmt.Println("end")

	// --- For 1/14
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// --- If with a short statement 6/14
	var pow func(x, n, lim float64) float64 = func(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		}
		// fmt.Println(v)
		return lim
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// --- Exercise: Loops and Functions 8/14
	var sqrt func(x float64) float64 = func(x float64) float64 {
		z := 1.0
		for i := 1; i < 10; i++ {
			z -= (z*z - x) / (2 * z)
		}
		return z
	}

	fmt.Println("-----")
	for i := 1.0; i < 10; i++ {
		fmt.Println(sqrt(i), math.Sqrt(i))
	}
}
