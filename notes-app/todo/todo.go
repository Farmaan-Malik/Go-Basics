package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"content"`
}

func New(content string) (*Todo, error) {
	if content == "" {
		return &Todo{}, errors.New("invalid input")
	}
	return &Todo{
		Content: content,
	}, nil
}
func (n *Todo) Display() {
	fmt.Printf("\n\n\ncontent: %v \n\n", n.Content)
}

func (todo *Todo) Save() (err error) {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, json, 0644)
	if err != nil {
		return err
	}
	return nil
}
