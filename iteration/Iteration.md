# Iteration:

### Initial Testing:
- In go there is only for, no while, do each, etc
- the for syntax is similar to other c-like languages
- Ex:
```go 
for i := 0; i < 5; i++ {
	repeated = repeated + character
}
```

- Unlike other languages like java or javascript, there are no parentheses around the three components of the loop

- There is also a change in the beginning of the function

```
var repeated string
```
- `:=` is short hand for declaring and initializing variables
- We can be more explicit with each step
- the above steps initializes a string. This will have a default value of ""

### Refactor:
- We can use += assignment operator
- This is the add AND assignment operator
- Adds the right operand to the left and assigns the result to left operand.
- Also works with other types like integers

### Benchmarking:
- Benchmarks are another first class feature of go. Similar to tests
- the testing.B gives access to the b.N
- when benchmark code runs it runs b.N times and measures how long it takes
- The framework determines what is a good number of times for code to run
- to run benchmarks do
```
go test -bench=
```

- Example Output:
```
goos: darwin
goarch: amd64
pkg: iteration
cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
BenchmarkIteration-8     8578542               134.2 ns/op
PASS
ok      iteration       1.394s
```

- This means it takes an average of 134.2 seconds to run

### Practice Exercise:
- Change test so caller can specify the number of repititions
- Write ExampleRepeat to document code
- Check out the strings package
