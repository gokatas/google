package search

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

type Search func(query string) Result

func New(kind string) Search {
	search := func(query string) Result {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		return Result(fmt.Sprintf("%s search result for %q\n", kind, query))
	}
	return search
}
