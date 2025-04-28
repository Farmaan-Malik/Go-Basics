package models

import (
	"errors"
	"example/rest-api/db"
	"fmt"
	"time"
)

type Events struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time"`
	UserID      int64     `json:"user_id"`
}
type Registration struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"user_id"`
	EventId int64 `json:"event_id"`
}

func (e *Events) Save() error {
	query := `
	INSERT INTO events (name,description,location,dateTime,user_id)
	VALUES (?,?,?,?,?)
	`
	sql, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer sql.Close()
	result, err := sql.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Events, error) {
	query := `
	SELECT * FROM events
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Events
	for rows.Next() {
		var e Events
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func GetEventById(id int64) (*Events, error) {
	query := `
	SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var e Events
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e *Events) UpdateEvent() error {
	query := `
	UPDATE events 
	SET name=?,description=?,location=?,dateTime=?
	WHERE id = ?
	`
	_, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.ID)

	if err != nil {
		fmt.Println("in Function")
		fmt.Println(err)
		return err
	}
	return nil
}

func (e *Events) DeletEvent() error {
	query := `
	DELETE FROM events WHERE id = ?
	`
	_, err := db.DB.Exec(query, e.ID)
	if err != nil {
		fmt.Println("in Function")
		fmt.Println(err)
		return err
	}
	return nil
}
func (e *Events) MakeRegistration(userId int64) error {
	query := `
	INSERT INTO registrations (user_id,event_id) VALUES (?,?)`
	_, err := db.DB.Exec(query, userId, e.ID)
	if err != nil {
		return errors.New("could not create entry")
	}
	return nil
}

func (e *Events) DeleteRegistration(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	_, err := db.DB.Exec(query, e.ID, userId)
	if err != nil {
		return errors.New("could not find entry")
	}
	return nil
}
func GetRegistrations() ([]Registration, error) {
	query := `SELECT * FROM registrations`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New("error fetching data")
	}
	defer rows.Close()
	var registrations []Registration
	for rows.Next() {
		var r Registration
		err = rows.Scan(&r.ID, &r.EventId, &r.UserID)
		if err != nil {
			return nil, err
		}
		registrations = append(registrations, r)
	}
	return registrations, nil
}
