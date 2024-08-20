package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func validateToken(token string) bool {
	localToken := string(os.Getenv("TOKEN"))
	// don't use authentication if TOKEN=none in .env
	if localToken == "none" {
		return true
	}
	return token == localToken
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	// validate token
	authHeader := r.Header.Get("Authorization")
	var bearerToken string
	if authHeader == "" { // check if auth header not received
		bearerToken = ""
	} else {
		bearerToken = string(strings.Split(r.Header.Get("Authorization"), " ")[1])
	}
	if !validateToken(bearerToken) {
		fmt.Fprint(w, "Invalid token")
		return
	}
	// initiate shutdown
	err := exec.Command("sudo", "shutdown", "-h", "now").Run()
	if err != nil {
		fmt.Fprint(w, err)
		log.Println(err)
		return
	}
	fmt.Fprint(w, "The system will shutdown now")
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	http.HandleFunc("/shutdown", shutdownHandler)

	log.Println("Listening on port 3009")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
