# Maps:

- Now we learn how to store data as key value pairs
- With maps

## Dictionary:
- start with test

## Map Definition:
- Map definitions are similar to making arrays
- start with map keyword and two types. First is the key type, second is the value type
- Key type is special. Must be a compatible type. Without the ability to tell if 2 keys are equal there's no way to ensure we have the correct value

```go
dictionary := map[string]string{"test": "this is just a test"}
```

### Basic Definition:
```go
func Search(dictionary map[string]string, word string) string {
	return ""
}
```
- Accessing a value from a map is same as getting it out of an array
- `return map[key]`

### Refactoring:
- For refactoring, make a string comparer in the test
- Can also improve dictionary by creating a new type and making search a method

## Further Requirements:
- What happens if we search for a word that doen't currently exist?
- Currently we get nothing back. It is better to report that the word is not in dictionary

### Basic Definition:

```go
t.Run("word exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("word does not exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("blah")
		want := "could not find word you were looking for"
		if err == nil {
			t.Fatal("Expected an error")
		}
		assertStrings(t, err.Error(), want)
	})
```
- The way to handle the not found case is to return an error along with the search term
- In the second test case we can convert the error to a string with .Error(). Also protect assert strings from panicing when it receives a null

```go
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
```
- here we make use of the fact a map look up can optionally return 2 values
- The second is a boolean denoting if look up succeeded.
- for refactoring we can move the error into its own variable. This also improves the test

## Implement Add method:
- start with test
- Adding a new key value pair is very similar to an array
- `map[word] = definition`
- You are able to modify maps without passing an address to it ex. &myMap
- `A map value is a pointer to a hmap structure`
### Note:
- Be careful because maps can be nil. This behaves like an empty map when reading but trying to write to it will cause a runtime panic
- Never initialize an empty map like 
```go
var m map[string]string
```
- always do one of the following:
```go
var dictionary = map[string]string{}
var dictionary = make(map[string]string)
```
### What happens when adding something that already exists:
- Maps in go do not throw an error if the key already exists. Instead it overwrites the new one
- We don't want the Add function possibly updating values.
- We add just enough to test and then refactor by adding a switch statement so we can handle the case where some other unexpected error gets thrown

### Constant Errors:
- Make the errors into constants. I think the idea is that this way anytime you compare two constants with same value, the errors are equal
- If you do not do this, then there is the possiblity that comparing two errors of the same type may not be equal


## Implement Update Functionality:
- similar to add functionality
- We follow a similar process to the add functionality
- We don't want to accidentally add a new word definition in the update function
- We want our functions to be lazy and only do one thing (no unintended side effects)

### Adding Error for Update
- We could have reused ErrNotFound
- It is better to use specific error messages. This allows you to perform different functionality around specific errors.

## Implement Delete Functionality:
- follow the process to create some new functionality
- Once the test and method signature have been created you can use go's built in delete method for maps
- syntax as follows
```go
delete(map,word)
```
- We don't need to complicate the code with any additional checks or errors, because deleting a word that doesn't exist doesn't do anything