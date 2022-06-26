package configs

// We defined environmental elements

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	// Read env file and load that env
	err := godotenv.Overload()

	if err != nil {
		log.Fatalln("Error in .env")
	}
	// That returns MongoDB adress in .env
	mongoURI := os.Getenv("MONGOURI")
	return mongoURI
}
