package main

import (
	"google.golang.org/api/people/v1"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

func getClient(config *oauth2.Config) *http.Client {
	tokenFile := "token.json"
	tok, err := tokenFromFile(tokenFile)

	if err != nil {
		tok = getTokenFromNet(config)
		saveToken(tokFile, tok)
	}

	return config.Client(context.Background(), tok)
}

func getTokenFromNet(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the  authorization  code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)

	f, err := os.Openfile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {
	
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, people.ContactsReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(config)

	srv, err := people.New(client)
	if err != nil {
		log.Fatalf("Unable to create people Client %v", err)
	}

	r, err := srv.People.Connections.List("people/me").PageSize(10).PersonFields("names,emailAddress").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve people. %v", err)
	}
	
	if len(r.Connections) > 0 {
		fmt.Print("List 10 connection name:\n")
		for _, c := range r.Connections {
			names := c.Names
			if len(names) > 0 {
				name := names[0].DisplayName
				fmt.Printf("%s\n", name)
			}
		}
	} else {
		fmt.Println("No connections found.")
	}

}
