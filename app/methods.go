package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// --- Exercise: Stringers 18/26
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

// --- Exercise: Errors 20/26
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return x, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	// --- Methods 1/26
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// --- Methods continued 3/26
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// --- Pointer receivers 4/26
	v.Scale(10)
	fmt.Println(v.Abs())

	p := &Vertex{3, 4}
	p.Scale(10)
	fmt.Println(p.Abs())
	p.Scale2(10)
	fmt.Println(p.Abs())

	// -- The empty interface 14/26
	describe := func(i interface{}) {
		fmt.Printf("(%v, %T)\n", i, i)
	}
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = Vertex{3, 4}
	// i.Abs undefined (type interface {} is interface with no methods)
	// fmt.Println(i.Abs())

	// --- Type assertions 15/26
	x, ok := i.(Vertex)
	fmt.Println(x, ok)
	fmt.Println(x.Abs())

	// --- Type switches 16/26
	do := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Twice %v is %v\n", v, v*2)
		case string:
			fmt.Printf("%q is %v bytes long\n", v, len(v))
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}
	do(21)
	do("hello")
	do(true)

	// --- Exercise: Stringers 18/26
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	// --- Exercise: Errors 20/26
	fmt.Println("-----")
	fmt.Println(sqrt(-2))
}
