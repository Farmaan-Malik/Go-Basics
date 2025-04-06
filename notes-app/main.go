package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/notes-app/note"
	"example.com/notes-app/todo"
)

type saver interface {
	Save() error
}
type outputtable interface {
	saver
	Display()
}

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(userNote)
	todoData := getUserInput("Enter Todo")
	userTodo, err := todo.New(todoData)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(userTodo)
}

func saveData(data saver) {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving file")
		return
	}
	fmt.Println("File saved successfully")

}
func outputData(data outputtable) {
	data.Display()
	saveData(data)
}

func getUserInput(prompt string) string {
	fmt.Printf("%v: ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func getNoteData() (string, string, error) {
	title := getUserInput("Enter Note title")
	content := getUserInput("Enter Note content")
	if title == "" || content == "" {
		return "", "", errors.New("title or content cannot be empty")
	}
	return title, content, nil
}
