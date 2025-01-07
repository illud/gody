package env

import (
	"os"

	"github.com/joho/godotenv"
)

// Add here your .env data
type Env struct {
	VERSION string
	NAME    string
	IP      string
	PORT    string
	URL     string
}

func Load() Env {
	godotenv.Load() //This loads your .env

	//Returns .env data int Env struct
	return Env{
		VERSION: os.Getenv("VERSION"),
		NAME:    os.Getenv("NAME"),
		IP:      os.Getenv("IP"),
		PORT:    os.Getenv("PORT"),
		URL:     os.Getenv("URL"),
	}
}
