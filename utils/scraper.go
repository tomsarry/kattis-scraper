package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

// App stores the client info
type App struct {
	Client *http.Client
}

// AuthenticityToken stores the token for login
type AuthenticityToken struct {
	Token string
}

// Problem holds info for each problem solved
type Problem struct {
	Name       string
	Difficulty float64
	Link       string
}

// iniatialize global variables
func init() {
	godotenv.Load(".env")
	username = os.Getenv("EMAIL")
	password = os.Getenv("PASSWORD")

	// verify that found values for login
	if username == "" || password == "" {
		fmt.Println("Error: could not load username and/or password.")
		os.Exit(1)
	}
}

const baseURL = "https://open.kattis.com"

var (
	username = ""
	password = ""
)

// getToken retrieves the token for login
func (app *App) getToken() AuthenticityToken {
	loginURL := baseURL + "/login/email?"
	client := app.Client

	response, err := client.Get(loginURL)

	if err != nil {
		log.Fatalln("Error fetching response. ", err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// find the hidden input for token
	token, find := document.Find("input[name='csrf_token']").Attr("value")

	if find == false {
		log.Fatal("Did not find input field.")
	}

	authenticityToken := AuthenticityToken{
		Token: token,
	}

	return authenticityToken
}

// Login logs in on Kattis
func (app *App) Login() {
	client := app.Client

	authenticityToken := app.getToken()

	loginURL := baseURL + "/login/email?"

	data := url.Values{
		"csrf_token": {authenticityToken.Token},
		"user":       {username},
		"password":   {password},
	}

	response, err := client.PostForm(loginURL, data)

	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
}

// GET NEXT PAGE (url page=1...)
// GetProblems returns the list of solved problems
func (app *App) GetProblems() []Problem {
	// get only solved problems
	projectsURL := baseURL + "/problems?show_solved=on&show_tried=off&show_untried=off"

	client := app.Client

	response, err := client.Get(projectsURL)

	if err != nil {
		log.Fatalln("Error fetching response. ", err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	var problems []Problem

	// seek login button (unique element with class .dark-bg)
	document.Find(".dark-bg").Each(func(i int, s *goquery.Selection) {
		fmt.Println("Error: Could not login to Kattis.")
		fmt.Println("HINT: Check that you have the correct username/password in your .env file.")
		fmt.Println("HINT: Also check to see if website is down ?")
		os.Exit(1)
	})

	// get all solved problems on the problem page
	document.Find(".solved").Each(func(i int, s *goquery.Selection) {
		// not selecting first one (sorting button)
		if i == 0 {
			return
		}

		name := s.Find(".name_column").Text()

		difficulty, err := strconv.ParseFloat(s.Find(".numeric").Last().Text(), 64)
		if err != nil {
			// TODO: check if error if difficulty is weird : 1.7-7.5 for ex -> should fail
			fmt.Println("Error, could not parse the difficulty.")

			// set it to 0 if error
			difficulty = 0
		}

		// find link to problem
		link, ok := s.Find("a").Attr("href")

		if ok == false {
			fmt.Println("Error: Could not find URL of problem.")
			link = ""
		}

		link = baseURL + link

		problem := Problem{
			Name:       name,
			Difficulty: difficulty,
			Link:       link,
		}

		problems = append(problems, problem)
	})

	return problems
}
