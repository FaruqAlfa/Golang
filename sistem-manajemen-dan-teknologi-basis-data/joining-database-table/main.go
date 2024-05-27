package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "test_db_camp"
)

func Connect() (*sql.DB, error) {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }

    return db, nil
}

type CustomerOrder struct {
    ID        int `sql:"order_id"`
    Name      string `sql:"customer_name"`
    Address   string `sql:"customer_address"`
    OrderDate string `sql:"order_date"`
}

func main() {
    // connect to database using func `Connect`
    db, err := Connect()
    if err != nil {
        panic(err)
    }

    // join table `customers` and `departments`
    rows, err := db.Query(`
    SELECT
        orders.id AS order_id,
        customers.name AS customer_name,
        customers.address AS customer_address,
        orders.order_date
    FROM orders
    INNER JOIN customers
    ON orders.customer_id = customers.id
    `)

    if err != nil {
        panic(err)
    }

    var customerOrders []CustomerOrder

    // iterate over rows
    for rows.Next() {
        var customerOrder CustomerOrder

        // scan data from rows to struct
        err := rows.Scan(
            &customerOrder.ID,
            &customerOrder.Name,
            &customerOrder.Address,
            &customerOrder.OrderDate,
        )

        if err != nil {
            panic(err)
        }

        customerOrders = append(customerOrders, customerOrder)
    }

    fmt.Println(customerOrders)
}
