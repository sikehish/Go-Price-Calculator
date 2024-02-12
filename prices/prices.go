package prices

import (
	"fmt"

	"github.com/sikehish/Go-Price-Calculator/conversion"
	"github.com/sikehish/Go-Price-Calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

// Constructor
func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30}, //This isnt mandatory as we eventually load data from a file into the array
		TaxRate:     taxRate,
		IOManager:   iom,
	}
}

// // OR
// func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
// 	return &TaxIncludedPriceJob{
// 		InputPrices: []float64{10, 20, 30},
// 		TaxRate:     taxRate,
// 		IOManager:   filemanager.New("prices.txt", fmt.Sprintf("results/result_%.0f.json", taxRate*100)),
// 	}
// }

func (job *TaxIncludedPriceJob) Process() error {

	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price + price*job.TaxRate
		result[fmt.Sprintf("%v", price)] = fmt.Sprintf("%.2f", taxIncludedPrice) //we're storing in string format as opposed to using float64 only so that we can elimnate the redundant decimal places
	}

	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job) //This returns an error and since its the last linein the function, we dont have to handle the error, and return it as it is. If there is no error, itll be nil

}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	//We wrap the below logic in a new function(Strings to Float) under conversion package
	// priceData := make([]float64, len(lines))
	// for idx, price := range lines {
	// 	floatPrice, err := strconv.ParseFloat(price, 64)
	// if err != nil {
	// 	fmt.Println("Converting price to float failed")
	// 	fmt.Println(err)
	// 	pricesFile.Close()
	// 	return
	// }
	// 	priceData[idx] = floatPrice
	// }
	// job.InputPrices = priceData
	// pricesFile.Close()

	priceData, err := conversion.StringsToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = priceData
	return nil

}
