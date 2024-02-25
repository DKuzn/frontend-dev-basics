package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Event struct {
	Id         int    `json:"id"`
	EventName  string `json:"event_name"`
	EventPlate string `json:"event_plate"`
	EventType  string `json:"event_type"`
	EventDate  string `json:"event_date"`
	EventTime  string `json:"event_time"`
}

type DatabaseWrapper struct {
	conn *sql.DB
}

func (dw DatabaseWrapper) CreateTable() error {
	_, error := dw.conn.Exec(`
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_name TEXT,
		event_plate TEXT,
		event_type TEXT,
		event_date TEXT,
		event_time TEXT
	);`,
	)

	return error
}

func (dw DatabaseWrapper) CreateEvent(event Event) error {
	_, error := dw.conn.Exec(
		`INSERT INTO events (event_name, event_plate, event_type, event_date, event_time) VALUES (?, ?, ?, ?, ?)`,
		event.EventName, event.EventPlate, event.EventType, event.EventDate, event.EventTime,
	)

	return error
}

func (dw DatabaseWrapper) ReadEventById(id int) Event {
	var event Event
	dw.conn.QueryRow(`SELECT id, event_name, event_plate, event_type, event_date, event_time FROM events WHERE id = ?`, id).Scan(
		&event.Id, &event.EventName, &event.EventPlate, &event.EventType, &event.EventDate, &event.EventTime,
	)

	return event
}

func (dw DatabaseWrapper) ReadEvents() []Event {
	rows, _ := dw.conn.Query(`SELECT id, event_name, event_plate, event_type, event_date, event_time FROM events`)
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		rows.Scan(&e.Id, &e.EventName, &e.EventPlate, &e.EventType, &e.EventDate, &e.EventTime)
		events = append(events, e)
	}

	return events
}

func (dw DatabaseWrapper) UpdateEventById(id int, event Event) error {
	_, error := dw.conn.Exec(
		`UPDATE events SET event_name = ?, event_plate = ?, event_type = ?, event_date = ?, event_time = ? WHERE id = ?`,
		event.EventName, event.EventPlate, event.EventType, event.EventDate, event.EventTime, id,
	)

	return error
}

func (dw DatabaseWrapper) DeleteEventById(id int) error {
	_, error := dw.conn.Exec(
		`DELETE FROM events WHERE id = ?`, id,
	)

	return error
}
