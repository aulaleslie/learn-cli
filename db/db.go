package db

import "learn-cli/models"

type IDatabase interface {
	Open() error
	Close() error
}

type FileDB struct {
}

var users []*models.User

func NewFileDB() IDatabase {
	return &FileDB{}
}

func (fileDB *FileDB) Open() error {
	// TODO unmarshal json store ke Users
	// ReadFile -> Users
	return nil
}

func (fileDB *FileDB) Close() error {
	// TODO marshal json store ke file
	// Users -> WriteFile
	return nil
}

func GetUsers() []*models.User {
	return users
}
