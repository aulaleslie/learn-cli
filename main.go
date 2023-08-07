package main

import (
	"learn-cli/app"
	"learn-cli/atm_service"
	"learn-cli/db"
	"learn-cli/helper"
)

func main() {
	file := db.NewFileDB("users.json")

	err := file.Open()
	helper.PanicIfError(err)

	atmService := atm_service.NewService()

	cmd := app.NewApp(atmService)
	err = cmd.Run()
	helper.PanicIfError(err)

	_ = file.Close()
	helper.PanicIfError(err)
}
