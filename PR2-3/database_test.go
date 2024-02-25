package main

import (
	"database/sql"
	"fmt"
	"testing"
)

import _ "modernc.org/sqlite"

func TestCreateEvent(t *testing.T) {
	conn, _ := sql.Open("sqlite", "database.sqlite3")
	db := DatabaseWrapper{conn}

	event := Event{0, "gjdls1", "gjdls1", "gjdls", "gjdls", "gjdls"}

	db.CreateEvent(event)
}

func TestReadById(t *testing.T) {
	conn, _ := sql.Open("sqlite", "database.sqlite3")
	db := DatabaseWrapper{conn}

	event := db.ReadEventById(1)

	fmt.Println(event)
}

func TestReadAll(t *testing.T) {
	conn, _ := sql.Open("sqlite", "database.sqlite3")
	db := DatabaseWrapper{conn}

	event := db.ReadEvents()

	fmt.Println(event)
}

func TestUpdateById(t *testing.T) {
	conn, _ := sql.Open("sqlite", "database.sqlite3")
	db := DatabaseWrapper{conn}

	id := 1

	event := db.ReadEventById(id)

	event.EventName = "test"

	db.UpdateEventById(id, event)
}

func TestUpdateByIdNil(t *testing.T) {
	conn, _ := sql.Open("sqlite", "database.sqlite3")
	db := DatabaseWrapper{conn}

	id := 3

	event := Event{}

	db.UpdateEventById(id, event)
}

func TestDeleteById(t *testing.T) {
	conn, _ := sql.Open("sqlite", "database.sqlite3")
	db := DatabaseWrapper{conn}

	id := 3

	db.DeleteEventById(id)
}
