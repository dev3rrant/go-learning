# First Program

## Writing the program:
- Writing a hello world
- You will always need a main function
- func keyword is how you define a function
- We also import the fmt package because that has the println function in it

## How to Test:
- It is good to seperate domain code from the outside world
- the fmt.println function is a side effect
- Lets separate these concerns

- Create a new function that returns the hello world string
- Functions are typed, need to denote that you are returning a string

## Writing the Test:
	
```

func HelloTest(t *testing.T)  {
	
}
```
- functions with this parameter denote package testing
- intended to be used with the go test command

- The test will initially fail because it cannot find a go module	

- You don't need to download and install testing frameworks.
- With go it is all part of the language. 

## Writing Test Rules:
- File must have format of XXXX_test.go
- Test function must begin with Test
- Test function must only have parameter: `t *testing.T`
- In order to use `t *testing.T` you must import the testing package

- The `t *testing.T` piece is our hook into the testing framework
- We can use it for things like `t.Fail()`

## Language Basics:
- `if` statements are similar to other languages
	- I notice that you don't need ()
- Variable declaration occurs with := 
	- Ex:
	```
	varName := value
	```
	
- We are calling `t.Errorf`, this lets us call the error function with a formatted string. `%q` lets us double quote the value we want to add to the string 	
- Apparently go automatically documents your code!!!

## Testing going forward:
- We first wrote the code and then the test
- We want to write the test first to capture requirements and then implements


## Constants:
- Constants are to efficiently capture the meaning of values and possibly aid performance
- Ex:
```
const englishHelloPrefix = "Hello, "
```


## SubTests:
- Now we introduce a new testing tool, subtests
- sometimes we want to group tests around a thing and have subtests verify different scenarios
- This lets us set up shared code between tests
- Tests should be clear specifications of what the code needs to do
- We should refactor tests as well as production code

```go
func TestHello(t *testing.T) {

    assertCorrectMessage := func(t testing.TB, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'World'", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

}

```

- We refactor the assertion into a function, this reduces duplication
- In Go, you can can declare functions inside other funtions and assign them to variables, then you call them like normal functions
- We also use testing.TB which is an interface that testing.T and testing.B both satisfy so we can call the helper functions of a test or benchmark
- We include t.Helper() to tell the test suite this is a helper method.
- That way failing tests display line number of the calling function rather than the helper.
	- This aids in debugging

## Testing Discipline:
- It is important to maintain discipline and follow the feedback loop
- This helps you design good code and ensure that your error messages are useful
- Writing fast tests and setting up tools to simply run tests help you get to flow.

## Adding Language Support:
- Now we start adding multiple greeting prefixes based on the language sent as a secondary parameter

## Switch:
- When we have lots of switch statements checking a particular value it can be helpful to use a switch statement instead
- this refactors the code and makes it easier to read

```
switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    }

```
- I am not a big fan of switch statements
- Can go do object literals with defaults similar to javascript
- I think the go colloary is [maps](https://blog.golang.org/maps)

- Ex:
```
greetingsPrefixDefinition := map[string]string{
		"English":   "Hello, ",
		"Spainish":  "Hola, ",
		"French":    "Bonjour, ",
		"Esperanto": "Saluton, ",
	}
```

## Last Refactor:
- Our function is getting a little large.
- We could convert our prefix calculator into its own function
- They do it like this:

```
func greetingPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return
}
```

- In the funct signature they have made a named return value `(prefix string)`
- this creates a variable called prefix in the funct
- it will be assigned the zero value. This depends on the type
	- ints are 0, and strings are ""
	- you can return whatever is set to it with `return` instead of `return prefix`
- This will display in the godoc, making your code more readable
- In Go, functions that start with a capital letter are publically accesible, functions starting with lowercase are private


## Why the steps are important:
- Write a failing test and see it fail so we know we have written a relevant test
- write the smallest amount of code to make it pass so  we know we have working software
- then refactor backed with safety of our tests 