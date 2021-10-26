# Property Based Tests:

- Going to implement the roman numeral kata using test driven development
- Think about how it can be implemented using small slices of functionality and then iterating
- Start with 1 or I


## Implementing basic functionality:
- Start off with a basic case 
- Try with one, then two
- Follow the TDD discipline
- Then get to the four numeral situation

## Refactoring Test Code:
- When testing scenarios like given this input, expect this input, table based testing is a good option.
- Also, starting with basic if statements works for a while, but eventually it gets unwieldly and we can use **strings.Builder** to concatenate strings

- From Docs:
```A Builder is used to efficiently build a string using Write methods. It minimizes memory copying.
```

## Refactoring to Switch Statement:
- Following the discipline eventually get to a place where using a switch statement makes sense, something like this:
```go
for arabic > 0 {
		switch {
		case arabic > 9:
			result.WriteString("X")
			arabic -= 10
		case arabic > 8:
			result.WriteString("IX")
			arabic -= 9
		case arabic > 4:
			result.WriteString("V")
			arabic -= 5
		case arabic > 3:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic--
		}
	}
```

- At this point it starts getting unwieldly
- This is a sign that we are capturing a concept or data inside imperative code when it could be captured in a class structure
- The switch statement is describing some truths about the roman numerals behavior
- Can create a type romanNumeral with string symbol and integer value.
- like this:
```go
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

//All these values have a maximum degree
//They can all be defined in terms of getting the max degree and then handling the remainder

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{3, "III"},
	{2, "II"},
	{1, "I"},
}
```

- Add some more cases like this:
```go
{"40 gets converted to XL", 40, "XL"},
{"47 gets converted to XLVII", 47, "XLVII"},
{"49 gets converted to XLIX", 49, "XLIX"},
{"50 gets converted to L", 50, "L"},
```

- We can keep going with this, but it is now trivial to add support for greater and greater numbers

## Parsing Roman Numerals.
- Now write a function that converts from a roman numeral to an int
- Can use test existing test cases as an examples

## Intro to Property Based Tests:
- As the code has been tested for a variety of scenarios there are certain rules that develop
	- Can't have more than 3 consecutive symbols
	- only 1, 5, 10 can be subtractors
	- Taking the result of convertToRoman(N) and passing it to convertToInt(n) should give N

- Can we take these rules and exercise them against the code?
- **Property based tests** allow you to do this by throwing random data at your code and checking the rules you state beforehand
- Need a good understanding of the domain in order to properly test with property based testing
- Ex:
```go
func TestPropertiesOfConversion(t *testing.T) {
    assertion := func(arabic int) bool {
        roman := ConvertToRoman(arabic)
        fromRoman := ConvertToArabic(roman)
        return fromRoman == arabic
    }

    if err := quick.Check(assertion, nil); err != nil {
        t.Error("failed checks", err)
    }
}
```

- This is done in technical example with the quick package.
- Give it an assertion function that it throws a bunch of type safe data at.
- Running the test against the code communicates there are some issues with it
- It currently uses int, but ints can be positive, while roman numerals cannot
- It would be better to use `uint16`
- This handles the negative case, but the max of uint16 is still way higher than what the code currently supports


## Closing thoughts:
- when coding tough problems, start with something simple and take small steps