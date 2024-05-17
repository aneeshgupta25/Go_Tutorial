package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := a[1:4] // slice
	fmt.Println(b)

	// dynamic arrays -> to add element (append), to delete any element, use slice
	c := []int{}
	c = append(c, 20, 34)
	fmt.Println(c)
}
