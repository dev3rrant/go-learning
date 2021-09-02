# Mocking:

## Coding Iteratively:
- Want to write code that prints out a countdown from 3 and go
- We want to implement this in an iterative manner
- This means we take smallest steps possible to ensure we have useful software
- Spending a lot of time writing a big section of code you theoretically think works can lead to a lot of time spent running down rabbit holes.
- **Slick up requirements as small as you can to have working software**


## Separation:
- Start with printing 3
- Print 3,2,1 and Go 
- Wait a second between each line

- We want to write data somewhere, io.Writer is the standard library way of doing this using an interface in go
-Start off by creating a func that takes a bytes.Buffer object. This object implements the io.Writer interface
- Once we verify that works, we can replace bytes.Buffer with io.Writer (more general interface)

**Take a thin slice of functionality and make it work end to end, backed by tests**
- This lets us iterate on the solution safely and easily. Now we don't need to stop and rerun the program to be confident it works.

#### Writing the test:
- Since we are testing output on different lines we can use the backtick character to represent this.
- Ex:
```go
want := `3
2
1
go`
```

- Next, lets get it to print 2,1, and then go
- Once we have the test passing we are going to add 1 second pauses
- We are able to do this using time.Sleep

## Mocking:
- Tests pass, but it takes 4 seconds to run
- slow tests reduce developer productivity
- We are also not testing an important component of our function
- We are depending on Sleep. We should extract this, so we can control it in tests

- We can mock time.Sleep. Then use dependency injection to use it instead of a real time.Sleep call and then spy on the calls to make assertions about them

### Writing Mocked Test:
- Start with defining dependency as an interface
- Then we can use a real sleeper in main, and a spy sleeper in tests
- Countdown is now oblivious to sleep implementation. We give more flexibility to caller

#### Mocked Sleeper Ex:
```go
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
```
- Spys are a type of mock that record how a dependancy is used
- they can record arguments sent, how many times it is called, etc.
- After we get the test passing there still an issue
- We are just checking that sleep is called four times. They could be called out of sequence
- Wew create a countDownOperationsSpy. If this implements io.Writer and Sleeper interfaces then we can record its calls
```go
type CountdownOperationsSpy struct {
    Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

const write = "write"
const sleep = "sleep"
```


## Adding further requirements:
- now want the sleeper to be configurable. We want to be able to change the length of time it sleeps.

- Start off by creating a new configurable sleeper
- This sleeper has two params
- One is duration. This tells it how long to sleep for. (instead of hardcoding it like default sleeper)
- It also takes a sleep function. The param name is sleep, but then we pass it the func and type of param it expects
- Then in actual implementation we can call the param function and pass it the other param, which matches the type the func expects
```go
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
```

- create a spy struct for tracking the time
- then give it a sleep method so it implements the sleeper interface
- We create a configurable sleeper and pass it a duration and the spys sleep function

## Closing Thoughts:
- People can get into a bad state with mocking when they don't listen to tests, or respect the refactor step
- If your mocking code is too complicated, or you have to mock lots of things, you should really considering a different approach
	- Break the module into smaller pieces
	- Consolidate dependencies into a meaningful module
	- test expected behavior rather than implementation

### Rules of Thumb:
- Refactoring, the code changes, but the behavior stays the same
	- if I refactor, would that require a lot of changes to tests?
- Go lets you test private functions, but avoid it because they are implementation details. Test the public behavior. Don't couple less stable private methods.
- If a test requires 3 or more mocks, it is a red flag. 
- Use spies with caution. They can lead to tighter coupling because code tests and implementation

```
"When to use iterative development? You should use iterative development only on projects that you want to succeed."
- Martin Fowler.
```