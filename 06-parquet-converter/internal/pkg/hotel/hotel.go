package hotel

type Hotel struct {
	// Id     string
	Name   string  `json:"name" parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	City   string  `json:"city" parquet:"name=city, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Review float64 `json:"review" parquet:"name=review, type=FLOAT64"`
}

type Booking struct {
	CheckintDate string `parquet:"checkin_date"`
	CheckoutDate string `parquet:"checkout_date"`
	Hotel        Hotel  `parquet:"hotel"`
}

func NewHotel() *Hotel {
	return &Hotel{}
}
