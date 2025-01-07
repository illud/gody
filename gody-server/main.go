package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	//Uncomment next line when you want to connect to a database
	db "github.com/gody-server/adapters/database"
	router "github.com/gody-server/router"
)

type ConfigFile struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port string `json:"port"`
	Url  string `json:"url"`
}

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
	// Open the JSON file
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Read the file contents into a byte slice
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Unmarshal the JSON data into a Person struct
	var configFile ConfigFile
	err = json.Unmarshal(fileBytes, &configFile)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	//Uncomment next line when you want to connect to a database
	//Connect to database
	db.Connect()

	err = db.Migrate()
	if err != nil {
		fmt.Println("DATABASE MIGRATE ERROR: ", err)
	}

	//Load .env port
	// port := "5000"

	// if port == "" {
	// 	fmt.Println("$PORT must be set")
	// }

	router.Router().Run(configFile.IP + ":" + configFile.Port)
}
