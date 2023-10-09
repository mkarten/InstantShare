package db

import (
	"InstantShare/models"
	"errors"
	"github.com/google/uuid"
)

func CreateEvent(name, userUUID string) (models.Event, error) {
	// check if event name already exists
	e, _ := GetEventByName(name)
	if e.Name != "" {
		return models.Event{}, errors.New("event name already exists")
	}
	// create a token
	token := uuid.New().String()
	newEvent := models.Event{
		Name:     name,
		Token:    token,
		UserUUID: userUUID,
	}
	// insert the event into the database
	_, err := DB.Conn.Exec("INSERT INTO events (name, token, userUUID) VALUES (?, ?, ?)", newEvent.Name, newEvent.Token, newEvent.UserUUID)
	if err != nil {
		return models.Event{}, err
	}
	return newEvent, nil
}

func GetEventById(id int) (models.Event, error) {
	var event models.Event
	err := DB.Conn.QueryRow("SELECT * FROM events WHERE id = ?", id).Scan(&event.Id, &event.Name, &event.Token, &event.UserUUID)
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

func GetEventByToken(token string) (models.Event, error) {
	var event models.Event
	err := DB.Conn.QueryRow("SELECT * FROM events WHERE token = ?", token).Scan(&event.Id, &event.Name, &event.Token, &event.UserUUID)
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

func GetEventByName(name string) (models.Event, error) {
	var event models.Event
	err := DB.Conn.QueryRow("SELECT * FROM events WHERE name = ?", name).Scan(&event.Id, &event.Name, &event.Token, &event.UserUUID)
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

func GetEventsByUserUUID(userUUID string) ([]models.Event, error) {
	var events []models.Event
	rows, err := DB.Conn.Query("SELECT * FROM events WHERE userUUID = ?", userUUID)
	if err != nil {
		return []models.Event{}, err
	}
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.Id, &event.Name, &event.Token, &event.UserUUID)
		if err != nil {
			return []models.Event{}, err
		}
		events = append(events, event)
	}
	return events, nil
}

func UpdateEvent(event models.Event) (models.Event, error) {
	// update the event
	_, err := DB.Conn.Exec("UPDATE events SET name = ?, token = ?, user_uuid = ? WHERE id = ?", event.Name, event.Token, event.UserUUID, event.Id)
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

func DeleteEvent(id int) error {
	// delete the event
	_, err := DB.Conn.Exec("DELETE FROM events WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
