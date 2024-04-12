// V3.0 introduces replication. It means we have multiple search services
// (replicas) for each kind and we take the first result returned by the fastest
// replica. This way we dramatically lower the likelihood of discarding results.
// This is a fast and robust program.
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

	go func() { c <- firstResult(query, web1, web2) }()
	go func() { c <- firstResult(query, image1, image2) }()
	go func() { c <- firstResult(query, video1, video2) }()

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
	web1   = search.New("web")
	web2   = search.New("web")
	image1 = search.New("image")
	image2 = search.New("image")
	video1 = search.New("video")
	video2 = search.New("video")
)

func firstResult(query string, replicas ...search.Search) search.Result {
	c := make(chan search.Result)
	for i := range replicas {
		go func(i int) { c <- replicas[i](query) }(i)
	}
	result := <-c
	return result
}
