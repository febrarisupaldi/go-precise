package main

import (
	"github.com/febrarisupaldi/go-precise/db"
	"github.com/febrarisupaldi/go-precise/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
