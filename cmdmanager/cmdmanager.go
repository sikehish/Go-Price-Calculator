package cmdmanager

import (
	"fmt"
)

// Using CLI i/ps and o/ps
type CMDManager struct {
}

func New() CMDManager {
	return CMDManager{}
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")
	var prices []string
	for { //Infinite for loop
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

// any and interface{} are the same
func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
