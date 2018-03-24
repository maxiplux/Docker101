package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

/*
 * Declare the routes and their handlers
 */
var routes = Routes{
	// User Routes	(Basic User handling calls)
	Route{"CreareUser", "POST", "/users", CreateUser},
	// ValidServiceProvider
	Route{"ValidateUserName", "GET", "/validate-username/{username}", ValidateUserName},
	// AuthServiceProvider
	Route{"UserLogin", "POST", "/validate-username", UserLogin},
	Route{"EmailLogin", "POST", "/validate-email", EmailLogin},
	// API server healt Routes
	Route{"APIhealt", "GET", "/healt", APIhealt},
}

/*
 * APIhealt check if AuthServer is up
 */
func APIhealt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Auth API server is UP!")
	dbinfo := GetDBConfig()
	dbconfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s ",
		dbinfo.DB_USER, dbinfo.DB_PASSWORD, dbinfo.DB_NAME, dbinfo.SSL_MODE, dbinfo.HOST)
	db, err := sql.Open("postgres", dbconfig)
	defer db.Close()
	if err = db.Ping(); err != nil {
		fmt.Fprintln(w, "Couldn't conect to Data base : ", err)
		log.Printf("Func\tAPIHealt:\tCouldn't conect to Data base: %v ", err)
	} else if err == nil {
		fmt.Fprintln(w, "Data base engine is UP and connected to Auth API server!")
		log.Printf("Func\tAPIHealt:\tHealt check passed")
	}
}

/*
 * END OF FILE!
 */
