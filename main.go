package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var (
		endpoint, username, password, getCookie string
	)

	endpoint, username, password, getCookie = os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	fmt.Printf("Logging into %s with user/pass %s/%s and retrieving header %s\n", endpoint, username, password, getCookie)
	_, cookie := basicAuth(endpoint, username, password, getCookie)
	fmt.Printf("got cookie: %s\n", cookie)
}

func basicAuth(endpoint string, username string, password string, getCookie string) (string, string) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(fmt.Sprintf("email=%s&password=%s", username, password)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Accept", "application/json")

	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := string(bodyText)

	cookie := cookieByName(resp, getCookie)
	return body, cookie
}

func cookieByName(resp *http.Response, name string) string {
	for _, c := range resp.Cookies() {
		if c.Name == name {
			return c.Value
		}
	}
	return ""
}
