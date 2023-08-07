package db

import (
	"encoding/json"
	"learn-cli/helper"
	"learn-cli/models"
	"os"
)

type IDatabase interface {
	Open() error
	Close() error
}

type FileDB struct {
	filePath string
}

var users []*models.User

func NewFileDB(filePath string) IDatabase {
	return &FileDB{
		filePath: filePath,
	}
}

func (fileDB *FileDB) Open() error {
	// Check if the file exists before trying to open it
	_, err := os.Stat(fileDB.filePath)
	if os.IsNotExist(err) {
		// File does not exist, initialize users slice and return
		users = make([]*models.User, 0)
		return nil
	} else if err != nil {
		// Return an error if there was an issue with the file
		return err
	}
	//file exists, open it and read its content
	file, err := os.Open(fileDB.filePath)
	helper.PanicIfError(err)
	defer file.Close()

	// Use json.NewDecoder to decode the JSON data into the users slice
	err = json.NewDecoder(file).Decode(&users)
	helper.PanicIfError(err)
	return nil
}

func (fileDB *FileDB) Close() error {
	writer, err := os.Create(fileDB.filePath)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	return json.NewEncoder(writer).Encode(&users)
}

func GetUsers() []*models.User {
	return users
}

func SetUsers(newUsers []*models.User) {
	users = newUsers
}

func GetLoggedInUser() *models.User {
	for _, user := range users {
		if user.IsLogin {
			return user
		}
	}
	return nil
}
