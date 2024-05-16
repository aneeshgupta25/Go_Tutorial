package fileops

import (
	"os"
	"errors"
	"strconv"
)

func WriteToFile(fileName string, value string) {	
	os.WriteFile(fileName, []byte(value), 0644)
}

func GetFromFile(fileName string) (float64, error) {
	byteCollection, err := os.ReadFile(fileName)
	if(err != nil) {		
		return 0.0, errors.New("Couldn't read File")
	}
	data, err := strconv.ParseFloat(string(byteCollection), 64)
	if(err != nil) {		
		return 0.0, errors.New("Something went wrong!")
	}
	return data, nil
}