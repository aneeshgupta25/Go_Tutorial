package conversion

import (
	"errors"
	"strconv"	
)

func ConvertStringsToFloats(stringSlice []string) ([]float64, error) {
	
	floatSlice := []float64{}
	for _, strValue := range stringSlice {
		floatValue, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			return nil, errors.New("Failed to convert from string to float")
		}
		floatSlice = append(floatSlice, floatValue)
	}
	return floatSlice,nil
}