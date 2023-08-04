package main

import (
	"learn-cli/app"
	"learn-cli/atm_service"
	"learn-cli/db"
)

func main() {
	file := db.NewFileDB()

	err := file.Open()
	panicIfError(err)

	atmService := atm_service.NewService()

	cmd := app.NewApp(atmService, db.GetUsers())
	err = cmd.Run()
	panicIfError(err)

	_ = file.Close()
	panicIfError(err)
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func atm() {
	// read file

	// parsing json store ke variable

	// baca argumen

	// proces input

	// close file
}
