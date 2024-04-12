// V1.0 invokes (fake) Web, Image and Video searches serially.
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
	results = append(results, web(query))
	results = append(results, image(query))
	results = append(results, video(query))
	return
}

var (
	web   = search.New("web")
	image = search.New("image")
	video = search.New("video")
)
