package main

import {
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
}


//Storing credentials for Twitter RESTful API

type Credentials struct {
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
}

//twitterClient is a helper function that will return twitter
//client that will be used to send tweets or stream new tweets
//this will take in a pointer to a Credentials struct which will
//contains everything needed to authenticate and return a pointer
//to twiter client or an error

func  twitterClient(creds *Credentials)(*twitter.Client, error){

	//Passing in consumer key (API key) and Cosumer Secret
	// (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.CosumerSecret)

	//Passing in the Access Token and Access Toke Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	//Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus: twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	//we cn retrieve the user and verify if the credentials
	//we have used successfully allow us to login
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's Account:\n%+v\n", user)
	return client, nil
}

func main(){
	fmt.Println("Go-Twitter-Bot v0.01")
	creds := Credentials{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey: os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET")
	}

	client, err := getClient(&creds)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
}
