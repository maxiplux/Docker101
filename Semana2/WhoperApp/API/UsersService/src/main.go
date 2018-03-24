package main

/*
 * Whoper API SERVER
 */
import (
	"log"
	"net/http"
)

func main() {
	// Start the router
	router := NewRouter()
	// listening on port
	log.Println("Whoper RestFul API now listening on port:\t3001")
	log.Fatal(http.ListenAndServe(":3001", router))
}

/*
 * END OF FILE!
 */
