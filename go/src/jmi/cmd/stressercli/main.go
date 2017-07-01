package main

import (
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"randomstring"
	"sync"
	"time"
)

var wc sync.WaitGroup

func main() {

	wc.Add(1)

	totalKeys := 120000

	keys := make([]string, totalKeys, totalKeys)

	for i := 0; i < totalKeys; i++ {
		keys[i] = randomstring.RandStringRunes(120)
	}

	for i := 0; i < 4; i++ {
		go func() {
			var netClient = &http.Client{
				Timeout: time.Second * 10,
			}

			http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000

			for {
				resp, err := netClient.Get("http://localhost:9092/" + keys[rand.Intn(totalKeys)])
				if err != nil {
					log.Println("error")
				} else {
					io.Copy(ioutil.Discard, resp.Body)
					resp.Body.Close()
				}
				time.Sleep(100 * time.Millisecond)
				//if resp != nil {
				//	log.Println(resp.Status)
				//}
			}
		}()
	}

	wc.Wait()
}
