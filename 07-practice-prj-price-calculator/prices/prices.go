package prices

type TaxIncluedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncluedPriceJob(taxRate float64) *TaxIncluedPriceJob {
	return &TaxIncluedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
