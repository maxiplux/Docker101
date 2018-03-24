package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
 * ListUsersHandler displays a list with summaries of users in the database.
 */
func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Fprintf(w, "could not retrieve users: %v", err)
	}
	var users []User
	defer rows.Close()
	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.Id, &user.User_name, &user.User_login, &user.Pwd, &user.Email, &user.Gender,
			&user.Location, &user.Salt, &user.Status, &user.Image)
		if err != nil {
			log.Print(err.Error())
			json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to select user from database"})
		} else {
			users = append(users, user)
		}
	}
	json.NewEncoder(w).Encode(users)
	log.Printf("Func:\tListUsersHandler:\tSuccesfuly retrieved users from db")
}

/*
 * GetUsersHandler retrieves a user info from its user id
 */
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Printf("Not a Valid id	%v", err)
	}
	var user User
	err = db.QueryRow("SELECT * FROM users WHERE id = $1", userID).Scan(
		&user.Id, &user.User_name, &user.User_login, &user.Pwd, &user.Email,
		&user.Gender, &user.Location, &user.Salt, &user.Status, &user.Image)
	if err != nil {
		log.Printf("Func\tGetUsersHandler:\tcould not find user:\t%v", err)
	} else {
		json.NewEncoder(w).Encode(user)
		log.Printf("Func\tGetUsersHandler:\tSuccesfuly retrieved user with id\t%v from db", user.Id)
	}

}

/*
 * DeleteUserHandler set status to 0 from its user id
 */
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	key := r.FormValue("salt")

	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
	}

	stmt, err := db.Prepare("UPDATE users SET status = '0' WHERE id = $1 AND salt = $2")
	_, err = stmt.Exec(userID, key)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Func\tDeleteUserHandler:\tSuccesfuly updated user: %d status to 0 (off)", userID)
	}
}

/*
 * END OF FILE!
 */
