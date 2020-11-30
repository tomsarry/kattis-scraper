package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"github.com/tomsarry/kattis-scraper/utils"
)

func main() {
	// create a cookiejar to store cookies
	jar, _ := cookiejar.New(nil)

	app := utils.App{
		Client: &http.Client{Jar: jar},
	}

	app.Login()

	// get alphabetical order
	problems := app.GetProblems()

	// get increasing order of difficulty
	utils.GetIncreasing(problems)
	// utils.GetDecreasing(problems)

	for index, pb := range problems {
		fmt.Printf("%d: %s, %.1f, %s\n", index+1, pb.Name, pb.Difficulty, pb.Link)
	}
}
