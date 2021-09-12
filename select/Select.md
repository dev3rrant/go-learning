# Select:

- We will be building a website racer
- We will use:
	- net/http to make requests
	- net/http/httptest to test
	- goroutines
	- select to synchronise processes

## Beginning:
- We start off with contrived example, following tdd patterns we've been using
- However, we are testing real websites, so that takes a while
- Go has tools in the standard library to test urls
- We can use httptest to mock an http server

## httptest:
- httptest.NewServer takees in a http.HandlerFunc we pass as an anonymous func
- The function needs a ResponseWriter and Request
- This is also how you set up a real server in go

## Refactor:
- There is duplication in how we measure the response time.
- Move that to its own method
- Can also move the construction of the test servers into their own method

## defer:
- Prefixing a func call with **defer** will cause the function to be called at the end of the containing function
- This is helpful for cleanup like closing a server or file

## select:
- go is great at concurrency, we should be able to check the speed of the sites at the same time. 
- We don't care about the times, we care about which one comes back first
- first we start be creating a func ping that creates a chan struct and returns it
	- chan structs are the smallest data type, since this is a simple case we don't need to allocate anything special

## make channels
- defining a channel like `var ch chan struct{}` creates the objects zero value. In the case of channels this means it is nil
- Therefore anytime we try to send to it, it will block forever
- You can wait for values to be sent to a channel with `myVar := <-ch`
	- this is a blocking call
- **select** lets you wait on multiple channels. The first one to send a value wins and the code underneath that executes

## Adding a Timeout:
- First, add the tests
- Once we have a failing test we add an error to the return objects and then fix the existing tests
- In order to throw an error after a certain time has occurred we can use time.After in conjunction with the select
- something like this:
```go
case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
```


## Handling Slow Tests:
- This takes a long time to test some simple functionality
- Can use a similar idea as dependency injection. We can add a configurable timeout.
- However, adding a timeout does not really help the actual behavior of the racer. The requirements are specific on 10s so we can rework the racer so it is configurable, while not requiring additional configuration to the users that don't need it

