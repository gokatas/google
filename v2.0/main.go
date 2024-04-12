// V2.0 runs the three kinds of search concurrently using the fan-in pattern.
// Concurrency makes the program faster.
package main

import (
	"fmt"
	"time"

	"google/search"
)

func main() {
	start := time.Now()
	results := google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func google(query string) (results []search.Result) {
	c := make(chan search.Result)

	go func() { c <- web(query) }()
	go func() { c <- image(query) }()
	go func() { c <- video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}

	return
}

var (
	web   = search.New("web")
	image = search.New("image")
	video = search.New("video")
)
