package main

/*
 * Whoper AuthServer
 */
import (
	"log"
	"net/http"
)

func main() {
	// Start the router
	router := NewRouter()
	// Load the SignKey
	initKeys(&SignKey)
	log.Println("SignKey loaded")
	// listening on port
	log.Println("Whoper Auth API now listening on port:\t5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}

/*
 * END OF FILE!
 */
