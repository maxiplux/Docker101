package main

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
