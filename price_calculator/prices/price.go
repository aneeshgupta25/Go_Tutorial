package prices

import (	
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"example.com/mygo/price_calculator/conversion"
	// "example.com/mygo/filemanager"	
	"example.com/mygo/price_calculator/iomanager"
)

const fileName = "prices.txt"

type TaxIncludedPriceJob struct {
	TaxRate              float64 `json:"tax_rate"`
	InputPriceList       []float64 `json:"input_prices"`
	TaxIncludedPriceList map[string]string `json:"tax_included_prices"`
	IOManager iomanager.IOManager `json:"-"` //minus sign helps to remove this key-value pair from the file
}

func (job *TaxIncludedPriceJob) Load() error {	
	prices,err := job.IOManager.ReadFile()
	if(err != nil) {		
		fmt.Println(err)
		return err
	}
	
	pricesInFloat, err := conversion.ConvertStringsToFloats(prices)
	if err != nil {		
		fmt.Println(err)
		return err
	}

	job.InputPriceList = pricesInFloat	
	return nil	
}

func (job *TaxIncludedPriceJob) Process() error {	
	err := job.Load() 
	if err != nil {
		fmt.Println("Unable to Load Input File!")
		return err
	}	
	updatedPriceList := make(map[string]string)
	for _, price_value := range job.InputPriceList {			
		taxIncludedPrice := price_value*(1+job.TaxRate)
		updatedPriceList[fmt.Sprintf("%.2f", price_value)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(updatedPriceList)	
	job.TaxIncludedPriceList = updatedPriceList
	return job.IOManager.WriteResult(job)	
}

func New(iom iomanager.IOManager, rate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob {
		TaxRate: rate,
		InputPriceList: []float64{10,20,30},
		IOManager: iom,
	}
}

// My New Function
// func New(rate float64) (*TaxIncludedPriceJob, error) {
// 	inputPrices, err := readInputPrices()
// 	if(err != nil) {
// 		return &TaxIncludedPriceJob{}, errors.New("Couldn't read file")
// 	}
// 	return &TaxIncludedPriceJob{
// 		TaxRate:              rate,
// 		InputPriceList:       inputPrices,
// 		// TaxIncludedPriceList: (inputPrices),
// 	}, nil
// }

// Reading file using os.ReadFile()
func (job *TaxIncludedPriceJob) readInputPrices() ([]float64, error) {
	pricesByteColl, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error encountered!")
		return nil,errors.New("Couldn't read prices!")
	}	
	pricesAsStrings := strings.Split(string(pricesByteColl), "\r\n")		
	inputPrices := make([]float64, len(pricesAsStrings))
	for index, stringPrice := range pricesAsStrings {	
		inputPrices[index], _ = strconv.ParseFloat(stringPrice, 64)
	}
	fmt.Println(inputPrices)
	return inputPrices, nil
}