package atm_service

import (
	"fmt"
	"learn-cli/db"
	"learn-cli/models"
	"os"
)

type IService interface {
	Login(name string)
	Deposit(amount int64)
	Withdraw(amount int64)
	Transfer(amount int64, name string)
	Logout(name string)
}

type Service struct {
}

func NewService() IService {
	return &Service{}
}

func (svc *Service) Login(name string) {
	users := db.GetUsers()
	// Check if any user is already logged in
	for _, user := range users {
		if user.IsLogin {
			fmt.Println("Another user is already logged in. Please logout first.")
			fmt.Println("logged in user :", user.Name)
			return
		}
	}

	var existUser *models.User
	for _, user := range users {
		if user.Name == name {
			existUser = user
		}
	}

	if existUser == nil {
		existUser = &models.User{
			Name: name,
		}
		users = append(users, existUser)
		db.SetUsers(users)
	}

	existUser.IsLogin = true
	fmt.Println("your balance is :", existUser.Balance)

}

func (svc *Service) Deposit(amount int64) {
	user := db.GetLoggedInUser()
	if user == nil {
		fmt.Println("You need to login before making transaction")
		os.Exit(1)
	}
	user.Balance += amount
}

func (svc *Service) Withdraw(amount int64) {
	user := db.GetLoggedInUser()
	if user == nil {
		fmt.Println("You need to login before making transaction")
		os.Exit(1)
	}
	user.Balance -= amount
}

func (svc *Service) Transfer(amount int64, name string) {
	users := db.GetUsers()
	loginUser := db.GetLoggedInUser()
	if loginUser == nil {
		fmt.Printf("you need to login before making a transaction")
		return
	}
	var toUser *models.User
	for _, user := range users {
		if user.Name == name {
			toUser = user
		}
	}
	if toUser == nil {
		fmt.Printf("user with the name `%s` not found", name)
		return
	}

	if loginUser.Balance < amount {
		fmt.Printf("insufficient balance. Your current balance is %d", loginUser.Balance)
		return
	}
	//perform to transfer
	loginUser.Balance -= amount
	toUser.Balance += amount
}

func (svc *Service) Logout(name string) {
	users := db.GetUsers()
	for _, user := range users {
		if user.Name == name {
			user.IsLogin = false
		}
	}
}
