package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const searchAPI = "https://api.twitter.com/1.1/search/tweets.json"

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
	fmt.Printf("VALUE: %s\nTYPE: %T\n", r.URL.Path[:1], r.URL.Path[:1])
}

func getTweets(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("QUERY: %s", r.URL.Query())
	fmt.Fprint(w, "Here are your tweets!")
	var res, err = http.Get(searchAPI)

	if err != nil {
		panic("Issue with request")
	}

	fmt.Printf("response: %s\n", io.ReadAtLeast(res.Body))
}
