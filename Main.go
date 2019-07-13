package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	const port = "3000"

	//Endpoints
	http.HandleFunc("/", handler)
	http.HandleFunc("/tweets", handleTweets)

	// Logs
	log.Printf("listening on %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func twitterAPI() (twtAPI *twitter.Client) {
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

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("VALUE: %s\nTYPE: %T\n", r.URL.Path[:1], r.URL.Path[:1])
}

func handleTweets(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["user"]

	if !ok || len(keys[0]) < 1 {
		http.Error(w, "error: Missing 'user' param", 422)
		return
	}

	tweets, _, err := twitterAPI().Search.Tweets(&twitter.SearchTweetParams{
		Query: string(keys[0]),
	})

	if err != nil {
		http.Error(w, "Failed to retrieve tweets", 422)
		return
	}

	t, _ := json.Marshal(tweets)
	w.Write(t)
}
