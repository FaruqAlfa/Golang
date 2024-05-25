package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	dns := "host=localhost port=5432 user=postgres password=postgres dbname=test_db_camp sslmode=disable"

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE employee (
		id INT,
		name VARCHAR(255) NOT NULL,
		age INT NOT NULL,
		address VARCHAR(255),
		salary INT
	)`)


	if err != nil {
		panic(err)
	}

	fmt.Println("Table created successfully")
}