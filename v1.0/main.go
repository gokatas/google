// V1.0 invokes (fake) web, image and video searches serially.
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
	results = append(results, web(query))
	results = append(results, image(query))
	results = append(results, video(query))
	return
}

var (
	web   = google.NewSearch("web")
	image = google.NewSearch("image")
	video = google.NewSearch("video")
)
