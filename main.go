package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15} //0%,7%,10%,15%
	//We apply different tax rates to each price

	result := make(map[float64][]float64, 3)

	for _, taxRate := range taxRates {
		for _, price := range prices {
			result[taxRate] = append(result[taxRate], price+(price*taxRate))
		}
	}
	// //OR (More efficient way as we dont overwrite the slice in each iteration)
	// for _, taxRate := range taxRates {
	// 	taxIncludedPrices := make([]float64, 3)
	// 	for priceIdx, price := range prices {
	// 		taxIncludedPrices[priceIdx] = price * (1 + taxRate)
	// 	}
	// 	result[taxRate] = taxIncludedPrices
	// }

	fmt.Println(result)
}
