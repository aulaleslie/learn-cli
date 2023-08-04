package app

import (
	"fmt"
	"os"

	"learn-cli/atm_service"
	"learn-cli/models"
)

type IApp interface {
	Run() error
}

type App struct {
	service atm_service.IService
	users   []*models.User
}

func NewApp(service atm_service.IService, users []*models.User) IApp {
	return &App{
		service: service,
		users:   users,
	}
}

func (a *App) Run() error {

	// baca input command
	args := os.Args

	// cek user

	for _, arg := range args {
		fmt.Println(arg)

		switch arg {
		case "login":
			fmt.Println("from login")
		case "withdraw":
			//user := login()
			//a.service.Withdraw(1, user)
			fmt.Println("from withdraw")
		}
	}
	return nil
}

func login() *models.User {
	// logic cari user
	// balikin user
	return nil
}

func logout() {

}
