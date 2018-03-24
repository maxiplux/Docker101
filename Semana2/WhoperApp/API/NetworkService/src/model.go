package main

import "net/http"

const (
	Publicgroup = "Public"
)

/*
 * Struct for User
 */
type User struct {
	Id         int    `json:"id"`
	User_name  string `json:"user_name"`
	User_login string `json:"user_login"`
	Pwd        string `json:"pwd"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	Location   string `json:"location"`
	Salt       string `json:"salt"`
	Status     int8   `json:"estado"`
	Image      string `json:"imageURL"`
}

/*
 * Struct for Posts
 */
type Post struct {
	Id         int    `json:"id"`
	Post_date  string `json:"post_date"`
	Rating     string `json:"post_rating"`
	Price      int64  `json:"price"`
	Location   string `json:"location"`
	Caption    string `json:"caption"`
	Status     int8   `json:"status"`
	Post_owner int    `json:"owner_id"`
}

/*
 * Struct for Groups
 */
type Group struct {
	Id          int    `json:"id"`
	Group_name  string `json:"group_name"`
	Status      int8   `json:"status"`
	Group_owner int    `json:"owner_id"`
	Image       string `json:"imageURL"`
}

/*
 * Struct for json responses
 */
type HttpResp struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

/*
 * Struct for Data base configuration
 */
type DBconfig struct {
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	SSL_MODE    string
	HOST        string
}

/*
 * Type Route
 */
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

/*
 * END OF FILE!
 */
