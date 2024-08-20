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

var PORT int = 3009

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
	bearerToken := string(strings.Split(r.Header.Get("Authorization"), " ")[1])
	if !validateToken(bearerToken) {
		fmt.Fprint(w, "Invalid token")
		return
	}
	// initiate shutdown
	out, err := exec.Command("ls").Output()
	if err != nil {
		fmt.Fprint(w, err)
		log.Println(err)
		return
	}
	fmt.Fprint(w, string(out))
}

func main() {
	godotenv.Load(".env")
	http.HandleFunc("/shutdown", shutdownHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
