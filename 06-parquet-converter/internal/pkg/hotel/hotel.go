package hotel

type Hotel struct {
	Name   string  `parquet:"name"`
	City   string  `parquet:"city"`
	Review float64 `parquet:"review"`
}

func NewHotel() *Hotel {
	return &Hotel{}
}
