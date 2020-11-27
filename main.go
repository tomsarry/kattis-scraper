package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"github.com/tomsarry/kattis-scraper/models"
)

func main() {
	// create a cookiejar to store cookies
	jar, _ := cookiejar.New(nil)

	app := models.App{
		Client: &http.Client{Jar: jar},
	}

	app.Login()

	problems := app.GetProblems()

	for index, pb := range problems {
		fmt.Printf("%d: %s, %.1f, %s\n", index+1, pb.Name, pb.Difficulty, pb.Link)
	}

}
