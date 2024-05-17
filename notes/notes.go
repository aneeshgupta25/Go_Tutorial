package notes

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title, Content string
	CreatedAt time.Time
}

func (note *Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)	
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")	
	fileName = strings.ToLower(fileName) + ".json"
	
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	} else {
		text = strings.Trim(text, "\n")
		text = strings.Trim(text, "\r")
		return text
	}
}

func New() *Note {
	var noteTitle, noteContent string

	fmt.Print("Note title: ")
	noteTitle = readInput()	
	if noteTitle == "" {
		return &Note{}
	}

	fmt.Print("Note content: ")
	noteContent = readInput()	
	if noteContent == "" {
		return &Note{}
	}

	note := &Note{
		Title: noteTitle,
		Content: noteContent,
		CreatedAt: time.Now(),
	} 
	note.Display()
	err := note.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return &Note{}
	} 

	fmt.Println("Successfully saved the note!")
	return note
}