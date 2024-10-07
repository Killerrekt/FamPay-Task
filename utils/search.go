package utils

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/youtube/v3"
)

var State = false
var Query = ""

func Search(service *youtube.Service) {
	call := service.Search.List([]string{"id", "snippet"}).
		Q(Query).
		MaxResults(10)
	response, err := call.Do()
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, item := range response.Items {
		fmt.Println(item)
	}
}

func ContinuousFetch() {
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				if State {
					fmt.Println(State)
				} else {
					log.Println("The State is currently at False")
				}
			}
		}
	}()
}
