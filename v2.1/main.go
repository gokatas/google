// V2.1 times out the search after 80ms using the timeout pattern. Consequently
// it sometimes returns only partial results. Thus it is fast but not very
// robust.
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

	timeout := time.After(time.Millisecond * 80)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}

	return
}

var (
	web   = search.New("web")
	image = search.New("image")
	video = search.New("video")
)
