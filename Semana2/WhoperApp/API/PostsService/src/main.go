package main

/*
 * Whoper Posts Servive API
 */
import (
	"log"
	"net/http"
)

func main() {
	// Start the router
	router := NewRouter()
	// listening on port
	log.Println("Whoper Posts RestFul API now listening on port:\t3002")
	log.Fatal(http.ListenAndServe(":3002", router))
}

/*
 * END OF FILE!
 */
