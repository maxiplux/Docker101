package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

/*
 * CreareUserHandler adds a user to the database.
 */
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := userFromForm(r)

	db := ConnectDB()
	defer db.Close()
	// INSERT user into DB
	UserInsertQuery := "INSERT INTO users (user_name, user_login, pwd, email, gender, location, salt, status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id"
	err := db.QueryRow(UserInsertQuery,
		user.User_name, user.User_login, user.Pwd, user.Email, user.Gender, user.Location, user.Salt, user.Status).Scan(&user.Id)
	// Error handling and return
	if err != nil {
		log.Printf("func\tCreareUserHandler:\tError adding user:\t%v", err)
	} else {
		// Create the user Public Group
		_, err = http.PostForm(groupsUrl, url.Values{"group_name": {"Public"}, "owner_id": {strconv.Itoa(user.Id)}})
		// Create the user Kong Consumer
		_, err = CreateConsumer(Consumer{Username: user.User_login, Custom_id: user.Salt})
		json.NewEncoder(w).Encode(nil)
		log.Printf("func\tCreareUserHandler:\tSuccesfuly added user:\t{%s}\tid\t%d to DataBase", user.User_name, user.Id)
	}
}

/*
 * GetUsers retrieves a user info from its user_login
 */
func GetUsers(user_login string, as *AuthSession) (string, error) {
	db := ConnectDB()
	defer db.Close()

	var pswdhash string

	err := db.QueryRow("SELECT id,user_login,pwd,salt FROM users WHERE user_login = $1", user_login).Scan(
		&as.Id, &as.User_login, &pswdhash, &as.Salt)
	if err != nil {
		log.Printf("func\tGetUsers:\tcould not find user:\t%v", err)
		return "", err
	} else {
		log.Printf("func\tGetUsers:\tSuccesfuly retrieved user:\t%s", as.User_login)
		return pswdhash, nil
	}
}

/*
 * GetEmail retrieves a user info from its email
 */
func GetEmail(email string, as *AuthSession) (string, error) {
	db := ConnectDB()
	defer db.Close()

	var pswdhash string

	err := db.QueryRow("SELECT id,user_login,pwd,salt FROM users WHERE email = $1", email).Scan(
		&as.Id, &as.User_login, &pswdhash, &as.Salt)
	if err != nil {
		log.Printf("func\tGetEmail:\tcould not find email:\t%v", err)
		return "", err
	} else {
		log.Printf("func\tGetUsers:\tSuccesfuly retrieved user:\t%s", as.User_login)
		return pswdhash, nil
	}
}

/*
 * userFromForm populates the fields of a User from form values
 */
func userFromForm(r *http.Request) User {
	imageURL := r.FormValue("imageURL")

	if imageURL == "" {
		imageURL = "https://upload.wikimedia.org/wikipedia/commons/6/67/User_Avatar.png"
	}

	passwordhash, _ := HashPassword(r.FormValue("pwd"))

	user := User{
		Id:         1,
		User_name:  r.FormValue("user_name"),
		User_login: r.FormValue("user_login"),
		Pwd:        passwordhash,
		Email:      r.FormValue("email"),
		Gender:     r.FormValue("gender"),
		Location:   r.FormValue("location"),
		Salt:       GenerateRandomString(64),
		Status:     1,
		Image:      imageURL,
	}

	return user
}

/*
 * END OF FILE!
 */
