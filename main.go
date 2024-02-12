package main

import (
	"github.com/sikehish/Go-Price-Calculator/cmdmanager"
	"github.com/sikehish/Go-Price-Calculator/prices"
)

func main() {
	// pricesData := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15} //0%,7%,10%,15%
	//We apply different tax rates to each price
	// result := make(map[float64][]float64, 3)

	// for _, taxRate := range taxRates {
	// 	for _, price := range pricesData {
	// 		result[taxRate] = append(result[taxRate], price+(price*taxRate))
	// 	}
	// }
	// //OR (More efficient way as we dont overwrite the slice in each iteration)
	// for _, taxRate := range taxRates {
	// 	taxIncludedpricesData := make([]float64, 3)
	// 	for priceIdx, price := range pricesData {
	// 		taxIncludedPrices[priceIdx] = price * (1 + taxRate)
	// 	}
	// 	result[taxRate] = taxIncludedPrices
	// }
	//OR
	for _, taxRate := range taxRates {

		// fm := filemanager.New("prices.txt", fmt.Sprintf("results/result_%.0f.json", taxRate*100))
		// pricesJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		//OR
		// pricesJob := prices.NewTaxIncludedPriceJob(taxRate)

		cmdm := cmdmanager.New()
		pricesJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		pricesJob.Process()
	}

}
