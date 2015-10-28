package main

import "fmt"
import "reflect"

func main() {

	var a [5]int
	var c = new([5]int)
	d := make([]int, 4)
	c = &a
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(c))
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
