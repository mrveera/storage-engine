package main

import "db/db"

func main() {

	dbx := db.New("rocks")
	defer dbx.Close()
	dbx.Set("id", 123)
	dbx.Set("name", "veera")
	println(dbx.Get("name"))
}
