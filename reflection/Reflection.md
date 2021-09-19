# Reflection:

- Challenge Statement:
```
golang challenge: write a function walk(x interface{}, fn func(string)) which takes a struct x and calls fn for all strings fields found inside. difficulty level: recursively.
```

- Need to use reflection:
```
Reflection in computing is the ability of a program to examine its own structure, particularly through types; it's a form of metaprogramming. It's also a great source of confusion.
```

## interface:
- Sometimes you want to write a function, but you don't know the type at compile time
- Go lets you get around this with the  interface{}
	- This essentially means any type
- So this, `walk(x interface{}, fn func(string))` will accept any type of object for x
- Be advised though, you lose type safety doing this
- Therefore, you need to use reflection to figure out the type and what you can do with it
	- Generally less performant and can be difficult to understand


## First try at Reflection:
```go
func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)
    field := val.Field(0)
    fn(field.String())
}
```
- This works, but it is making a lot of risky assumptions
- Namely, that there is a field on x and that it has a String func
- The ValueOf function allows us to inspect a value and get at its fields
- Val also has a method numField, that returns number of the field
- Can use this in order to support multiple string fields

- Then need to check the type of the field and ensure it is a string
- However, if we try to use a struct with its own internal struct then it still faiils
- Fortunately, we can call the walk method again. Thus we get the recursion
- This still fails if we an object is passed in with a pointer
	- This is because we can't use NumField on a pointer. Extract value with Elem method
- Following the same process we can do arrays and maps. 
- Arrays are handled the same way as slices so we can combine them in case with following:

```go
witch val.Kind() {
    case reflect.String:
        fn(val.String())
    case reflect.Struct:
        numberOfValues = val.NumField()
        getField = val.Field
    case reflect.Slice, reflect.Array:
        numberOfValues = val.Len()
        getField = val.Index
    }
```
