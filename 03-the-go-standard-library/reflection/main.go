package main

import (
	"fmt"
	"pluralsight/reflection/media"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func main() {
	var myFav media.Catalogable = &media.Movie{}
	myFav.SetMovie("Harry Potter", media.G, 90.2)
	fmt.Println(myFav.GetMovie())
	m := media.NewMovie("Lord of the rings", media.R, 99.9)
	fmt.Println(m)
	m.SetTitle("Lord of the Rings")
	fmt.Println(m)
	fmt.Printf("The value is %v \n", reflect.TypeOf(m))
	fmt.Printf("The value is %v \n", reflect.ValueOf(m))
	fmt.Printf("The value is %v \n", reflect.ValueOf(m).Kind())
	fmt.Printf("The value is %v \n", reflect.Pointer)

	movies := make([]media.Movie, 3)
	movies = append(movies, *media.NewMovie("Lord of the rings", media.R, 99.9))
	movies = append(movies, *media.NewMovie("Dracula", media.R, 85.9))
	movies = append(movies, *media.NewMovie("Interstelar", media.R, 50.9))
	fmt.Println(movies)

	eType := reflect.TypeOf(movies)
	newMovies := reflect.MakeSlice(eType, 0, 0)
	newMovies = reflect.Append(newMovies, reflect.ValueOf(*media.NewMovie("Lord of the rings", media.R, 99.9)))
	fmt.Println(newMovies)

	outTitle := "Func test"
	timed := MakeTimedFunction(properTitle).(func(string) string)
	newTitle := timed(outTitle)
	fmt.Println(newTitle)
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

func MakeTimedFunction(f interface{}) interface{} {
	rf := reflect.TypeOf(f)

	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}
	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}
