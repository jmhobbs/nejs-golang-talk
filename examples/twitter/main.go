package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	http.HandleFunc("/tweets", func(w http.ResponseWriter, r *http.Request) {
		tweets, err := getTweets("nejsconf", "recent", "20")
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(tweets)
	})

	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Println("Listening on http://127.0.0.1:9090/")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getTweets(query, result_type, count string) ([]byte, error) {
	v := url.Values{
		"q":                []string{query},
		"result_type":      []string{result_type},
		"count":            []string{count},
		"include_entities": []string{"false"},
	}

	u, err := url.Parse("https://api.twitter.com/1.1/search/tweets.json")
	if err != nil {
		return []byte{}, err
	}
	u.RawQuery = v.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TWITTER_TOKEN")))
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	return ioutil.ReadAll(resp.Body)
}
