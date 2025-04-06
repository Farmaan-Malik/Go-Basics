package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("invalid input")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
func (n *Note) Display() {
	fmt.Printf("\n\ntitle:%v \ncontent: %v \n\n\ncreatedAt:%v\n\n", n.Title, n.Content, n.CreatedAt)
}

func (note *Note) Save() (err error) {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, json, 0644)
	if err != nil {
		return err
	}
	return nil
}