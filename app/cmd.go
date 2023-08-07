package app

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"learn-cli/atm_service"
	"learn-cli/helper"
)

type IApp interface {
	Run() error
}

type App struct {
	service atm_service.IService
}

func NewApp(service atm_service.IService) IApp {
	return &App{
		service: service,
	}
}

func (a *App) Run() error {
	args := os.Args

	for _, arg := range args {

		switch {
		case strings.HasPrefix(arg, "login="):
			a.loginCommand(arg)
			// name := strings.TrimPrefix(arg, "login=")
			// a.service.Login(name)
		case strings.HasPrefix(arg, "deposit="):
			a.depositCommand(arg)
			// amountString := strings.TrimPrefix(arg, "deposit=")
			// amount, err := strconv.Atoi(amountString)
			// helper.PanicIfError(err)
			// a.service.Deposit(int64(amount))
		case strings.HasPrefix(arg, "withdraw="):
			a.withdrawCommand(arg)
			// amountString := strings.TrimPrefix(arg, "withdraw=")
			// amount, err := strconv.Atoi(amountString)
			// helper.PanicIfError(err)
			// a.service.Withdraw(int64(amount))
		case strings.HasPrefix(arg, "transfer="):
			a.transferCommand(args)
			// // find to Args
			// var toArg string
			// for _, a := range args {
			// 	if strings.HasPrefix(a, "to=") {
			// 		toArg = strings.TrimPrefix(a, "to=")
			// 		break
			// 	}
			// }
			// if toArg == "" {
			// 	fmt.Println("Error: 'to=' argument not found for transfer.")
			// }
			// toUser := toArg
			// amountString := strings.TrimPrefix(arg, "transfer=")
			// amount, err := strconv.Atoi(amountString)
			// helper.PanicIfError(err)
			// a.service.Transfer(int64(amount), toUser)
		case strings.HasPrefix(arg, "logout="):
			a.logoutCommand(arg)
			// name := strings.TrimPrefix(arg, "logout=")
			// a.service.Logout(name)
		}
	}
	return nil
}

func (a *App) loginCommand(arg string) {
	name := strings.TrimPrefix(arg, "login=")
	a.service.Login(name)
}

func (a *App) depositCommand(arg string) {
	amountString := strings.TrimPrefix(arg, "deposit=")
	amount, err := strconv.Atoi(amountString)
	helper.PanicIfError(err)
	a.service.Deposit(int64(amount))
}

func (a *App) withdrawCommand(arg string) {
	amountString := strings.TrimPrefix(arg, "withdraw=")
	amount, err := strconv.Atoi(amountString)
	helper.PanicIfError(err)
	a.service.Withdraw(int64(amount))
}

func (a *App) transferCommand(args []string) {
	var toArg string
	for _, arg := range args {
		if strings.HasPrefix(arg, "to=") {
			toArg = strings.TrimPrefix(arg, "to=")
			break
		}
	}

	if toArg == "" {
		fmt.Println("Error: 'to=' argument not found for transfer.")
	}
	name := toArg
	amountString := strings.TrimPrefix(args[1], "transfer=")
	amount, err := strconv.Atoi(amountString)
	helper.PanicIfError(err)
	a.service.Transfer(int64(amount), name)
}

func (a *App) logoutCommand(arg string) {
	name := strings.TrimPrefix(arg, "logout=")
	a.service.Logout(name)
}
