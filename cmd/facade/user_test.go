package main

import (
  "testing"
  "github.com/markbates/goth"
)

func TestUserCopyValuesFromGothUser(t *testing.T) {
  gothUser := goth.User{
  	NickName: "TestUser",
  	Name: "Test User",
    FirstName: "Firstname",
    LastName: "Lastname",
    AvatarURL: "http://example.com/newavatar.png",
  }

  user := &User{
  	"testUserId",
  	"OldNickName",
  	"OldName",
  	"OldEmail@example.com",
  	"http://example.com/oldavatar.png",
  }

  user.CopyValuesFrom(gothUser)

  if user.Id != "testUserId" {
  	t.Errorf("Id was changed, expected '%s' got '%s'.", "testUserId", user.Id)
  }

  if user.NickName != gothUser.NickName {
  	t.Errorf("user.NickName != gothUser.NickName, expected '%s' got '%s'", gothUser.NickName, user.NickName)
  }

  if user.Name != gothUser.Name {
	  t.Errorf("user.Name != gothUser.Name, expected '%s' got '%s'", gothUser.Name, user.Name)
  }

  if user.AvatarURL != gothUser.AvatarURL {
    t.Errorf("Expected '%s' got '%s'", gothUser.AvatarURL, user.AvatarURL)
  }

  gothUser = goth.User{
    FirstName: "Firstname",
    LastName: "Lastname",
  }

  user.CopyValuesFrom(gothUser)

  expectedName := gothUser.FirstName + " " + gothUser.LastName
  if user.Name != expectedName {
	  t.Errorf("Expected '%s' got '%s'", expectedName, user.Name)
  }
}
