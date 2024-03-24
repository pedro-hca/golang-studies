package utils

import "encoding/json"

func IsJson(s string) error {
	var js struct{}

	if err := json.Unmarshal([]byte(s), &js); //o primeiro parametro passado é a string que deve ser um json convertida para um slice de bytes e o segundo parametro um pointeiro para struct onde o json deve ser armazenado quando deserializado
	err != nil {
		return err
	}

	return nil
}
