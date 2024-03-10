# ErrLess: Simplified Error Handling in GoLang Inspired by [GoLang 2.0 Error Handling Proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md)

ErrLess is a Go library that brings the simplicity and elegance of the proposed GoLang 2.0 error handling to Go
developers **today**.
Inspired by the draft design for error handling in Go 2.0

## Error Handling — Problem Overview

Dealing with errors in GoLang can be a bit cumbersome. The common pattern is to check for errors after each function call that returns an error. This can lead to a lot of repetitive code and can make the code harder to read and maintain.

Details can be found in the [**GoLang 2.0 Error Handling Overview**](https://github.com/golang/proposal/blob/master/design/go2draft-error-handling-overview.md) document as starting point of error handling proposal of golang 2.0

## Error Handling Proposal for GoLang 2.0
This below text is taken from the [GoLang 2.0 Error Handling Proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md)

### Design

The draft design introduces the keywords `check` and `handle`, which we will introduce first by example.

Today, errors are commonly handled in Go using the following pattern:

```go
func printSum(a, b string) error {
    x, err := strconv.Atoi(a)
    if err != nil {
        return err
    }
    y, err := strconv.Atoi(b)
    if err != nil {
        return err
    }
    fmt.Println("result:", x + y)
    return nil
}
```

With the `check/handle` construct, we can instead write:

``` go
func printSum(a, b string) error {
        handle err { return err }
        x := check strconv.Atoi(a)
        y := check strconv.Atoi(b)
        fmt.Println("result:", x + y)
        return nil
}
```

For each check, there is an implicit **handler** chain function, explained in more detail
in [the link](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md#design). Here, the handler
chain is the same for each check and is defined by the single handle statement to be:

```go
func handleChain(err error) error {
    return err
}

```

## ErrLess: Bringing Go 2.0 Error Handling to Go Today

Below code sniped is the same function(`printSum`) implemented with ErrLess.
The implementation is almost the same with the GoLang 2.0 Error Handling Proposal.

```go
func printSum(a, b string) (err error) {
    defer errless.Handle(&err, func (err error) error {
        return err
    })
    x := errless.Check1(strconv.Atoi(a))
    y := errless.Check1(strconv.Atoi(b))
    fmt.Println("result:", x + y)
    return nil
}


```

**Check1** is a generic function. It can be used with any function that returns a single value and error.
The signature of **Check1** is:

```go
func Check1[A any](a A, err error) A
```

`strconv.Atoi` is returning a value and an error that's why **Check1** is used here. If number of return value
changes, you can use **Check2**, **Check3**, **Check4**, **Check5**.


Please closely look at the function's signature(`printSum(a, b string) (err error)`), where we use a named return value for the error(`err`).
This allows the defer statement to access and modify the error value returned by the handle function. If you want to know more about **[Named Return Values](https://go.dev/doc/effective_go#named-results)** and [**Defer**](https://go.dev/doc/effective_go#defer), please check  [here](https://go.dev/doc/effective_go#named-results).

ErrLess implementation fo all example in
the [GoLang 2.0 Error Handling Proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md)
can be found in here.

## Features

### **Simplified Error Checking**:

With functions like Check, **Check1**, **Check2**, **Check3**, **Check4**, and **Check5**,
ErrLess allows developers to handle errors without cluttering their code with repetitive error checks.

**Before:**

```go
x, err := strconv.Atoi(a)
    if err != nil {
    return err
}
```

After:

```go
x := errless.Check1(strconv.Atoi(a))
```

### **Scoped Error Handler**: 
If you want to add context to the error for specific error check,
you can implement custom error handler. In this case, we need to use **Check1W**, **Check2W**, **Check3W**, **Check4W**,
**Check5W** functions.
This allow us to reuse the error handler for multiple error checks.  

```go
convertionError:= func (err error) error {
return fmt.Errorf("failed to convert to int: %w", err)
}
x := errless.Check1W(strconv.Atoi(a)).With(convertionError)
y := errless.Check1W(strconv.Atoi(a)).W(convertionError)
z := errless.Check1W(strconv.Atoi(a)).W(convertionError)
```

**Function Scoped Error Handler**: If you want to add context to the error for **any** errors can happen in a function call,you can use **errless.Handle** method.
want to handle the error, you can use **errless.EmptyHandler**.

Imagine you want to add "sum of values failed" to the returning error of the function. You can use below code.

```go
func sum(a, b string) (res int, err error) {
    defer errless.Handle(&err, func (err error) error {
        return fmt.Errorf("sum of values failed: %w", err)
    })
    x := errless.Check1(strconv.Atoi(a))
    y := errless.Check1(strconv.Atoi(b))
    return x + y, nil
}
```
it will add "sum of values failed" to the returning error of the function not mather error is coming from `strconv.Atoi(a)` or `strconv.Atoi(b)`.


### **Static Type Check**: 
Leveraging Go's generics, ErrLess provides a flexible way to work with functions that return
multiple values along with an error. Thanks to generics, **all type checking is done at compile time**.



## **Getting Started**

Using it is quite straightforward.

Just `defer errless.Handle` method beginning of the function and wrap any error returning function with `Check` ,`Check1`, `Check2`, `Check3`, `Check4`, `Check5` and remove the error checking code.

```go

import (
    "fmt"
    "log"
    "strconv"
    
    "github.com/mfatihercik/errless"
)

func sum(a, b string) (res int, err error) {
    defer errless.Handle(&err, errless.EmptyHandler)
    x := errless.Check1(strconv.Atoi(a))
    y := errless.Check1(strconv.Atoi(b))
    return x + y, nil
}

func main() {
    res, err := sum("10", "20")
    if err != nil {
        log.Fatalf("Error occurred: %v", err)
    }
    fmt.Println("result:", res)

}
```

Please check examples in the [/example](/example) folder.


## Contributing

Build Project
```shell
make build
```

Run Tests
```shell
make test
```


## License

MIT License