package main

import (
	"fmt"

	"github.com/pedro-hca/go-studies/07-practice-prj-price-calculator/filemanager"
	"github.com/pedro-hca/go-studies/07-practice-prj-price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
