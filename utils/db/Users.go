package db

import (
	"InstantShare/models"
	"InstantShare/utils/hash"
	"errors"
	"github.com/google/uuid"
)

func CreateUser(username, password, firsName, lastName string) (models.User, error) {
	// check if the username is already taken
	u, _ := GetUserByUsername(username)
	if u.Username != "" {
		// username already taken
		return models.User{}, errors.New("username already taken")
	}
	// create a uuid
	userUUID := uuid.New().String()
	// hash the password
	hashedPass, err := hash.HashPassword(password)
	if err != nil {
		return models.User{}, err
	}
	// create the user
	newUser := models.User{
		UUID:      userUUID,
		Username:  username,
		Password:  hashedPass,
		FirstName: firsName,
		LastName:  lastName,
	}
	// insert the user into the database
	_, err = DB.Conn.Exec("INSERT INTO users (uuid, username, password, first_name, last_name) VALUES ($1, $2, $3, $4, $5)", newUser.UUID, newUser.Username, newUser.Password, newUser.FirstName, newUser.LastName)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func GetUserByUUID(uuid string) (models.User, error) {
	var user models.User
	err := DB.Conn.QueryRow("SELECT * FROM users WHERE uuid = $1", uuid).Scan(&user.UUID, &user.Username, &user.Password, &user.FirstName, &user.LastName)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := DB.Conn.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&user.UUID, &user.Username, &user.Password, &user.FirstName, &user.LastName)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func UpdateUser(user models.User) (models.User, error) {
	// update the user
	_, err := DB.Conn.Exec("UPDATE users SET username = $1, password = $2, first_name = $3, last_name = $4 WHERE uuid = $5", user.Username, user.Password, user.FirstName, user.LastName, user.UUID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func DeleteUser(user models.User) error {
	// delete the user
	_, err := DB.Conn.Exec("DELETE FROM users WHERE uuid = $1", user.UUID)
	if err != nil {
		return err
	}
	return nil
}
