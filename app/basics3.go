package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	// --- Pointers 1/27
	i := 42
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	// p = 10 // cannot use 10 (type int) as type *int in assignment

	// --- Pointers to structs 4/27
	v := Vertex{1, 2}
	pv := &v
	pv.X = 1e9
	fmt.Println(v)
	(*pv).X = 2
	fmt.Println(v)

	// --- Struct Literals 5/27
	v1 := &Vertex{1, 2}
	v2 := &Vertex{3, 4}
	v3 := Vertex{5, 6}
	v4 := Vertex{7, 8}
	v1.X = 9
	v3.X = 9
	fmt.Println(v1, v2, v3, v4)

	var f1 func(v *Vertex) = func(v *Vertex) {
		v.Y = 100
	}
	var f2 func(v Vertex) = func(v Vertex) {
		v.Y = 100
	}
	f1(v2)
	f2(v4)
	fmt.Println(v1, v2, v3, v4)

	// --- Slices are like references to arrays 8/27
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// 関数に渡して値を変えても影響なし
	var f3 func(arr [4]string) = func(arr [4]string) {
		arr[3] = "XXX"
	}
	f3(names)
	fmt.Println(names)
	// &で渡すと参照になる
	var f4 func(arr *[4]string) = func(arr *[4]string) {
		arr[3] = "XXX"
	}
	f4(&names)
	fmt.Println(names)

	// 配列をスライスとしては渡せない
	// var f func(arr []string) = func(arr []string) { }
	// f(names) // cannot use names (type [4]string) as type []string in argument to f

	// --- Slice length and capacity 11/27
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s[:4])
	printSlice(s[:0])
	// printSlice(s[:10]) // panic: runtime error: slice bounds out of range [:10] with capacity 6
	fmt.Printf("len=%d cap=%d %v\n", len(names), cap(names), names)

	// capが同じなら渡せるとかもない。型が違う
	// var f func(arr [6]int) = func(arr [6]int) { }
	// f(s[:]) // cannot use s (type []int) as type [6]int in argument to f

	// --- Nil slices 12/27
	var s2 []int
	if s2 == nil {
		fmt.Println("nil!")
	}
	// fmt.Println(s2[0]) // panic: runtime error: index out of range [0] with length 0

	// --- Appending to a slice 15/27
	printSlice(s)
	s = append(s, 1)
	printSlice(s) // capが2倍に

	// --- Range 16/27
	for i, v := range s {
		fmt.Printf("i %d = %d\n", i, v)
	}

	// --- Exercise: Slices 18/27
	// pic.Show(func(dx, dy int) [][]uint8 {
	// 	var result [][]uint8
	// 	for i := 0; i < dy; i++ {
	// 		x := make([]uint8, dx)
	// 		for j := 0; j < dx; j++ {
	// 			x = append(x, uint8(j))
	// 		}
	// 		result = append(result, x)
	// 	}
	// 	return result
	// })

	// --- Maps 19/27
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["test"] = Vertex{1, 2}
	fmt.Println(m["test"])

	m2 := make(map[string]Vertex)
	m2["test"] = Vertex{3, 4}
	fmt.Println(m2["test"])

	var m3 map[string]Vertex
	if m3 == nil {
		fmt.Println("nil!")
	}

	m4 := map[string]Vertex{"test": {1, 2}}
	fmt.Println(m4["test"])

	// --- Mutating Maps 22/27
	elm, ok := m4["hoge"]
	fmt.Println(elm, ok)

	// --- Exercise: Maps 23/27
	var WordCount func(s string) map[string]int = func(s string) map[string]int {
		result := make(map[string]int)
		for _, v := range strings.Fields(s) {
			_, ok := result[v]
			if ok {
				result[v]++
			} else {
				result[v] = 1
			}
		}
		return result
	}
	fmt.Println(WordCount("hoge hoge fuga"))

	// --- Function values 24/27
	hoge := func(s string) {
		fmt.Println(s)
	}
	hoge("hogehoge")

	// --- Function closures 25/27
	adder := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			fmt.Println(sum)
			return sum
		}
	}
	adds := adder()
	adds(1)
	adds(5)
	adds(4)

	// --- Exercise: Fibonacci closure 26/27
	fibonacci := func() func() int {
		n := []int{0}
		return func() int {
			if len(n) < 3 {
				current := n[len(n)-1]
				n = append(n, 1)
				return current
			} else {
				current := n[len(n)-1]
				before := n[len(n)-2]
				n = append(n, current+before)
				return current
			}
		}
	}
	fmt.Println("---")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
