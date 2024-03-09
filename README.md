# ErrLess: Simplified Error Handling in GoLang Inspired by [GoLang 2.0 Error Handling Proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md)

ErrLess is a Go library that brings the simplicity and elegance of the proposed GoLang 2.0 error handling to Go
developers **today**.
Inspired by the draft design for error handling in Go 2.0

## Error Handling Proposal for GoLang 2.0
This below text is taken from the [GoLang 2.0 Error Handling Proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md)

#### Design

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

**Check1** is a generic function. It can be used with any function that returns a single value and an error. 
The signature of **Check1** is:

```
func Check1[A any](a A, err error) A
```

`strconv.Atoi` is returning a single value and an error that's why **Check1** is used here. If number of return value
changes, you can use **Check2**, **Check3**, **Check4**, **Check5**.


Please closely look at the function's signature(`printSum(a, b string) (err error)`), where we use a named return value for the error(`err`).
This allows the defer statement to access and modify the error value returned by the handle function. If you want to know more about **[Named Return Values](https://go.dev/doc/effective_go#named-results)** and [**Defer**](https://go.dev/doc/effective_go#defer), please check  [here](https://go.dev/doc/effective_go#named-results).

## Features

**Simplified Error Checking**: With functions like Check, **Check1**, **Check2**, **Check3**, **Check4**, and **Check5**, ErrLess allows developers to handle errors without cluttering their code with repetitive error checks.

**Generic Support**: Leveraging Go's generics, ErrLess provides a flexible way to work with functions that return
multiple values along with an error.

**Custom Error Handler**: You can define your own error handler function to enrich the returning error. If you don't want to handle the error, you can use **errless.EmptyHandler**.


## **Getting Started**

Using it is quite straightforward. 

Just `defer errless.Handle` method bigining of the function and wrap any error returning function with `Check` ,`Check1`, `Check2`, `Check3`, `Check4`, `Check5` and remove the error checking code.

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