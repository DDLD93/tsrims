package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlCon() *sql.DB {
	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	createMigrations(db)
	return db
}

func createMigrations(conn *sql.DB) {
	log.Println("Creating migrations ...")
	customerTableQuery := `CREATE TABLE IF NOT EXISTS customers(
		id int PRIMARY KEY auto_increment, 
		accountNo varchar(12) NOT NULL,
		firstName varchar(50) NOT NULL,
		lastName varchar(50) NOT NULL,
		dob varchar(12) NOT NULL,
		sex varchar(6) NOT NULL,
		address varchar(200) NOT NULL,
		phone varchar(12) NOT NULL,
		bvn varchar(12) NOT NULL,
		pin varchar(4) NOT NULL
		)`
	result, err := conn.Exec(customerTableQuery)
	if err != nil {
		log.Printf("Error Creating Customer table, %v", err)
	}
	fmt.Println(result)
}
