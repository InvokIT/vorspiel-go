package main

import (
	//"encoding/json"
	//"fmt"
	"os"

	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

const defaultPublicUrl = "http://localhost:8080"

func main() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		logger.Printf("error when loading .env: %s", err)
	}

	publicUrl := defaultPublicUrl
	if v, ok := os.LookupEnv("PUBLIC_URL"); ok {
		publicUrl = v
	}

	router := BuildRouter()
	mq, err := BuildMq()
	if err != nil {
		logger.Fatalf("error when building mq client: %s", err)
	}

	app := &App{Port: 8080, PublicUrl: publicUrl, MQ: mq, Router: router}

	app.Start()
}
