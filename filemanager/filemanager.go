package filemanager

import (
	"bufio"
	"errors"
	"os"
	"encoding/json"
)

type FileManager struct {
	InputFileName string
	OutputFileName string
}

func New(inputFileName string, outputFileName string) FileManager {
	return FileManager{
		InputFileName: inputFileName,
		OutputFileName: outputFileName,
	}
}

func (fm FileManager) ReadFile() ([]string, error) {
	file, err := os.Open(fm.InputFileName)
	if err != nil {		
		return nil,errors.New("Couldn't open given file!")
	}
	scanner := bufio.NewScanner(file)
	fileContent := []string{}
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}
	fileError := scanner.Err()
	if(fileError != nil) {
		file.Close()
		return nil, errors.New("Couldn't read content of file!")
	}
	file.Close()
	return fileContent,nil
}

// func WriteJSON(filename string, content interface{}) error {
// 	jsonEncoding, err := json.Marshal(content)
// 	if(err != nil) {
// 		return errors.New("Couldn't convert to JSON!")
// 	}
// 	os.WriteFile(filename, jsonEncoding, 0644)
// 	return nil
// }

func (fm FileManager) WriteResult(content interface{}) error {
	
	file, err := os.Create(fm.OutputFileName)
	if err != nil {
		return errors.New("Couldn't convert to JSON!")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(content)
	if err != nil {
		return errors.New("Couldn't convert to JSON!")
	}
	return nil
}