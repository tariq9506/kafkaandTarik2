package main

import (
	"log"
	"os"
	"tarik_consumer/kafkaandTarik2/router"

	"github.com/gin-contrib/pprof"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {

		log.Fatal("Error loading .env file -> ", err)
	}
	r := router.SetUpRouter()
	pprof.Register(r)
	r.Run(":" + os.Getenv("PORT"))
}
