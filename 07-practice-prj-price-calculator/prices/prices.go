package prices

import (
	"fmt"

	"github.com/pedro-hca/go-studies/07-practice-prj-price-calculator/conversion"
	"github.com/pedro-hca/go-studies/07-practice-prj-price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{},
		TaxRate:     taxRate,
	}
}
