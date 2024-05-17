package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/mygo/notes"
	"example.com/mygo/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoContent := getUserInput("Todo Content: ")

	userNote, err := note.New(title, content)	
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(userNote)
	if err != nil {		
		return
	}

	todo, err := todo.New(todoContent)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todo)
	if err != nil {		
		return
	}
	
}

// function accepting anything (Any value allowed)
func randomFunction2(value any) {
	fmt.Println(value)
}
// function accepting anything
func randomFunction(value interface{}) {
	// fmt.Println(value)
	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float64:", value)
	case string:
		fmt.Println("String:", value)
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the content failed")
		return err
	}
	fmt.Println("Saving the content succeeded")
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