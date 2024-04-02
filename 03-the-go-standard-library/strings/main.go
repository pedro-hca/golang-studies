package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	string1 := "i love go"
	string2 := "i love go"
	fmt.Println(CompareCaseIns(string1, string2))
	string3 := "one |two |three |four"
	SplitString(string3)
	result := strings.Contains(string3, "two")
	fmt.Println(result)
	result = strings.HasPrefix(string3, "one")
	fmt.Println(result)
	result = strings.HasSuffix(string3, "four")
	fmt.Println(result)
	newString1 := strings.Replace(string1, "go", "golang", -1)
	fmt.Println(newString1)
	string4 := " string test string test "
	r, _ := regexp.Compile(`s([a-z]+)g`)
	fmt.Println(r.MatchString(string4))
	fmt.Println(r.FindAllString(string4, -1))
	fmt.Println(r.FindStringIndex(string4))
	fmt.Println(r.ReplaceAllString(string4, "golang"))
	fmt.Println(strings.Trim(string4, " "))
	fmt.Println(strings.TrimLeft(string4, " "))
	fmt.Println(strings.ToUpper(string4))
	s := properTitle("welcome to the go programming language")
	fmt.Println(s)
}

func CompareCaseIns(a, b string) bool {
	if len(a) == len(b) {
		if strings.ToLower(a) == strings.ToLower(b) {
			return true
		}
	}
	return false
}

func SplitString(s string) {
	stringCollection := strings.Split(s, "|")

	for i := range stringCollection {
		fmt.Println(stringCollection[i])
	}
}

func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
