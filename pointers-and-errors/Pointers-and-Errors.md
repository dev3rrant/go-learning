# Errors & Pointers:

- Fintech likes go, apparently
- We will start with a wallet that lets us deposit bitcoin

## Note:
- The gopkg manager works better when you open one vscode folder for each go mod. It finds the path better if you open the project at the pointers-and-errors section rather than root of go-learning

## Wallet
### Deposit:
- Follow the discipline, start with unit test and then write just enough to progress and learn more from the compiler
- We use structs here, because instead of access fields directly, we want to hide implementation details from users
- One we have the test failing because we return the wrong value we are going to add a balance param
- In Go, if a symbol (variable, type, functions, etc.) starts with a lowercase symbol then it is private outside the package its defined in
	- We want the methods of struct to be able to manipulate balance param, but nothing else should be able to 

- We eventually get to this point:
```go
package main

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}

```
- However, our test still fails
- It fails because, **when you call a function or method the arguments are copied**
- The memory location of the wallet in the test is different from the location of the wallet in the struct
- We have not experienced this issue previously because in structs chapter we just use the values of the copies, we don't try to change
- Since we want to change the param itself, balance, we need to access it in memory
- We can fix this with pointers

## Pointers:
- Pointers let us point to some values and lets us change them
- Instead of getting a copy of the wallet we take a pointer to the wallet so we can change it

- Pointers in go are defined in a similar fashion to c
```
func (w *Wallet) Deposit(amount int) {
    w.balance += amount
}

func (w *Wallet) Balance() int {
    return w.balance
}
```

### Automatic Dereferencing:
- we didn't dereference the pointers like `(*w).balance` 
- The above is valid, however, the developers of go deemed this cumbersome, so we can use what are called struct poointers that are automatically dereferenced


## New from Existing Types:
- we want to deposit bitcoins, but int has all the functionality we need
- However, it is not immeadiately obvious what our intention is
- We can build off of existing types and extend them to add our own functionality
- syntax is as follows:
```go
type MyName OriginalType

```
- This allows us to also add methods and other functionality to existing types
- for example, we can implement Stringer interface on Bitcoin
- Stringer is an interface from the fmt package and lets you define how your type gets printed
- when used with the %s string in prints
- The syntax for creating a method on a type alias is the same as a struct
- Next, update the test format string so they use String() instead
```go
if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
```
### Withdraw:
- Next, we want to add a withdraw function. The opposite of deposit
- Write the tests
- In order to test just the withdraw function, you can define the wallet with an initial state
- Can define that as following:
```go
		wallet := Wallet{balance: Bitcoin(20)}
```
- Follow the discipline to get the test to pass
- Then refactor the test to remove duplication

## Overdraft Issue:
- What happens when we try to withdraw an amount greater than the balance?
- We are currently assuming there will be no overdraft
- In go, when you want to indicate an error for your function, return an error foir the caller to check and act on.
- Can demonstrate expected behavior in a test

```go
t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(20))
		assertBalance(t, wallet, Bitcoin(10))
		if err == nil {
			t.Error("wanted an error, but didn't get one")
		}
	})
```
- In go, nil is synonmous to null in other programming languages
- We can make the return type of the withdraw function, error
- this is an interface and can be made nillable
- use `errors.New` to create a new error with a message of your choice

```go
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("oh no")
	}
	w.balance -= amount
	return nil
}

```

### Refactor:
- We can convert the error checking block into its own function

```go
assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error, but didn't get one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
```
- Can also update the test to verify the error message that comes back
- We replace the first t.Errorf with t.Fatal. 
- This stops the test if its reached
- We do this because we don't want to to make any more assertions if there is no error
- Without the test this would go to next step and trigger a panic due to nil pointer
- We still have duplication of the error message in withdraw and test code
- If someone rewords the error this would fail. This is currently a brittle test
-  we can make the error message global with a var declaration. That way there is only one source of truth about it

```go
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
```

- Another additional tool for help handling errors is a linter called errcheck
```bash
go get -u github.com/kisielk/errcheck
```