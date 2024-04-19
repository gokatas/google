// V2.0 runs the three kinds of search concurrently using the fan-in pattern.
// Concurrency makes the program faster.
package main

import (
	"fmt"
	"time"

	"google"
)

func main() {
	start := time.Now()
	results := googleIt("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func googleIt(query string) (results []google.Result) {
	c := make(chan google.Result)

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
	web   = google.NewSearch("web")
	image = google.NewSearch("image")
	video = google.NewSearch("video")
)
