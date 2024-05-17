package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"	
)

type Todo struct {	
	Text   string    `json:"text"`	
}

func (todo Todo) Display() {
	fmt.Printf("Your Todo content is %v\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.txt"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Invalid input.")
	}

	return Todo{		
		Text:   text,		
	}, nil
}