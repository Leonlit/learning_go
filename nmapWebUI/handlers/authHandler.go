package handlers

import (
	"net/http"
)

func GetLoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/pages/login.html")
}

func GetRegistraterPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/pages/register.html")
}

func VerifyUserLogin(w http.ResponseWriter, r *http.Request) {
	// Fetch host by ID
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Fetch host by ID
}
