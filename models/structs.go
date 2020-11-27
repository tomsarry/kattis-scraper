package models

import "net/http"

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
