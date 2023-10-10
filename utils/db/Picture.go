package db

import (
	"InstantShare/models"
	"github.com/google/uuid"
)

func AddPictureToEvent(filePath, userUUID, eventToken string) error {
	// generate a uuid for the picture
	uuid := uuid.New().String()
	// insert the picture into the database
	_, err := DB.Conn.Exec("INSERT INTO pictures (uuid, file_path, event_token, user_uuid) VALUES (?, ?, ?, ?)", uuid, filePath, eventToken, userUUID)
	if err != nil {
		return err
	}
	return nil
}

func GetPicturesByEventToken(eventToken string) ([]models.Picture, error) {
	// get all pictures from the database
	rows, err := DB.Conn.Query("SELECT * FROM pictures WHERE event_token = ?", eventToken)
	if err != nil {
		return nil, err
	}
	// create a slice of pictures
	pictures := make([]models.Picture, 0)
	for rows.Next() {
		// create a new picture
		picture := models.Picture{}
		// scan the row into the picture
		err := rows.Scan(&picture.UUID, &picture.FilePath, &picture.EventToken, &picture.UserUUID, &picture.Likes)
		if err != nil {
			return nil, err
		}
		// append the picture to the slice
		pictures = append(pictures, picture)
	}
	// sort the pictures by likes
	for i := 0; i < len(pictures); i++ {
		for j := i + 1; j < len(pictures); j++ {
			if pictures[i].Likes < pictures[j].Likes {
				pictures[i], pictures[j] = pictures[j], pictures[i]
			}
		}
	}
	return pictures, nil
}

func GetPicturesByUserUUIDAndEventToken(userUUID, eventToken string) ([]models.Picture, error) {
	// get all pictures from the database
	rows, err := DB.Conn.Query("SELECT * FROM pictures WHERE user_uuid = ? AND event_token = ?", userUUID, eventToken)
	if err != nil {
		return nil, err
	}
	// create a slice of pictures
	pictures := make([]models.Picture, 0)
	for rows.Next() {
		// create a new picture
		picture := models.Picture{}
		// scan the row into the picture
		err := rows.Scan(&picture.UUID, &picture.FilePath, &picture.EventToken, &picture.UserUUID, &picture.Likes)
		if err != nil {
			return nil, err
		}
		// append the picture to the slice
		pictures = append(pictures, picture)
	}
	return pictures, nil
}

func GetPictureByUUID(uuid string) (models.Picture, error) {
	// create a new picture
	picture := models.Picture{}
	// get the picture from the database
	err := DB.Conn.QueryRow("SELECT * FROM pictures WHERE uuid = ?", uuid).Scan(&picture.UUID, &picture.FilePath, &picture.EventToken, &picture.UserUUID, &picture.Likes)
	if err != nil {
		return picture, err
	}
	return picture, nil
}

func DeleteEventPictures(eventToken string) error {
	// delete all pictures from the database
	_, err := DB.Conn.Exec("DELETE FROM pictures WHERE event_token = ?", eventToken)
	if err != nil {
		return err
	}
	return nil
}
