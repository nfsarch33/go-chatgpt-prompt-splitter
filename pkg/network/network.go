package network

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func WaitForServer(url string) {
	for {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				log.Println("Server is up!")
				return
			}
		}
		log.Println("Waiting for the server...")
		time.Sleep(time.Second)
	}
}
