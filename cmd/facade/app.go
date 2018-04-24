package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"os"

	//jwt "github.com/dgrijalva/jwt-go"
  "github.com/gorilla/pat"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/steam"
  "github.com/invokit/vorspiel-lib/mq"
)

import (
)

type App struct {
  Port uint16
  PublicUrl string
  MQ mq.Client
  Router *pat.Router
}

func (app *App) Start() {
  goth.UseProviders(buildSteamProvider(app.PublicUrl))

  address := fmt.Sprintf(":%s", app.Port)
	fmt.Printf("Listening on address %s\n", address)
	logger.Fatal(http.ListenAndServe(address, app.Router))
}

func buildSteamProvider(publicUrl string) *steam.Provider {
	steamKey := os.Getenv("STEAM_KEY")
	if steamKey == "" {
		logger.Panic("STEAM_KEY is not defined.")
	}

  return steam.New(steamKey, fmt.Sprintf("%s/auth/steam/callback", publicUrl))
}
