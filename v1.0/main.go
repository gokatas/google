// V1.0 invokes (fake) Web, Image and Video searches serially.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

type Result string

func Google(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

var (
	Web   = NewSearch("web")
	Image = NewSearch("image")
	Video = NewSearch("video")
)

type Search func(query string) Result

func NewSearch(kind string) Search {
	search := func(query string) Result {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		return Result(fmt.Sprintf("%s search result for %q\n", kind, query))
	}
	return search
}
