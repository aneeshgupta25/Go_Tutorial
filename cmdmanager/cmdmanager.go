package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() CMDManager {
	return CMDManager{}
}

func (cmd CMDManager) ReadFile() ([]string, error) {
	fmt.Println("Kindly enter the prices...")

	prices := []string{}
	for {
		var price string
		fmt.Scanln(&price)

		if price == "0" {
			break;
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmd CMDManager) WriteResult(content interface{}) error {
	fmt.Println(content)
	return nil
}