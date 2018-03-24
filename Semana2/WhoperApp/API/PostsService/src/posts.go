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
 * CreatePostHandler adds a post to the database.
 */
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	post := postFromForm(r)

	db := ConnectDB()
	defer db.Close()

	PostInsertQuery := "INSERT INTO posts (rating, price, location, caption, status, post_owner) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id, post_date"
	err := db.QueryRow(PostInsertQuery,
		post.Rating, post.Price, post.Location, post.Caption, post.Status, post.Post_owner).Scan(&post.Id, &post.Post_date)

	if err != nil {
		log.Printf("Func\tCrearePostHandler:\tError adding user:\t%v", err)
	} else {
		json.NewEncoder(w).Encode(post)
		log.Printf("Func\tCrearePostHandler:\tSuccesfuly added post with id :\t%d to DataBase", post.Id)
	}
	postOnGroup(post.Post_owner, post.Id, r.FormValue("group"))
}

/*
 * GetPostHandler retrieves a post info from its post id
 */
func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
	}

	var post Post
	err = db.QueryRow("SELECT * FROM posts WHERE id = $1", postID).Scan(
		&post.Id, &post.Post_date, &post.Rating, &post.Price, &post.Location,
		&post.Caption, &post.Status, &post.Post_owner)
	if err != nil {
		fmt.Fprintf(w, "could not find post: %v \n", err)
	} else {
		json.NewEncoder(w).Encode(post)
		log.Printf("Func:\tGetPostHandler:\tSuccesfuly retrieved post with id\t%v from db", post.Id)
	}
}

/*
 * DeletePostHandler set status to 0 from its user id
 */
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["post_id"])

	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
	}

	ImageDeleteQuery := "UPDATE posts SET status=0 WHERE id=$1"
	db.QueryRow(ImageDeleteQuery,
		postID)

	if err != nil {
		log.Printf("Func\tImagePostHandler:\tError Deleting image:\t%v", err)
	} else {
		json.NewEncoder(w).Encode(nil)
		log.Printf("Func\tImagePostHandler:\tSuccesfuly DELETE post\t%d status set to 0", postID)
	}
}

/*
 * postOnGroup asing a post to a group
 */
func postOnGroup(owner int, postID int, groupName string) {
	db := ConnectDB()
	defer db.Close()
	// Retrieve group_id
	rows, err := db.Query("SELECT id FROM groups WHERE admin=$1 AND group_name=$2", owner, groupName)
	if err != nil {
		log.Printf("Func\tpostOnGroup:\tcould not retrieve group_id:\t%v", err)
	}
	// asing post to group
	var groupID int
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&groupID)
	}
	db.QueryRow(`INSERT INTO posts_groups (post_id, group_id)
		VALUES($1, $2)`, postID, groupID)
}

/*
 * postFromForm populates the fields of a Post from form values
 */
func postFromForm(r *http.Request) Post {
	ItemPrice, err := strconv.ParseInt(r.FormValue("price"), 10, 64)
	if err != nil {
		log.Printf("Error on operation Parseint for %d : %v", ItemPrice, err)
	}

	ownerID, err := strconv.Atoi(r.FormValue("post_owner"))
	if err != nil {
		log.Printf("Error on operation Strng-int: %v", err)
	}

	post := Post{
		Id:         1,
		Rating:     r.FormValue("rating"),
		Price:      ItemPrice,
		Location:   r.FormValue("location"),
		Caption:    r.FormValue("caption"),
		Status:     1,
		Post_owner: ownerID,
	}

	return post
}

/*
 * END OF FILE!
 */
