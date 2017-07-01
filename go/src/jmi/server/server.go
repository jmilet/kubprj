package main

import (
	"fmt"
	"jmi/counter"
	"net/http"
)

func main() {
	fmt.Println("Server staring...")

	input := make(chan string)

	for i := 0; i < 8; i++ {
		go counter.Run(i, input)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		input <- string(r.URL.Path)
	})

	http.ListenAndServe(":9092", nil)
}
