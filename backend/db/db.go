package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB
var err error

func SetupPostgres() {
	host := "dhrztxpidhpbqxmhrjzc.retooldb.com"
	port := 5432
	user := "production"
	password := "EfOe0QPl1pe8lccUpsc_IG"
	dbname := "production"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
