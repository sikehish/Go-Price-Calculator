package prices

import (
	"fmt"
	"os"

	"github.com/sikehish/Go-Price-Calculator/conversion"
	"github.com/sikehish/Go-Price-Calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

// Constructor
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30}, //This isnt mandatory as we eventually load data from a file into the array
		TaxRate:     taxRate,
		IOManager:   fm,
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

func (job *TaxIncludedPriceJob) Process() {

	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price + price*job.TaxRate
		result[fmt.Sprintf("%v", price)] = fmt.Sprintf("%.2f", taxIncludedPrice) //we're storing in string format as opposed to using float64 only so that we can elimnate the redundant decimal places
	}

	fmt.Println(job.TaxRate, result)
	job.TaxIncludedPrices = result

	//Creates a directory if it doesnt exist
	if err := os.MkdirAll("results", 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	err := job.IOManager.WriteResult(job)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
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
		fmt.Println(err)
		return
	}

	job.InputPrices = priceData

}
