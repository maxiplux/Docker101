package main

import (
	"net/http"
	"net/url"
)

/*
 * Gobal const
 */
const (
	consumersURL = "http://kong:8001/consumers/"
	pluginsURL   = "http://kong:8001/apis/"
	groupsUrl    = "http://ntwrs:3003/groups"
)

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
 * Signature key
 */
var SignKey []byte

/*
 * Struct for AuthSession
 */
type AuthSession struct {
	Id         int    `json:"id"`
	User_login string `json:"user_login"`
	Salt       string `json:"salt"`
	Token      string `json:"token"`
}

/*
 * Struct for Kong consumers
 */
type Consumer struct {
	Username  string
	Custom_id string
}

/*
 * Struct for Kong API key-auth
 */
type APIKey struct {
	ConsumerId string
	CreatedAt  int
	ID         string
	Key        string
}

/*
 * Struct for Kong plugin
 */
type Plugins struct {
	Plugin string
	values url.Values
}

/*
 * END OF FILE!
 */
