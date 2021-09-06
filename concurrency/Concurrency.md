# Concurrency:

- We have the following code.
```go
package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)

    for _, url := range urls {
        results[url] = wc(url)
    }

    return results
}
```

- Pass a map of urls and a function for checking if those urls are valid
- Here is the test:
```go
package concurrency

import (
    "reflect"
    "testing"
)

func mockWebsiteChecker(url string) bool {
    if url == "waat://furhurterwe.geds" {
        return false
    }
    return true
}

func TestCheckWebsites(t *testing.T) {
    websites := []string{
        "http://google.com",
        "http://blog.gypsydave5.com",
        "waat://furhurterwe.geds",
    }

    want := map[string]bool{
        "http://google.com":          true,
        "http://blog.gypsydave5.com": true,
        "waat://furhurterwe.geds":    false,
    }

    got := CheckWebsites(mockWebsiteChecker, websites)

    if !reflect.DeepEqual(want, got) {
        t.Fatalf("Wanted %v, got %v", want, got)
    }
}
```
- They use dependency injection to allow them to pass a checker function that makes testing easy and not require actually hitting a url
- We want to speed it up
- First, we add a benchmark to see the effect of our changes:

```go
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
```
- Here we use a slice of 100 urls and a fake implementation of websitechecker
- then we trigger the benchmark with `go test -bench=.`

## What is Concurrency:
- concurrency: having more than one thing process
- instead of waiting for a website to respond before sending a request to the next website we tell the computer to make the next request while it waits
- Normally in go we wait for a function to return, for it to complete. This operation is **blocking**: makes us wait for it to finish
- An operation that does not block in go runs in a separate process called a **goroutine**
	- A normal process has one read that starts at the top and then reads down the page one line at a time. Going inside a process, completing it and coming back out to continue
	- When a separate process starts with a go routine, it is like another reader starts that process and the original one continues with the page.

## Creating a goroutine:
- to start a new goroutine, turn a function call into a go statement using the go keyword in front of it like `go doSomething()`
- Ex:
```go
package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)

    for _, url := range urls {
        go func() {
            results[url] = wc(url)
        }()
    }

    return results
}
```
- This is the only way to start a goroutine in go. Because of this we often use anonymous functions when we want to start a goroutine

## Anonymous Functions:
- They can be executed the same time they are declared. that is what the () represents in the goroutine function
- All the variables defined at the time when the anon func is called are available to the func
- The current code can result in a panic. Parallelism is hard. We need to make sure we handle concurrency predictably


## Race Condition:
- Currently, the tests will still fail because the goroutines do not have enough time to add their results to the map
- We could add a sleep of a couple of seconds, but the url is used in each goroutine, its not a copy
- So they all rewrite the value and we only end up getting the last url
- If we add a separate url to the anonymous function like this:
```go
for _, url := range urls {
		go func(u string) {
			results[url] = wc(url)
		}(url)
	}
```
- We will still probably see an error about concurrent map rewrites
- This is a **race condition**
- Race Condition: a bug that occurs when output of our software is dependant on the timing and sequence of events we don't control
- We don't control when each gorouting adds to the map, so we are vulnerable to them writing to the map at the same time
- Go has a built in race detector. You can use it with `go test -race`

## Channels:
- We can solve the race condition using **channels**.
- Channels are a go data structure that can send and receive values. These operations allow communication between processes
- Ex:
```go
package concurrency

type WebsiteChecker func(string) bool
type result struct {
    string
    bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)
    resultChannel := make(chan result)

    for _, url := range urls {
        go func(u string) {
            resultChannel <- result{u, wc(u)}
        }(url)
    }

    for i := 0; i < len(urls); i++ {
        r := <-resultChannel
        results[r.string] = r.bool
    }

    return results
}
```
- The results channel receives a result (where the website url is checked)
	- We do this using the send operator `<-`
- Then we iterate through the result to construct a map of results we can return
- Note, when you aren't sure what to name params on a struct you can anonymously use their types.

## Parallism and Linearity:
- With the goroutines we parallized the part of the code that we want to make faster
- Using the channel we were also able to keep parts of the code that need to happen linearly


## Words of Wisdom:
```
- Make it work, make it right, make it fast


- Premature optimization is the root of all evil

```