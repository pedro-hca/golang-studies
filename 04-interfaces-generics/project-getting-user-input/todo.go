package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pedro-hca/go-studies/04-interfaces-generics/project-getting-user-input/note"
	"github.com/pedro-hca/go-studies/04-interfaces-generics/project-getting-user-input/todo"
)

type saver interface {
	Save() error
}
type displayer interface {
	Display()
}
type outputtable interface {
	saver
	displayer
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(45)
	result := add(1, 2)
	fmt.Println(result)

	title, content := getNoteData()
	todoText := getUserInput("Todo text:")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

}
func add[T int | float64 | string](a, b T) T {
	return a + b
}
func printSomething(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Int: ", intVal)
	}
	float64Val, ok := value.(float64)
	if ok {
		fmt.Println("Int: ", float64Val)
	}
	stringVal, ok := value.(string)
	if ok {
		fmt.Println("Int: ", stringVal)
	}
	switch value.(type) {
	case int:
		fmt.Println("Int: ", value)
	case float64:
		fmt.Println("Float64: ", value)
	case string:
		fmt.Println("String: ", value)
	}
}
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
