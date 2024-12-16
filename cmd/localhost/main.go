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
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil)

	res, err := fetch.Get("/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resBody))

	wg := sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resError, err := fetch.Get("/error")
			if err != nil {
				panic(err)
			}
			defer resError.Body.Close()

			resErrorBody, err := io.ReadAll(resError.Body)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(resErrorBody))
		}()
	}

	wg.Wait()

	clone := fetch.Clone()
	clone.Headers["Content-Type"] = "application/xml"
	clone.CodesToRetry = []int{500}

	resClone, err := clone.Get("/")
	if err != nil {
		panic(err)
	}
	defer resClone.Body.Close()

	resCloneBody, err := io.ReadAll(resClone.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resCloneBody))

	resOriginal, err := fetch.Get("/")
	if err != nil {
		panic(err)
	}
	defer resOriginal.Body.Close()

	resOriginalBody, err := io.ReadAll(resOriginal.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resOriginalBody))

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resError, err := clone.Get("/error")
			if err != nil {
				panic(err)
			}
			defer resError.Body.Close()

			resErrorBody, err := io.ReadAll(resError.Body)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(resErrorBody))
		}()
	}

	wg.Wait()
}
