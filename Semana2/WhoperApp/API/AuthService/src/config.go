package main

import (
	"io/ioutil"
	"log"
)

/*
 * RSA KEYS AND INITIALISATION
 */

const (
	privKeyPath = "/etc/keys/app.rsa"
)

func initKeys(SignKey *[]byte) {
	var err error
	*SignKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
	}
}

/*
 * GetDBConfig returns a DBconfig struct with the data base credentials
 */
func GetDBConfig() *DBconfig {
	dbconfig := &DBconfig{
		DB_USER:     "whoperadmin",
		DB_PASSWORD: "montanoarango",
		DB_NAME:     "whoperdb",
		SSL_MODE:    "disable",
		HOST:        "postgres",
	}

	return dbconfig
}
