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
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

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
	
	for rate_index, rate_value := range taxRates {
		doneChans[rate_index] = make(chan bool)
		errChans[rate_index] = make(chan error)
		fm := filemanager.New("prices/prices.txt", fmt.Sprintf("results_%.0f.json", rate_value*100))
		// cmd := cmdmanager.New()
		priceJob := prices.New(fm, rate_value)
		go priceJob.Process(doneChans[rate_index], errChans[rate_index]) // goroutines don't return any values	
		// if err != nil {
		// 	fmt.Println("Could not process job!")
		// 	fmt.Println(err)
		// }
	}

	// Tells go to wait unless every channel has emitted at least one value
	// for _, doneChan := range doneChans {
	// 	<- doneChan
	// }

	// To wait for one or more channels to emit a value ->
	for index := range taxRates {
		select {
		case err := <- errChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <- doneChans[index]:
			fmt.Println("Done!")
		}
	}
}

func printMap(m *map[float64][]float64) {
	for key,value := range *m {
		fmt.Println(key, value)
	}
}