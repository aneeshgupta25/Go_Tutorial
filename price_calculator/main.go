package main

import (
	"fmt"
	// "strconv"
	// "example.com/mygo/cmdmanager"
	"example.com/mygo/price_calculator/filemanager"
	"example.com/mygo/price_calculator/prices"
	// "example.com/mygo/cmdmanager"
	// "example.com/mygo/iomanager"
)

func main() {
	// var prices []float64 = []float64{1,2,3,4}
	var taxRates []float64 = []float64{0,0.01,0.05,0.07}

	// result := map[float64][]float64{}
	// for _, rate_value := range taxRates {
	// 	updatedPriceList := make([]float64, len(prices))
	// 	for price_index, price_value := range prices {			
	// 		updatedPriceList[price_index], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", price_value*(1+rate_value)), 64)
	// 	}
	// 	result[rate_value] = updatedPriceList
	// }
	// printMap(&result)

	// taxIncludedJob, _ := prices.New(0.05)
	// fmt.Println(taxIncludedJob.InputPriceList)
	
	for _, rate_value := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.0f.json", rate_value*100))
		// cmd := cmdmanager.New()
		err := prices.New(fm, rate_value).Process()
		if err != nil {
			fmt.Println("Could not process job!")
			fmt.Println(err)
		}
	}
}

func printMap(m *map[float64][]float64) {
	for key,value := range *m {
		fmt.Println(key, value)
	}
}