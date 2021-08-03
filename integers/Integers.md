# Integers:

- Integers work similar to most other languages
- We will start with an addition function
- We first start off with building the test
- We are now going to use %d in the error string instead of %q

```go
package integers

import "testing"

func TestAdder(t *testing.T) {
    sum := Add(2, 2)
    expected := 4

    if sum != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
}
```
- Other change is that we are using package integers. This denotes we are working with code grouped around working with integers
- Then we write just enough code to satisfy the compiler. This ensures our tests fail for the reasons we expect

```go
func Add(first, second int) int {
	return 0
}
```
- Another note, is when you have parameters with the same type one after the other you only need to put the type after the last one
- About the only factoring we can do is adding a comment

## Godocs:
- You can access automatically generated documentation for this package at `http://localhost:6060/pkg/integers/`.
	- after you run `godoc -http=:6060`

### Examples:
- You can add examples. Go examples are different in that they are executed just like the tests
- As with typical tests, examples are functions in the _test.go files
- They are prepended with Example{Name}	
- Ex:

```go
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}

```
- The build will fail if the example is no longer valid