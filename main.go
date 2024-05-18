package main

import (
	"fmt"
)

func main() {
	a := []int{1,2,3,4}
	fmt.Println(transformNumbers(&a, double))
	fmt.Println(transformNumbers(&a, triple))

	myfunction := helper(2)
	fmt.Println(myfunction(3))

	// anonymous function
	anonymFunction(&a, func(num int) int {		
		return num
	})

	//variadic function
	fmt.Println(variadicFunction(1,2,3,4,5,6,7,8,9,10))
	fmt.Println(variadicFunction(a...))

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	a := []int{}
	for _,value := range *numbers {
		a = append(a, transform(value))
	}
	return a
}

func double(val int) int {
	return val*2
}

func triple(val int) int {
	return val*3
}

// returning a function
func helper(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func anonymFunction(a *[]int, anFunc func(int) int) {
	for _ , value := range *a {
		fmt.Println(anFunc(value))
	}
}

func variadicFunction(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}