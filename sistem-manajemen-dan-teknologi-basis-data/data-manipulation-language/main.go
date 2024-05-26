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
    // connect to database using func `Connect`
    // db, err := Connect()
    // if err != nil {
    //     panic(err)
    // }

    // // insert data to table `employees`
    // _, err = db.Exec(`INSERT INTO employee
    //   VALUES (1, 'Rizki', 25, 'Jl. Kebon Jeruk', 2000000),
    //   (2, 'Andi', 27, 'Jl. Kebon Sirih', 3000000),
    //   (3, 'Budi', 30, 'Jl. Kebon Melati', 4000000),
    //   (4, 'Caca', 32, 'Jl. Kebon Anggrek', 5000000),
    //   (5, 'Deni', 35, 'Jl. Kebon Mawar', 6000000)`)

    // if err != nil {
    //     panic(err)
    // }

    // fmt.Println("Insert data success")
	// connect to database using func `Connect`
    db, err := Connect()
    if err != nil {
        panic(err)
    }

    // update data to table `employees`
    _, err = db.Exec(`UPDATE employee SET salary = 1000000 WHERE id = 1`)

    if err != nil {
        panic(err)
    }

    fmt.Println("Update data success")


    // delete data to table `employees`
    _, err = db.Exec(`DELETE FROM employee WHERE id = 5`)
    if err != nil {
        panic(err)
    }

    fmt.Println("Delete data success")
}