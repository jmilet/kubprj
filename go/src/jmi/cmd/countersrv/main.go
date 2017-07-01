package main

import (
	"fmt"
	"jmi/counter"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Server staring...")

	input := make(chan string)

	for i := 0; i < 8; i++ {
		go counter.Run(i, input)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		input <- string(r.URL.Path)
		hostName, err := os.Hostname()
		if err != nil {
			log.Fatal("Error getting hostname")
		}

		hostName = fmt.Sprintf(`{"Hostname:" "%s"}`+"\n", hostName)

		w.Write([]byte(hostName))
	})

	http.ListenAndServe(":9092", nil)
}
