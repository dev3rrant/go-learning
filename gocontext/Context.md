# Context:

- Software can run as long running processes
- If the action is cancelled/dies then its needs to be able to stop the processes in a consistent manner
- The context package can help manage long running processes

- Once we have a basic test for the happy path, we can test a situation where the Store can't complete the Fetch before the request is cancelled.

```go
t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

	})
```

## Context Definition:
```
The context package provides functions to derive new Context values from existing ones. These values form a tree: when a Context is canceled, all Contexts derived from it are also canceled.
```

- Derive your contexts so cancellations are propagated through call stack
- The code above derives a new cancellingCtx from the request. This gives a cancel function
- Then we schedule the function to be called in 5 miliseconds with time.Afterfunc
- The context is then used in request by calling `request.WithContext`

- Can use the following robust solution:
```go
func Server(store Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(rw, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
```

- context has a method **Done()**.
	- This returns a channel that gets a signal when the context is done'ed or canceled
	- We listen for that signal and call store.Cancel
- However, we want to ignore that case if Fetch completes before that.
	- To implement this, can run fetch in a goroutine and write result to a new channel, data

- We can use select to race the two async processes and then we write a response or cancel
- This works, but should the server be responsible for managing cancels?
- main aspect of context is it offers a consistent way of managing cancellation

**Go Docs**:
```
Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a Context is canceled, all Contexts derived from it are also canceled.
```

**Go Blog**:
```
At Google, we require that Go programmers pass a Context parameter as the first argument to every function on the call path between incoming and outgoing requests. This allows Go code developed by many different teams to interoperate well. It provides simple control over timeouts and cancelation and ensures that critical values like security credentials transit Go programs properly.
```

- Based on this approach we should just pass context through in request and let all the subroutines handle their cancel process themselves
- In order to test the new context flow, need to replicate a slow request 

```go
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}
```
- 

- This simulates a slow server by slowly appending the string character by character. 
- When it finishes, it sends strings to results channel
- the goroutine also listens for the ctx.Done signal and will stop the process if that signal is received
- Outside of the goroutine there is another select channel that races two async process
	- If ctx done signal is received first, then it errors
	- If the data channel receives data and sends it to res variable first, then the response/result is returned
