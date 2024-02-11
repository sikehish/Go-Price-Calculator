package prices

import (
	"bufio"
	"fmt"
	"os"
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

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%v", price)] = price + price*job.TaxRate
	}

	fmt.Println(job.TaxRate, result)
}

func (job TaxIncludedPriceJob) LoadData() {
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

}
