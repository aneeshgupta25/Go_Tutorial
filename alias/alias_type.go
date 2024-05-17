package main

import "fmt"

type str string

func (text str) log() {
	fmt.Println(text)
}

func main4() {
	var x str = "Aneesh"
	x.log()
}