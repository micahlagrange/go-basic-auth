package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Implement a form login
func main() {
	var (
		endpoint, username, password, getCookie string
	)

	endpoint, username, password, getCookie = os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	fmt.Printf("Logging into %s with user/pass %s/%s and retrieving header %s\n",
		endpoint, username, password, getCookie)

	cookie := basicAuth(endpoint, username, password, getCookie)

	fmt.Printf("got cookie: %s\n", cookie)
}

// Do basic auth, return a cookie value by name
func basicAuth(endpoint string, username string, password string, getCookie string) string {
	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(fmt.Sprintf("email=%s&password=%s", username, password)))
	// Add headers needed to post a form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	cookie := cookieByName(resp, getCookie)

	return cookie
}

// Take a string and return the value for the cookie whos name == string
func cookieByName(resp *http.Response, name string) string {
	for _, c := range resp.Cookies() {
		if c.Name == name {
			return c.Value
		}
	}
	return ""
}
