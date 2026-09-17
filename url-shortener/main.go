package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type ShortUrl struct {
	URL string `json:"url"`
}

func homeHandler(writer http.ResponseWriter, requestData *http.Request) {
	fmt.Fprintf(writer, "http server")
}

func healthHandler(writer http.ResponseWriter, requestData *http.Request) {
	fmt.Fprint(writer, "u are on health page")
}

func generateShortId() {

}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var fetchUrl ShortUrl // decoded data will be stored in this var
	err := json.NewDecoder(r.Body).Decode(&fetchUrl)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	u, err := url.Parse(fetchUrl.URL)
	if err != nil {
		http.Error(w, "Invalid url", http.StatusBadRequest)
		return
	}

	if u.Scheme == "" || u.Host == "" ||
		(u.Scheme != "http" && u.Scheme != "https") {
		// invalid url
		http.Error(w, "Invalid url", http.StatusBadRequest)
		return
	}

	generateShortId()

	fmt.Printf("u check %v\n", u.Scheme)
	fmt.Fprint(w, fetchUrl.URL, " post req recieved")

}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/shorten", shortenHandler)

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%v \n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("server failed to start: %s ", err)
	}

}
