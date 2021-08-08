package main

import (
	"fmt"
	"math"
)

func addAndSub(x, y int) (int, int) {
	return x + y, x - y
}

func main() {
	// --- Functions 4/17
	fmt.Println(addAndSub(42, 13))

	// --- Variables 8/17
	// varやconstをつけるとフォーマット時にエラーになった
	hoge := "hoge"
	fmt.Println(hoge)
	hoge = "fuga"
	fmt.Println(hoge)

	// --- Type conversions 13/17
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// --- Type inference 14/17
	// overflows int
	// var v = 11111111111111111111111111111111111
	var v = 1.1
	fmt.Printf("v is of type %T\n", v)
}
