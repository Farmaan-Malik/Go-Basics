package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes-app/note"
)

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
	userNote.DisplayNote()
	err = userNote.SaveNoteJson()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// func getUserInput(prompt string) string {
// 	fmt.Println(prompt)
// 	var value string
// 	fmt.Scanln(&value)
// 	return value
// }

func getUserInput(prompt string) string {
	fmt.Printf("Enter %s: ", prompt)
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
	title := getUserInput("Note title")
	content := getUserInput("Note content")
	return title, content, nil
}
