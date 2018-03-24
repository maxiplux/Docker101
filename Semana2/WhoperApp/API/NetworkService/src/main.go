package main

/*
 * Whoper Network Service API
 */
import (
	"log"
	"net/http"
)

func main() {
	// Start the router
	router := NewRouter()
	// listening on port
	log.Println("Whoper Newtwork RestFul API now listening on port:\t3003")
	log.Fatal(http.ListenAndServe(":3003", router))
}

/*
 * END OF FILE!
 */
