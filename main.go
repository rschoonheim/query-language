package main

import "query-language/database"

func main() {

	configuration := database.ConfigurationNew()
	db := database.New(configuration)

	println(db)

}
