package models

import (
	"time"

	"example.com/event-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := "INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
	// events = append(events, e)
	// add DB save logic here
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
func (event Event) UpdateEvent() error {
	query := "UPDATE events SET name = ?, description = ?, location = ?, date_time = ?, user_id = ? WHERE id = ?"
	stmp, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmp.Close()
	_, err = stmp.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID)
	if err != nil {
		return err
	}

	return nil
}
func (event Event) DeleteEvent() error {
	query := "DELETE FROM events WHERE id = ?"
	stmp, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmp.Close()
	_, err = stmp.Exec(event.ID)
	return err
}
