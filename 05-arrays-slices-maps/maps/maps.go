package main

import "fmt"

type langMap map[int]string

func main() {
	websites := map[string]string{
		"Amazon Web Services": "http://amazon.com",
		"Google":              "http://google.com",
	}
	fmt.Println(websites)
	websites["Linkedin"] = "http://linkedin.com"
	fmt.Println(websites)
	delete(websites, "Google")
	fmt.Println(websites)

	programLang := make(langMap, 3)
	programLang[1] = "Go"
	programLang[2] = "C#"
	programLang[3] = "Typescript"
	fmt.Println(programLang)

	for key, value := range programLang {
		fmt.Println(key)
		fmt.Println(value)
	}
}
