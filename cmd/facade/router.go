package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/markbates/goth/gothic"
)

func BuildRouter() (*pat.Router) {
  r := pat.New()
	r.Get("/auth/{provider}/callback", authCallbackHandler)
	r.Get("/logout/{provider}", logoutHandler)
	r.Get("/auth/{provider}", authHandler)
  r.Get("/me", meHandler)

  return r
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
