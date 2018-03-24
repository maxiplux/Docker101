package main

import "net/http"

/*
 * Struct for User
 */
type Image struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	Status int    `json:"status"`
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
