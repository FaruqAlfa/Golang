package main

import (
	"database/sql"
	"errors"
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

// menambah tag sql pada struct Employee
type Employee struct {
    ID        int    `sql:"id"`
    Name      string `sql:"name"`
    Age       int    `sql:"age"`
    Address   string `sql:"address"`
    Salary    int    `sql:"salary"`
}



func main() {
	// buat koneksi ke database menggunakan func `Connect`
   db, err := Connect()
   if err != nil {
	   panic(err)
   }

//    // melakukan query untuk mendapatkan semua data dari tabel employee
//    rows, err := db.Query("SELECT * FROM employee")
//    if err != nil {
// 	   panic(err)
//    }

//    // membuat struct baru untuk menampung data dari tabel employee
//    var listEmployee []Employee

//    // melakukan looping untuk menampung data dari rows ke struct Employee
//    for rows.Next() {
// 	   var employee Employee

// 	   // kita tampung setiap baris data ke struct Employee
// 	   err := rows.Scan(&employee.ID, &employee.Name, &employee.Age, &employee.Address, &employee.Salary)
// 	   if err != nil {
// 		   panic(err)
// 	   }

// 	   // kemudian kita tambahkan struct Employee ke listEmployee
// 	   listEmployee = append(listEmployee, employee)
//    }

//    fmt.Println("Berhasil melakukan query tabel employee")
//    fmt.Println(listEmployee)

  // melakukan query untuk mendapatkan data dari tabel employee dengan id 10
  row := db.QueryRow("SELECT * FROM employee WHERE id = 10")

  // membuat struct baru untuk menampung data dari tabel employee
  var employee Employee

  // kita tampung data ke struct Employee
  err = row.Scan(&employee.ID, &employee.Name, &employee.Age, &employee.Address, &employee.Salary)

  // jika errornya karena tidak menemukan data, maka kita tampilkan pesan
  if errors.Is(err, sql.ErrNoRows) {
	  fmt.Println("Data tidak ditemukan")
	  return
  } else if err != nil {
	  // jika errornya karena kesalahan lain, maka kita panic
	  panic(err)
  }

  fmt.Println("Berhasil melakukan query tabel employee")
  fmt.Println(employee)
}