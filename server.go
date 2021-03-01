package main

import (
	"github.com/febrarisupaldi/go-learning-api/db"
	"github.com/febrarisupaldi/go-learning-api/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
