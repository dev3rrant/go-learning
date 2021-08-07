# Arrays and Slices:

## Definition:
- Arrays let you store multiple elements of the same type in a variable in a particular order
- It is common to iterate over the elements in an array performing some kind of action with them
- First, we want to create a Sum function that will take an array of numbers and return the total
- Start with the test

- Arrays can be defined like this:
```
numbers := [5]int{1,2,3,4,5}
```
- In go, arrays have a fixed capacity defined when you declare the variable
- Two ways of initializing
	- [N]type{value1, value2, ..., valueN} e.g. ` numbers := [5]int{1, 2, 3, 4, 5}`
	- [...]type{value1, value2, ..., valueN} e.g. `numbers := [...]int{1, 2, 3, 4, 5}`

## Write just enough to pass:

- You can define the array in the signature as follows:
```	
package main

func Sum(numbers [5]int) int {
    return 0
}
```

- Interestingly, the author sets the package as main

- Once we have a failing test, we add the amount of code necessary to get the test to pass
- Accessing the elements of the array function as many other languages do

```go
for i := 0; i < 5; i++ {
	sum += numbers[i]
}
```

## Refactor:
- Now we introduce range construct 
- range allows you to iterate over an array. 
- on each iteration range returns two values - the index and the value
- In this example, we ignore the index with the _ blank indentifier

```
func Sum(numbers [5]int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}
```

- The size of an array is encoded in the type
- If you try to pass an array of size 4 to function that takes an array of size 5, it will fail to compile
- Most of the time you don't use fixed length arrays.
- Instead you use slices


### Slices:
- The slice type allows you to have collections of any size
- Definition is similar to arrays, you just don't define the length

```
 numbers := []int{1, 2, 3}
 
```

- However, we need to edit the first test so it sends a slice, rather than a length defined array

### Refactoring Slices:
- The sum code itself has been refactored and is in good shape
- However, we should also take time to refactor the test code if possible
- Remember, we can create a named function as a variable
	- Don't forget to add the t.Helper() so the correct error line is displayed

- It is important to question the value of your tests. They should be written to give you confidence in the code, not just quantity of tests


## SumAll:
- next we want a new function SumAll that takes in a varying number of slices and returns a new slice containing the results of each individual slice
- For example, 
- `SumAll([]int{1,2}, []int{0,9}) would return []int{3, 9}`

### Write Test:
- First we write the test, usuage looks something like this
```go
got := SumAll([]int{1, 2}, []int{0, 9})
want := []int{3, 9}
```

- After you see the undefined error writing the test, write the basic definition so compilation is successful
- To define the method signature of SumAll we will use variadic functions
- These types of functions can be called with any number of trailing arguments.
- Can work with range and the varying variable
- Definition looks like this:
```go
func SumAll(numbers ...[]int) (sums []int) {
	return
}
```
- Here we use the return variable again
- The test still fails because we can't use equality operators with slices
- To get around this we will use a function from the standard library
- [DeepEqual](https://pkg.go.dev/reflect#DeepEqual)
	- reports whether x and y are deeply equal
- You need to be careful when using DeepEqual because it is not type safe
- This can lead to tricky to find bugs

### Get Test Passing:

```
func SumAll(numberGroups ...[]int) []int {
	lengthOfNumbers := len(numberGroups)
	sums := make([]int, lengthOfNumbers)
	for i, numbers := range numberGroups {
		sums[i] = Sum(numbers)
	}
	return sums
}
```
- Many new things
- len lets you get the length of the object 
- We use a new way of creating slices. `make` allows you to create a slice with a starting capacity of len of the numberGroups we are working with
- We also index the slice using `slice[N]`, to get a value or assign a new one

### Refactoring:
- Slices have a capacity and if you try to access/change a value that doesn't exist you get a runtime error
	- slice with capacity 2 and you try to access index 4
- To resolve this, we can use the append method 
	- This takes a slice and a new value and returns a new slice with all items (existing and new) in it


## SumAllTrails:
- Next requirement is to change SumAll to SumAllTails 
- This will calculate the total of the tails of each slice. 
- Tail of a collection is all items in the collection, minus the first (head) element

### Write enough for compilation
- rename the SumAll function

### Get Test Passing:
- write just enough for the test to pass
- We do this by slicing the slice
- Syntax: `slice[low:high]`
- `slice[1:]` takes everything from the index 1 to the end. 
	- beginning would be 0

### Refactor:
- Not much to refactor, but how do you handle case where we are try to take a slice of an empty array
- When we try to run the test we see that the test compiles, so we have a run time error
- We can do this by checking the length of the slice before using it

```go

func SumAllTails(numberGroups ...[]int) []int {
	var sums []int
	for _, numbers := range numberGroups {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

```

### Test Refactor:
- Remove duplication of the tests
- A handy effect of adding the helper compare function is that it adds type checking. We have to pass two slices to the function
- Slices work with other types of objects 