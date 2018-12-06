package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.101.59"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	go selectL(db)
	go update(db, "name B2")

	time.Sleep(time.Second * 3)
}

func selectL(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	_, err = tx.Exec(
		`SET TRANSACTION ISOLATION LEVEL SERIALIZABLE READ ONLY DEFERRABLE;`)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second)

	rr, err := tx.Query(`select name from lookbooks;`)
	if err != nil {
		panic(err)
	}


	time.Sleep(time.Second)
	err = tx.Commit()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func update(db *sql.DB, name string) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	_, err = tx.Exec(
		`update lookbooks SET name = $1 where lookbooks.lookbook_id='1946b25e-5140-464d-99b3-a71719c1ca9c'`, name)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second)
	err = tx.Commit()
	if err != nil {
		fmt.Println(name + ": " + err.Error())
	}
}