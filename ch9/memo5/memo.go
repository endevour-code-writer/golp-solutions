// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 278.

// Package memo provides a concurrency-safe non-blocking memoization
// of a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package memo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//!+Func

var done = make(chan struct{})
var test = make(chan int)

// Func is the type of the function to memoize.
type Func func(key string, done chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response

	if IsCancelled() {
		res.value = nil
	}
	// fmt.Println(res.value)
	// fmt.Println()
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

//!-get

//!+monitor

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e

			if IsCancelled() {
				for range memo.requests {
				}

				return
			}

			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key, done)
	// Broadcast the ready condition.

	if IsCancelled() {
		e.res.value = nil
		close(e.ready)

		return
	}

	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res

	if IsCancelled() {
		return
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

func httpGetBody(url string, done chan struct{}) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

//!-monitor
