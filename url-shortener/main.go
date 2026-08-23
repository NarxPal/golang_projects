package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(writer http.ResponseWriter, requestData *http.Request) {
	fmt.Fprintf(writer, "http server")
}

func healthHandler(writer http.ResponseWriter, requestData *http.Request) {
	fmt.Fprint(writer, "u are on health page")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%v \n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("server failed to start: %s ", err)
	}

}
