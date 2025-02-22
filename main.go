package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://www.lamoda.ru",
		"https://www.yandex.ru",
		"https://www.mail.ru",
		"https://www.google.com",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fmt.Printf("Starting fetch: %s\n", url)
			r := make(chan error)
			err := fetchUrl(url, r)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", url, err)
				return
			}
			fmt.Printf("Finished fetch: %s\n", url)
		}(url)
	}
	wg.Wait()
	fmt.Println("Program finished")
}

func fetchUrl(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func testGracefulShutdown() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan string, 2)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go print(ctx, wg, c)

	c <- "Hello channels"
	c <- "Hello Go"
	close(c)

	wg.Wait()

}

func print(ctx context.Context, wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():

			return
		case s, ok := <-c:
			if !ok {
				return
			}
			time.Sleep(2 * time.Second)
			println(s)
		}
	}
}
