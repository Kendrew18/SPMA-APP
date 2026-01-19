package main

import (
	"Template-golang/db"
	"Template-golang/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":38600"))
}
