package prices

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sikehish/Go-Price-Calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// Constructor
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {

	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price + price*job.TaxRate
		result[fmt.Sprintf("%v", price)] = fmt.Sprintf("%.2f", taxIncludedPrice) //we're storing in string format as opposed to using float64 only so that we can elimnate the redundant decimal places
	}

	fmt.Println(job.TaxRate, result)
}

func (job *TaxIncludedPriceJob) LoadData() {
	pricesFile, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(pricesFile)

	var lines []string

	for scanner.Scan() { //It'll iterate as long scanner.Scan() returns true.It returns false when the scan stops, either by reaching the end of the input or an error.
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err() //After Scan returns false, the Err method will return any error that occurred during scanning, except that if it was io.EOF, Err will return nil.

	if err != nil {
		fmt.Println("Reading the file content failed")
		fmt.Println(err)
		pricesFile.Close()
		return
	}

	//We wrap the below logic in a new function under conversion package
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

	pricesData, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		pricesFile.Close()
		return
	}

	job.InputPrices = priceData
	pricesFile.Close()

}
