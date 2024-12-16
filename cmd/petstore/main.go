package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MXLange/royalfetch/v2"
)

type Pet struct {
	ID        int      `json:"id"`
	Category  Category `json:"category"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []Tag    `json:"tags"`
	Status    string   `json:"status"`
}

// Define a struct para a categoria
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Define a struct para a tag
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fetch := royalfetch.New(royalfetch.RoyalFetch{
		BaseURL:              "https://petstore.swagger.io/v2",
		Retry:                3,
		CodesToRetry:         []int{500, 502, 503, 504},
		WaitingTime:          1000,
		WaitTimeIncreaseRate: 1.25,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, &http.Client{
		Timeout: 5000,
	})

	pet := Pet{
		ID: 1,
		Category: Category{
			ID:   2,
			Name: "categoria1",
		},
		Name:      "doggie",
		PhotoUrls: []string{"url1", "url2"},
		Tags: []Tag{
			{ID: 1, Name: "tag1"},
			{ID: 2, Name: "tag2"},
		},
		Status: "available",
	}

	json, err := json.Marshal(pet)
	if err != nil {
		fmt.Println(err)
	}
	res, err := fetch.Post("/pet", string(json))
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}
