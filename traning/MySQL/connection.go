package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connection_Link() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/MySqlDB")

	if err != nil {
		fmt.Println(err, "Connection error")
	} else {
		fmt.Println(db, "connection success")
	}
	return db
}

func CreateTable(db *sql.DB) {
	query := `create table table3(id int primary key auto_increment, name varchar(30))`
	ctx, canclefun := context.WithTimeout(context.Background(), 5*time.Second)
	defer canclefun()
	resp, err := db.ExecContext(ctx, query)
	if err != nil {
		fmt.Println(err, "query error ")
	}
	row, _ := resp.RowsAffected()
	fmt.Println(" Table Created ", row)
}

func main() {
	//Connection_Link()

	db := Connection_Link()
	CreateTable(db)

}
