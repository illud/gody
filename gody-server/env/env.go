package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

//Add here your .env data
type Env struct {
	PORT                   int
	VERSION                string
}

func Load() Env {
	godotenv.Load() //This loads your .env

	//Converts port string to int
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	//Returns .env data int Env struct
	return Env{
		PORT:                   port,
		VERSION:      os.Getenv("VERSION"),
	}
}