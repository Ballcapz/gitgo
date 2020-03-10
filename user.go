package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	color "github.com/fatih/color"
)

const (
	apiURL       = "https://api.github.com"
	userEndpoint = "/users/"
)

// User from github
type User struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	URL       string `json:"url"`
	ReposURL  string `json:"repos_url"`
}

func getUsers(name string) User {
	resp, err := http.Get(apiURL + userEndpoint + name)
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	var user User
	json.Unmarshal(body, &user)

	if (User{}) == user {
		color.Red("Sorry the GitHub user dosen't exist")
		os.Exit(1)
		return user
	}

	return user

}
