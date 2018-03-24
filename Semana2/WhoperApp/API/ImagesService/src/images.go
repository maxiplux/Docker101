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
 * ImageHandler adds a image to the database.
 */
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	image := imageFromForm(r)
	postID, _ := strconv.Atoi(r.FormValue("post_id"))
	db := ConnectDB()
	defer db.Close()

	ImageInsertQuery := "INSERT INTO images (url,status) VALUES ($1,$2) RETURNING id"
	err := db.QueryRow(ImageInsertQuery,
		image.URL, image.Status).Scan(&image.ID)

	if err != nil {
		log.Printf("Func\tCrearePostHandler:\tError adding user: %v", err)
	} else {
		json.NewEncoder(w).Encode(image.ID)
		log.Printf("Func\tCrearePostHandler:\tSuccesfuly added post with id : %d to DataBase", image.ID)
	}
	_ = ImagePost(postID, image.ID)
}

/*
 * ImagePost atach a image to a post info from its post id
 */
func ImagePost(postID int, imageID int) error {
	db := ConnectDB()
	defer db.Close()

	ImagePostInsertQuery := "INSERT INTO images_post (posts_id, image_id) VALUES ($1,$2)"
	db.QueryRow(ImagePostInsertQuery,
		postID, imageID)
	var err error
	if err != nil {
		log.Printf("Func\tImagePost:\tError adding image to post: %v", err)
	} else {
		log.Printf("Func\tImagePost:\tSuccesfuly added image\t%d to post\t%d", imageID, postID)
	}
	return err
}

/*
 * DeleteImageHandler set status to 0 from its image id
 */
func DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	imageID, err := strconv.Atoi(vars["image_id"])

	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
	}

	ImageDeleteQuery := "UPDATE images SET status=0 WHERE id=$1"
	db.QueryRow(ImageDeleteQuery,
		imageID)

	if err != nil {
		log.Printf("Func\tImagePostHandler:\tError Deleting image:\t%v", err)
	} else {
		json.NewEncoder(w).Encode(nil)
		log.Printf("Func\tImagePostHandler:\tSuccesfuly DELETE image\t%d status set to 0", imageID)
	}
}

/*
 * imageFromForm populates the fields of a Image from form values
 */
func imageFromForm(r *http.Request) Image {
	image := Image{
		ID:     1,
		URL:    r.FormValue("url"),
		Status: 1,
	}
	return image
}

/*
 * END OF FILE!
 */
