package main

import (
	_ "github.com/lib/pq"

	db "cs-backend/db"
	routerr "cs-backend/router"
)

type Data struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	db.SetupPostgres()
	defer db.Db.Close()
	router := routerr.Routes()
	router.Run(":8080")
}
