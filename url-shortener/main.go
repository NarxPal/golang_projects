package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"net/url"
	"strings"
)

type ShortUrl struct {
	URL string `json:"url"`
}

// urlStore stores id -> url
var urlStore = make(map[string]string)

// idStore stores url -> id (to find duplicate id's for single url)
var idStore = make(map[string]string)

func homeHandler(writer http.ResponseWriter, requestData *http.Request) {
	path := requestData.URL.Path
	if path == "/" {
		fmt.Fprintf(writer, "http server")
		return
	}
	shortId := strings.TrimPrefix(path, "/")
	fmt.Printf("url.path check %v\n", shortId)

	urlPath, exists := urlStore[shortId]
	if !exists {
		http.Error(writer, "url for id not found", http.StatusNotFound)
	}

	fmt.Printf("urlpath check %v\n", urlPath)
	http.Redirect(writer, requestData, urlPath, http.StatusFound)
}

func healthHandler(writer http.ResponseWriter, requestData *http.Request) {
	fmt.Fprint(writer, "u are on health page")
}

func generateShortId() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	emptyId := make([]byte, 6)
	for i := range emptyId {
		emptyId[i] = charset[rand.IntN(len(charset))]
	}
	randomStr := string(emptyId)
	return randomStr
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

	//  map the url wrt id using idStore
	shortId, exists := idStore[fetchUrl.URL]

	if exists {
		fmt.Fprint(w, shortId)
		fmt.Printf("map check %v\n", urlStore)
		return
	}
	randShortId := generateShortId()
	fmt.Printf("short id : %v\n", randShortId)
	urlStore[randShortId] = fetchUrl.URL
	idStore[fetchUrl.URL] = randShortId

	fmt.Fprint(w, randShortId)
	fmt.Printf("map check %v\n", urlStore)
	fmt.Printf("u check %v\n", u.Scheme)
	fmt.Fprint(w, " post req recieved")

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
