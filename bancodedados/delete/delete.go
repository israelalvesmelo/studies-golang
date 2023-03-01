package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		panic(err)
	}
	stmt.Exec(3)

}
