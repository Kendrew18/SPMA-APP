package main

import (
	"SPMA-APP/db"
	"SPMA-APP/routes"
)

func main() {
	db.DatabaseInit()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":38600"))
}
