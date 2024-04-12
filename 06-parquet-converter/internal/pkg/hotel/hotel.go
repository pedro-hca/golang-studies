package hotel

type Hotel struct {
	Name   string  `parquet:"name"`
	City   string  `parquet:"city"`
	Review float32 `parquet:"review"`
}

func NewHotel() *Hotel {
	return &Hotel{}
}
