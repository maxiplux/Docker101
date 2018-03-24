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
	// Groups Routes	(Basic Groups handling calls)
	Route{"CreateGroup", "POST", "/groups", CreateGroupHandler},
	Route{"CreateGroup", "POST", "/groups/{admin_id}/{group_id}/{user_id}", AddUserHandler},
	// Users network Routes	(Follow,GetFollowers,Unfollow)
	Route{"FollowHandler", "POST", "/ntw/{user_id}/{follow_id}", FollowHandler},
	// API server healt Routes
	Route{"APIhealt", "GET", "/healt", APIhealt},
}

/*
 * APIhealt check if API server is up
 */
func APIhealt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API server is UP!")
	dbinfo := GetDBConfig()
	dbconfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s ",
		dbinfo.DB_USER, dbinfo.DB_PASSWORD, dbinfo.DB_NAME, dbinfo.SSL_MODE, dbinfo.HOST)
	db, err := sql.Open("postgres", dbconfig)
	defer db.Close()
	if err = db.Ping(); err != nil {
		fmt.Fprintln(w, "Couldn't conect to Data base : ", err)
		log.Printf("Func\tAPIhealt:\tCouldn't conect to Data base: %v ", err)
	} else if err == nil {
		fmt.Fprintln(w, "Data base engine is UP and connected to API server!")
		log.Printf("Func\tAPIhealt:\tHealt check passed")
	}
}

/*
 * END OF FILE!
 */
