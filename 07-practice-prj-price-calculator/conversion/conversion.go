package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(stringsToConvert []string) ([]float64, error) {
	floatSlice := make([]float64, len(stringsToConvert))
	for index, strValue := range stringsToConvert {
		floatPrice, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			return nil, errors.New("Failed to convert strings")
		}
		floatSlice[index] = floatPrice
	}
	return floatSlice, nil
}
