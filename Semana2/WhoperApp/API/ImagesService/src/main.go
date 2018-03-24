package main

/*
 * Whoper Images Servive API
 */
import (
	"log"
	"net/http"
)

func main() {
	// Start the router
	router := NewRouter()
	// listening on port
	log.Println("Whoper Images RestFul API now listening on port:\t3004")
	log.Fatal(http.ListenAndServe(":3004", router))
}

/*
 * END OF FILE!
 */
