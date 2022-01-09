package main

import (
	"echo-gorm/db"
	"echo-gorm/route"
)

func main() {
	db.Init()
	r := route.Init()

	r.Logger.Fatal(r.Start(":1323"))
}
