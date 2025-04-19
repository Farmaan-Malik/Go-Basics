package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Events struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

var events = []Events{}

func (e *Events) Save() error {
	e.ID = len(events)
	e.UserID = len(events) + rand.Intn(10)
	e.DateTime = time.Now()
	events = append(events, *e)
	err := writeJson(events)
	if err != nil {
		if len(events) > 0 {
			events = events[:len(events)-1]
		}
		fmt.Println(err)
		return err
	}
	fmt.Println(events)
	return nil
}

func GetAllEvents() *[]Events {
	return &events
}

func writeJson[T Events | []Events](e T) error {
	jsonData, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = os.WriteFile("events.json", jsonData, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
