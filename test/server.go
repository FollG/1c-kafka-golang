package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "9092")
	})

	http.ListenAndServe(":9092", nil)
}

//https://juanmanuel-tirado.medium.com/using-the-kafka-rest-api-with-go-44352bc1c803
