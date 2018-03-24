package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
 * CreateGroupHandler adds a post to the database.
 */
func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	group := groupFromForm(r)

	db := ConnectDB()
	defer db.Close()
	// INSERT into groups
	GroupInsertQuery := "INSERT INTO groups (group_name, status, admin, imagen) VALUES ($1,$2,$3,$4) RETURNING id"
	err := db.QueryRow(GroupInsertQuery,
		group.Group_name, group.Status, group.Group_owner, group.Image).Scan(&group.Id)
	// Error handling and return
	if err != nil {
		log.Printf("func\tCreateGroupHandler:\tError creating group:\t%v", err)
	} else {
		json.NewEncoder(w).Encode(group.Id)
		log.Printf("func\tCreateGroupHandler:\tSuccesfuly added group:\t{%s}\tid\t%d to DataBase", group.Group_name, group.Id)
	}
}

/*
 * AddUsertoGroup adds a user to a group.
 */
func AddUsertoGroup(adminID int, userID int, groupID int) {
	if groupID == 0 {
		groupName := Publicgroup
	}
	db := ConnectDB()
	defer db.Close()
	// Retrieve group_id
	rows, err := db.Query("SELECT id FROM groups WHERE admin=$1 AND group_name=$2", adminID, groupName)
	if err != nil {
		log.Printf("Func\tAddUsertoGroup:\tcould not retrieve group_id:\t%v", err)
	}
	// insert user into group
	var groupID int
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&groupID)
	}
	log.Printf("id:\t%d\tname:\t%s", groupID, groupName)
	db.QueryRow(`INSERT INTO users_groups (user_id, group_id)
		VALUES($1, $2)`, userID, groupID)
}

/*
 * FollowHandler user follow user
 */
func AddUserHandler(w http.ResponseWriter, r *http.Request) {

	db := ConnectDB()
	defer db.Close()
	vars := mux.Vars(r)
	adminID, _ := strconv.Atoi(vars["admin_id"])
	userID, _ := strconv.Atoi(vars["user_id"])
	groupName, _ := strconv.Atoi(vars["group_id"])

	AddUsertoGroup(adminID, userID, groupID)
}

/*
 * groupFromForm populates the fields of a Group from form values
 */
func groupFromForm(r *http.Request) Group {
	imageURL := r.FormValue("imageURL")

	if imageURL == "" {
		imageURL = "https://upload.wikimedia.org/wikipedia/commons/b/b9/Group_font_awesome.svg"
	}

	ownerID, err := strconv.Atoi(r.FormValue("owner_id"))
	if err != nil {
		log.Printf("Func\tgroupFromForm:\tError on operation Strng-int:\t%v", err)
	}

	group := Group{
		Id:          1,
		Group_name:  r.FormValue("group_name"),
		Status:      1,
		Group_owner: ownerID,
		Image:       imageURL,
	}

	return group
}

/*
 * END OF FILE!
 */
