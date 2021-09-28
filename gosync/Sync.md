# Sync:

## Problem Statement:
- Want to make a counter that is safe to use concurrently
- Start with unsafe counter
- Then exercise its unsafeness with different goroutines

## Initial Testing:
- First we set up a basic counter using the test driven development discipline
	- Dry up the assert calls
- Next requirements are that it must be safe to use in concurrent contexts

```go
t.Run("it runs safely concurrently", func(t *testing.T) {
    wantedCount := 1000
    counter := Counter{}

    var wg sync.WaitGroup
    wg.Add(wantedCount)

    for i := 0; i < wantedCount; i++ {
        go func() {
            counter.Inc()
            wg.Done()
        }()
    }
    wg.Wait()

    assertCounter(t, counter, wantedCount)
})
```

- Loop through wantedCount and fire a goroutine to increment counter
- This uses **sync.WaitGroup**.
	- This is a convienient way to synchronize concurrent processes
```
A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.
```

- Trying to run the test as is, fails. 
- Multiple processes/goroutines are trying to mutate the value at the same time


## Mutex:
- A simple solution is to use a **mutex**
- mutex is mutual exclusion lock. 
	- zero value for a mutal lock is an unlocked mutex
- Can use it like this:
```go
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += 1
}
```

- This means the first goroutine to call Inc gets the lock on Counter.
- Then all other goroutines need to wait for it be unlocked before they can access it

## Embedding Mutex in struct:
- It is possible to embed the mutex as follows:
```go
func (c *Counter) Inc() {
    c.Lock()
    defer c.Unlock()
    c.value++
}
```

- This can be dangerous because clients of your code can incorrectly call the lock/unlock methods and impact the mutex.
- This is an example of why you need to be careful with the public interface you expose in your API's
	- Once something is exposed, it can be coupled to by clients

## go vet results:
- running `go vet` states that the code is currently copying the mutex.
- In the assert call we pass counter by value so it copies it. 
- From the documentation:
`A mutex must not be copied after its first use`

- To resolve this, we can pass a pointer to the mutex in the assert helper method
- This makes test fail to compile
- Can address this problem by adding a method that handles how to correctly create the object:


## When to Use Locks vs Channels/Goroutines:
- Use channels when passing ownership of data
- Use mutex for managing state.

## go vet
- Use go vet in build scripts because it can find subtle bugs around timing/concurrency

## Don't use embedding