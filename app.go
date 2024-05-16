package main

import (
	"errors"
	"fmt"	
	"example.com/mygo/fileops"
)

const fileName = "aneesh.txt"

func main2() {
	greetHello()
	revenue, err := fileops.GetFromFile(fileName)		
	if(err != nil) {
		revenue, err = printAndInput("Revenue")	
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	expenses, err := printAndInput("Expenses")	
	if(err != nil) {
		fmt.Println(err)
		return
	}	
	taxRate, err := printAndInput("Tax Rate")	
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := solve(revenue, expenses, taxRate)
	results := fmt.Sprintf("EBT: %0.2f\nPROFIT: %0.2f\nRATIO: %0.2f")
	fileops.WriteToFile(fileName, results)	

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
	// fmt.Printf("%T", &expenses)

	greetGoodbye()
}

func printAndInput(text string) (float64, error) {
	var userInput float64
	fmt.Printf("%v: ", text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Kindly enter a value greater than zero")
	}
	return userInput, nil
}

func solve(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
