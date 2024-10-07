package utils

import (
	"fmt"
	"log"

	"google.golang.org/api/youtube/v3"
)

func Search(service *youtube.Service) {
	call := service.Search.List([]string{"id", "snippet"}).
		Q("google").
		MaxResults(10)
	response, err := call.Do()
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, item := range response.Items {
		fmt.Println(item)
	}
}
