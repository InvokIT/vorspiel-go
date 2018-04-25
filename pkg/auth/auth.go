package auth

import (
  "errors"
  "net/http"
  "fmt"
  "os"
  "log"
  "github.com/markbates/goth"
  "github.com/markbates/goth/providers/steam"
  "github.com/gorilla/pat"
)

var PublicUrl string
var TokenWriter func(token string, http.ResponseWriter)
var TokenReader func(*http.Request) (token string)
var UserToTokenConverter func(user *User) token string
var TokenToUserConverter func(token string) *User

var logger = log.New(os.Stderr, "", log.LstdFlags | log.Llongfile)

func SetupRouter(router *pat.Router) {
  router.Get("/auth/{provider}/callback", authCallbackHandler)
	router.Get("/logout/{provider}", logoutHandler)
	router.Get("/auth/{provider}", authHandler)
  router.Get("/me", meHandler)
}

func init() {
  providerBuilders := make([]providerBuilder, 0)

  providerBuilders = append(providerBuilders, buildSteamProvider)
  // Append new providers here

  providers := make([]goth.Provider, len(providerBuilders))
  for i, b := range providerBuilders {
    p, err := b(publicUrl)
    if err != nil {
      return err
    }

    providers[i] = p
  }

  goth.UseProviders(...providers)
}

type providerBuilder func(publicUrl string) (goth.Provider, error)

func buildSteamProvider(publicUrl string) (goth.Provider, error) {
  steamKey, ok := os.LookupEnv("STEAM_KEY")
	if !ok {
    return nil, errors.New("STEAM_KEY is not defined in the environment")
	}

  provider := steam.New(steamKey, fmt.Sprintf("%s/auth/steam/callback", publicUrl))

  return provider, nil
}

func authCallbackHandler(res http.ResponseWriter, req *http.Request) {
	logger.Print("AuthCallbackHandler entered")

	gothUser, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		logger.Print(err)
		fmt.Fprint(res, err)
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	logger.Printf("AuthCallbackHandler got user: %s", gothUser)

	if err := Templates["authSuccess"].Execute(res, gothUser); err != nil {
		logger.Print(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	//res.WriteHeader(http.StatusOK)
}

func logoutHandler(res http.ResponseWriter, req *http.Request) {
	logger.Print("LogoutHandler entered")

	gothic.Logout(res, req)
	//res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusOK)
}

func authHandler(res http.ResponseWriter, req *http.Request) {
	logger.Print("AuthHandler entered")

	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		if err := Templates["authSuccess"].Execute(res, gothUser); err != nil {
			logger.Fatal(err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	gothic.BeginAuthHandler(res, req)
}

func meHandler(res http.ResponseWriter, req *http.Request) {
  // TODO Get user from Request and write it to Response
}
