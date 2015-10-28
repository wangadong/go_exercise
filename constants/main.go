package main

import "fmt"
import "math"
import "reflect"

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(reflect.TypeOf(int64(d)))
	dd := d / 2
	fmt.Println(dd)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	var x float64 = 123000000000000000000
	fmt.Println(x)

	y := 1 + 3i
	fmt.Println(y)
	fmt.Println(reflect.TypeOf(y))
}
