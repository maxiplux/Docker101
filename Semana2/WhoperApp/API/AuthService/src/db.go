package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

/*
 * ConnectDB conects to the data base
 */
func ConnectDB() *sql.DB {
	dbinfo := GetDBConfig()

	dbconfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s ",
		dbinfo.DB_USER, dbinfo.DB_PASSWORD, dbinfo.DB_NAME, dbinfo.SSL_MODE, dbinfo.HOST)

	db, err := sql.Open("postgres", dbconfig)
	if err != nil {
		log.Printf("Func\tConnectDB:\tCouldn't open Data base\n %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Printf("Func\tConnectDB:\tCouldn't conect to Data base:\n %v ", err)
	}
	return db
}

/*
 * END OF FILE!
 */
