# Dependency Injection:

- Want to write a greet function. However, we want to do the actual printing now
```go
func Greet(name string) {
    fmt.Printf("Hello, %s", name)
}
```
- Something like above is hard to test because it prints to std out
- Our function doesn't care where/how printing occurs. We can accept an interface rather than a concrete type
	- This lets us change the implementation of the printing mechanism easily
- This is the fmt.Printf implementation:
```go
//fmt.printf implementation

// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, format, a...)
}

//fprintf implementation
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
    p := newPrinter()
    p.doPrintf(format, a)
    n, err = w.Write(p.buf)
    p.free()
    return
}

//io.writer definition:
type Writer interface {
    Write(p []byte) (n int, err error)
}

```
- reading between the lines, we are using Writer to send our greeting somewhere
- We can write a test using bytes.Buffer because bytes.Buffer implemented the Writer interface because it has a Write method implementation

## Implementation:
- Follow the TDD process
- First implement is with regular fmt.Printf.
- This will fail, but will display greeting in output. It is printing to std out still
- Then fix the test by using Fprintf. This also takes in a writer that will be be sent the text

### Refactoring:
- Current implementation is brittle. 
- It is only accepting buffers, but what if we want to use something like os.Stdout?
- We then use something like this:
```go
func Greet(writer io.Writer, name string) {}
```

## Why use DI:
- If you can't test code easily it can be due to hardwired dependancies. DI lets you inject a database dependancy that you can then mock out with something you can control in your test.
- Separate concerns
- code can be used in different contexts