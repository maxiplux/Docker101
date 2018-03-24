package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
 * FollowHandler user follow user
 */
func FollowHandler(w http.ResponseWriter, r *http.Request) {

	db := ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	followID, err := strconv.Atoi(vars["follow_id"])

	if err != nil {
		log.Printf("Not a Valid user id	%v", err)
	}
	db.QueryRow("INSERT INTO following (user_id,follows_user_id) VALUES ($1,$2)", userID, followID)
	if err != nil {
		log.Printf("Func\tFollowHandler:\tcould not follow user %d:\t%v", followID, err)
	} else {
		AddUsertoGroup(userID, followID, Publicgroup)
		log.Printf("Func\tFollowHandler:\tUser %d Succesfuly follows user with id %d from db", userID, followID)
	}
}

/*
 * END OF FILE!
 */
