package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "root"
	password = "secret"
	dbname   = "postgres"
	sslmode  = "disable"
)

func connect_db() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", psqlconn)
	checkError(err)
	// defer db.Close()
	err = db.Ping()
	checkError(err)
	fmt.Println("Connected to database!")
	return db
}

func checkError(err error) {
	if err != nil {
		log.Printf(err.Error())
	}
}

func Service_get_db() {
	db := connect_db()
	rows, err := db.Query("SELECT id, feature_id FROM Earthquakes")
	checkError(err)
	fmt.Println(rows)
	for rows.Next() {
		var id int
		var feature_id string
		err_Scan := rows.Scan(&id, &feature_id)
		checkError(err_Scan)
		fmt.Printf("ID: %d, Feature ID: %s", id, feature_id)
	}
}
