package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/gorilla/mux"
)

/*
* Function Name: twitterAPI
* Description: Initialises the HTTP Client with keys and secrets
*
* Params: nil
* Return: twitter.Client
 */
func twitterAPI() *twitter.Client {
	const (
		consumerKey    = "237yGfkFsctxG2YKMhP5lxQyS"
		consumerSecret = "FkKAqXzlzQFLKcf1NIUgrsCduCic5ZDFu3jESDca2G4QkKdRTp"
		accessToken    = "2820529160-RTQ0CFS8qDhlW8kDeP0I26N6Vl6FhdhSsPIXYKY"
		accessSecret   = "Zkb495dO1bYGo0wvC7cfqy6qI5MIKqImMl5rlMRt0HzUa"
	)
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	return client
}

// API front page response
const apiHome = "<b>Welcome to the API version: 1.0.0 the available endpoints are:</b><br><br><ol><li>/</li><li>/tweets - Params: [username]</li><li>/tweets/top-monthly - Params: [username]</li></ol>"

func main() {
	// Environment Variables
	const port = "3000"
	const env = "development"
	const verision = "1.0.0"

	// Endpoints
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	router.HandleFunc("/tweets/{username}", handleTweets)
	router.HandleFunc("/tweets/{username}/top-monthly", handleTopTweets)

	http.Handle("/", router)

	// Environment Logs
	log.Printf(`🚨  Server started at: localhost:%s`, port)
	log.Printf(`🛰  API: localhost:%s`, port)
	log.Printf(`🍃  Enviroment: %s`, env)
	log.Printf(`🏷️  Version: %s`, verision)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	fmt.Fprint(w, apiHome)
}

func handleTweets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweets, err := getTweets(string(vars["username"]), 10)

	if err != nil {
		http.Error(w, "Failed to retrieve tweets", 422)
		return
	}

	twts, _ := json.Marshal(tweets)

	w.Write(twts)
}

func handleTopTweets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweets, err := getTweets(string(vars["username"]), 200)

	if err != nil {
		http.Error(w, "Failed to retrieve tweets", 422)
		return
	}

	sort.Slice(tweets.Statuses, func(i, j int) bool { return tweets.Statuses[i].FavoriteCount > tweets.Statuses[j].FavoriteCount })

	tweets.Statuses = tweets.Statuses[0:10]

	twts, _ := json.Marshal(tweets)

	w.Write(twts)
}

func getTweets(username string, tweetCount int) (*twitter.Search, error) {
	tweets, _, err := twitterAPI().Search.Tweets(&twitter.SearchTweetParams{
		Query: username,
		Count: tweetCount,
	})

	return tweets, err
}
