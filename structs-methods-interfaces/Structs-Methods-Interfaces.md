# Structs, Methods, and Interfaces

- We want a function that can calculate the perimeter of a rectangle given a height and width

## Initial Rectangle Perimeter
### Write a failing test:
- will be using float64 as the type, this denotes a floating point number. 
- ex. 10.04
- You can denote floating point numbers in a format string as `.2f`
- This denotes a floating point number up to 2 digits after decimal

### Write just enough code for compilation
- This is how I defined the signature
```go
func Perimeter(width, height float64) (perimeter float64) {
	return
}
```

### Write just enough for code to pass:
- Pretty basic, formula for perimeter is 2*(width + height)


## Area of Rectangle:
- Now we want to add functionality for calculating the area of a rectangle
- I will try this one on my own
- Follow tdd pattern
- Write a failing test, get compilation to work, get the test to pass

### Refactor:
- Our code works, but it doesn't contain any logic specific to rectangles
- We are assuming usage of this code will always be for rectangles
- An unwary developer might try to use this function for triangle
- An elegant solution is to  create our own type called Rectangle that encapsulates the concept of separate shapes
- Enforcing a type of rectangle conveys our intent more clearly

#### Structs:
- First we need to declare a struct
- Ex:
```
type Rectangle struct {
    Width float64
    Height float64
}

```
- Now we can refactor the tests/code to use the new type, instead of float64
- Creating a new instance of the new struct can look like this:
```
rectangle := Rectangle{10.0,10.0}
```

- Once the rectangle instance is defined in the code, we can access it like `rectangle.Width`
- syntax: `myStruct.field`


- Next requirement is writing a function for area of a circle


## Area of a Circle:
- Follow TDD discipline
- Write a failing test, get it to compile, get it passing, refactor

- Something along these lines for test:
```
func TestArea(t *testing.T) {
	assertEqual := func(t *testing.T, expected, got float64) {
		t.Helper()
		if expected != got {
			t.Errorf("expected %g, received %g", expected, got)
		}
	}
	t.Run("get area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Area(rectangle)
		expected := 100.0
		assertEqual(t, expected, got)
	})

	t.Run("circles", func(t *testing.T) {
        circle := Circle{10}
        got := Area(circle)
        want := 314.1592653589793

        if got != want {
            t.Errorf("got %g want %g", got, want)
        }
    })
}
```
- Use %g to get a more precise decimal point

- Then we create a similar struct to the rectangle
- If we try to create another area method with a circle param. i.e. like this
```go
func Area(circle Circle) float64 { ... }
func Area(rectangle Rectangle) float64 { ... }
```
- we get an error about Area being redeclared in the block
- Go does not allow the above.
- We can have functions with the same name in different packages, or we can define methods on our new types instead

## Methods:
- We have just been writing functions, but have used methods like t.Errorf(...)
- A method is a function with a receiver
	- method declaration binds an identifier (the method name) to a method and associates the method with the receivers base type
- Methods are similar to functions, but are called by invoking them on an instance of a particular object/type
	- Can only call methods on things 

- If we try to do something like this:
```go
func TestArea(t *testing.T) {
	assertEqual := func(t *testing.T, expected, got float64) {
		t.Helper()
		if expected != got {
			t.Errorf("expected %g, received %g", expected, got)
		}
	}
	t.Run("Area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		expected := 100.0
		assertEqual(t, expected, got)
	})

	t.Run("Area of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		expected := 314.1592653589793
		assertEqual(t, expected, got)
	})

}
```
- We get the following errors:
```
./shapes_test.go:29:19: rectangle.Area undefined (type Rectangle has no field or method Area)
./shapes_test.go:36:16: circle.Area undefined (type Circle has no field or method Area)
```

- The go compiler is very helpful in debugging. Make sure to read the error code and understand it

### Write just enough for compilation:
- Method Definitions:
```go
type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) Area() float64 {
	return 0
}

type Circle struct {
	radius float64
}

func (circle Circle) Area() float64 {
	return 0
}
```

- Method definition syntax: `func (receiverName ReceiverType) MethodName(args)`
- Method definitions are similar to funcs
- The difference is the syntax of the method receiver
- When your method is called on a variable of that type. You get the reference to its data via receiverName
- Convention is to have receiver variable be the first letter of the type
- Ex Definition:
```go

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

```
- Note we use the constant of pi from the math package


### Refactor:
- There is duplication in our tests
- We want to take a collection of shapes and call area method on each shape and then check the result
- We want to write a checkArea function we can pass shapes to, but fails for non shapes

## Interfaces:
- In go, we codify this intent with interfaces.
- they allow you to make functions that can be used with different types and create highly decoupled code and maintain type safey

- We define what we want in the tests:

```
func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if want != got {
			t.Errorf("got %g, want %g", want, got)
		}
	}
	t.Run("Area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("Area of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})

}

```

- This fails to compile because shape is undefined
- Interface definition:

```
type Shape interface {
    Area() float64
}
```
- Very similar to struct definition
- And just with that the test passes. You don't need to add an implements to the definition 
- the interface resolution in go is implicit 
- For example, the rectangle struct has an Area function that returns float64, so it satisfies the interface

#### Decoupling:
- The helper no longer needs to know what kind of shape is passed
- It just fulfils its single responsiblity of checking the results
- Thus, the helper is decoupled from the concrete types
- Using interfaces we only use what we need


#### Further Refactoring (Table Driven Tests):
- Now we introduce table driven tests
- This lets us build a list of test cases that can be tested in the same manner

```go
func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12, 6}, 72.0}
		{Circle{10},314.1592653589793}
	}

	for _, tt:= range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)	
		}
	}
}
```

- Here we use an anonymous struct, areaTests
- It declares a slice of structs with []struct with two fields, shape and want
- Then we iterate over the definitions using the structs fields to  run the tests
- This makes it simple to add additional cases/objects

## Area of a Triangle:

### Follow the testing discipline
- Write the test, write just enough for compiler, then get the tests working

### Refactor:
- Looking at the definition of the test cases they do not scan well. It can be difficult to read the intent of these params
- Structs allow you to optionally name the fields

## Test Output:
- Take a moment to ensure your test output is helpful
- If you have  a large table with many cases, just showing what you expected and what you want might not be clear where the issue/ or what tests the test is occuring in
- Can use this syntax: 
```go
%#v got %g want %g
```
- `%#v` will print our struct with the values in its field
- Could also use a t.Run with each instance. That way we can name the test and explicitly point where the issue is occuring
