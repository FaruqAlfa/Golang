package main

import (
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
    gorm.Model
    Name  string
    Age int
}

func main() {
    dsn := "host=localhost user=postgres password=postgres dbname=test_db_camp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
		  TablePrefix: "t_", // awalan nama tabel `User` akan menjadi `t_users`
		   SingularTable: true, // gunakan nama tabel tunggal, tabel untuk `User` akan menjadi `user` saat opsi ini diaktifkan
		   NoLowerCase: true, // skip nama snake_casing
		   NameReplacer: strings.NewReplacer( "CID" , "Cid " ), // gunakan nama pengganti untuk mengubah nama struct/field sebelum mengubahnya menjadi nama db
		 },
	})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&User{})

    // Create
    db.Create(&User{Name: "Aditira", Age: 23})

    // Read
    // SELECT * FROM "users" WHERE "users"."id" = 1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1
    var user User
    db.First(&user, 1) // temukan user dengan menggunakan primary key dan simpan di variabel user
    db.First(&user, "name = ?", "Aditira") // temukan user dengan nama Aditira

    fmt.Println(user)
}