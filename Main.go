package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	const port = "3000"

	//Endpoints
	http.HandleFunc("/", handler)
	http.HandleFunc("/tweets", getTweets)

	// Logs
	log.Printf("listening on %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("VALUE: %s\nTYPE: %T\n", r.URL.Path[:1], r.URL.Path[:1])
}

func getTweets(w http.ResponseWriter, r *http.Request) {
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

	tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})

	if err != nil {
		panic("Issue with request")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusBadRequest {
		w.Write(body)
	}

	if resp.StatusCode == http.StatusOK {
		t, _ := json.Marshal(tweets)
		w.Write(t)
	}
}
