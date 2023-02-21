package api

import (
	"fmt"
	"log"
	"os"

	"github.com/QuocBao92/go-sample/api/controllers"
	"github.com/QuocBao92/go-sample/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	var err error
	// Load env
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	// connect DB
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// mockup some data to test
	seed.Load(server.DB)

	server.Run(":8090")

}
