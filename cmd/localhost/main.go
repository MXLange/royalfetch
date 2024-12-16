package main

import (
	"fmt"
	"io"
	"sync"

	"github.com/MXLange/royalfetch/v2"
)

func main() {
	fetch := royalfetch.New(royalfetch.RoyalFetch{
		BaseURL:              "http://localhost:8080",
		Retry:                3,
		CodesToRetry:         []int{404},
		WaitingTime:          1000,
		WaitTimeIncreaseRate: 3,
	}, nil)

	res, err := fetch.Get("/")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(resBody))

	wg := sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resError, err := fetch.Get("/error")
			if err != nil {
				fmt.Println(err)
			}
			defer resError.Body.Close()

			resErrorBody, err := io.ReadAll(resError.Body)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(resErrorBody))
		}()
	}

	wg.Wait()
}
