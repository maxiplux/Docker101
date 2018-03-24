package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

/*
 * ValidateUserNameHandler
 */
func ValidateUserName(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	userName := vars["username"]

	if userName == "" {
		log.Printf("Func\tValidateUserNameHandler:\tNo user_login specified on request _ null search\t%s\n", userName)
		json.NewEncoder(w).Encode(nil)
	}

	var status int8
	err := db.QueryRow("SELECT status FROM users WHERE user_login = $1", userName).Scan(
		&status)
	if err != nil {
		log.Printf("Func\tValidateUserNameHandler:\tcould not find user_login:%s\t%vs\n", userName, err)
		json.NewEncoder(w).Encode(nil)
	} else {
		json.NewEncoder(w).Encode(status)
		log.Printf("Func\tValidateUserNameHandler:\tSuccesfuly retrieved user status with user_login\t%v\tfrom db", userName)
	}

}

/*
 * UserLogin log a user with user_name and return a login session token
 */
func UserLogin(w http.ResponseWriter, r *http.Request) {

	var as AuthSession

	hash, _ := GetUsers(r.FormValue("user_login"), &as)
	credentials := CheckPasswordHash(r.FormValue("pwd"), hash)

	if credentials == true {
		as.Token, _ = CreateAPIKey(Consumer{Username: as.User_login, Custom_id: as.Salt})
		log.Printf("func\tUserLogin:\tOn Login Credentials authentication for user:\t%s", r.FormValue("user_login"))
		json.NewEncoder(w).Encode(as)
	} else {
		json.NewEncoder(w).Encode(nil)
		log.Printf("func\tUserLogin:\tWrong credentials authentication for user:\t%s", r.FormValue("user_login"))
	}
}

/*
 * EmailLogin log a user with email and return a login session token
 */
func EmailLogin(w http.ResponseWriter, r *http.Request) {

	var as AuthSession

	hash, _ := GetEmail(r.FormValue("email"), &as)
	credentials := CheckPasswordHash(r.FormValue("pwd"), hash)

	if credentials == true {
		as.Token = string(SignKey)
		log.Printf("func\tEmailLogin:\tOn Login Credentials authentication for email:\t%s", r.FormValue("email"))
		json.NewEncoder(w).Encode(as)
	} else {
		json.NewEncoder(w).Encode(nil)
		log.Printf("func\tEmailLogin:\tWrong credentials authentication for email:\t%s", r.FormValue("email"))
	}

}

/*
 * CheckPasswordHash compares a hashed password with its possible plaintext equivalent.
 * Returns nil on success, or an error on failure.
 */
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/*
 * END OF FILE!
 */
