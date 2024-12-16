package main

import (
	"github.com/MXLange/royalfetch/v2"
	"github.com/MXLange/royalfetch/v2/auth"
	"github.com/MXLange/royalfetch/v2/proxy"
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
		Proxy: &proxy.Proxy{
			Host: "http://localhost",
			Port: 8080,
			Auth: &auth.BasicAuth{
				Username: "user",
				Password: "password",
			},
		},
		Auth: &auth.Auth{
			BasicAuth: &auth.BasicAuth{
				Username: "user",
				Password: "password",
			},
		},
	})
}

// func main() {
// 	fetch := royalfetch.New(royalfetch.RoyalFetch{
// 		BaseURL:              "http://localhost:5000",
// 		Retry:                4,
// 		WaitingTime:          1000,
// 		WaitTimeIncreaseRate: 2,
// 		CodesToRetry:         []int{500, 502, 503, 504},
// 		Timeout:              5000,
// 		Headers: map[string]string{
// 			"Content-Type": "application/json",
// 		},
// 	})

// 	res, err := fetch.Get("/")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer res.Body.Close()
// 	bytes, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(bytes))

// }

// func main() {
// 	fetch := royalfetch.New(royalfetch.RoyalFetch{
// 		BaseURL:              "https://petstore.swagger.io/v2",
// 		Retry:                3,
// 		WaitingTime:          1000,
// 		WaitTimeIncreaseRate: 1.25,
// 		Timeout:              5000,
// 		Headers: map[string]string{
// 			"Content-Type": "application/json",
// 		},
// 	})

// 	pet := Pet{
// 		ID: 1,
// 		Category: Category{
// 			ID:   2,
// 			Name: "categoria1",
// 		},
// 		Name:      "doggie",
// 		PhotoUrls: []string{"url1", "url2"},
// 		Tags: []Tag{
// 			{ID: 1, Name: "tag1"},
// 			{ID: 2, Name: "tag2"},
// 		},
// 		Status: "available",
// 	}

// 	json, err := json.Marshal(pet)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	res, err := fetch.Post("/pet", string(json))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer res.Body.Close()
// 	bytes, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(bytes))
// }
