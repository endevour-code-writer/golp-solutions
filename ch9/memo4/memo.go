// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 276.

// Package memo provides a concurrency-safe memoization a function of
// a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a Mutex.
package memo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var done = make(chan struct{})

// Func is the type of the function to memoize.
type Func func(key string, done chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// !+
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key, done)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}

	select {
	case <-done:
		e.res.value = nil
	default:
	}

	return e.res.value, e.res.err
}

func InitCancelListener() {
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()
}

func IsCancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func Test() {
	m := New(httpGetBody)
	incomingUrls := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://golang.org",
		"https://godoc.org",
		"https://pkg.go.dev/os",
	}
	InitCancelListener()

	for i, url := range incomingUrls {
		fmt.Println(url)
		_, err := m.Get(url)

		if err != nil {
			fmt.Println(err)
		}

		if IsCancelled() {
			fmt.Println("The operation is cancelled")

			unprocessed := incomingUrls[i : len(incomingUrls)-1]

			for _, url := range unprocessed {
				fmt.Printf("The url %v was unprocessed\n", url)
			}

			return
		}
	}
}

func httpGetBody(url string, done chan struct{}) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//!-
