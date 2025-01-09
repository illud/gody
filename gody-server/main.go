package main

import (
	"fmt"
	"log"

	//Uncomment next line when you want to connect to a database
	db "github.com/gody-server/adapters/database"
	configFile "github.com/gody-server/config"
	router "github.com/gody-server/router"
)

//The next lines are for swagger docs
// @title gody-server
// @version version(1.0)
// @description Description of specifications
// @Precautions when using termsOfService specifications

// @host localhost:5000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	//Uncomment next line when you want to connect to a database
	//Connect to database
	db.Connect()

	err := db.Migrate()
	if err != nil {
		fmt.Println("DATABASE MIGRATE ERROR: ", err)
	}

	//Load .env port
	// port := "5000"

	// if port == "" {
	// 	fmt.Println("$PORT must be set")
	// }

	configFile, err := configFile.ConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	router.Router().Run(configFile.IP + ":" + configFile.Port)
}
