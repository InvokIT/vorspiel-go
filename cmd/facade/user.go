package main

import (
	"fmt"
	"strings"
	"github.com/markbates/goth"
)

type User struct {
	Id string
	NickName string
	Name string
	Email string
	AvatarURL string
}

func (user *User) CopyValuesFrom(s goth.User) {
	if s.NickName != "" {
		user.NickName = s.NickName
	}

	switch {
	case s.Name != "":
		user.Name = s.Name
	case s.FirstName != "" || s.LastName != "":
		if n := strings.TrimSpace(fmt.Sprintf("%s %s", s.FirstName, s.LastName)); n != "" {
			user.Name = n
		}
	}

	if s.Email != "" {
		user.Email = s.Email
	}

	if s.AvatarURL != "" {
		user.AvatarURL = s.AvatarURL
	}
}
