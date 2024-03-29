# ErrLess: Simplified Error Handling in GoLang Inspired by [GoLang 2.0 Error Handling Proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md)

ErrLess is a Go library that brings the simplicity and elegance of the proposed GoLang 2.0 error handling to Go
developers **today**.
Inspired by the draft design and ["try proposal"](https://github.com/golang/proposal/blob/master/design/32437-try-builtin.md) for error handling in Go 2.0 

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
The implementation is almost the same with the GoLang 2.0 Error Handling Proposal. In here, we inspired from ["try proposal"](https://github.com/golang/proposal/blob/master/design/32437-try-builtin.md) for naming of the functions.

```go
func printSum(a, b string) (err error) {
    defer errless.Handle(&err, func (err error) error {
        return err
    })
    x := errless.Try1(strconv.Atoi(a))
    y := errless.Try1(strconv.Atoi(b))
    fmt.Println("result:", x + y)
    return nil
}


```

**Try1** is a generic function. It can be used with any function that returns a single value and error.
The signature of **Try1** is:

```go
func Try1[A any](a A, err error) A
```

`strconv.Atoi` is returning a value and an error that's why **Try1** is used here. If number of return value
changes, you can use **Try2**, **Try3**, **Try4**, **Try5**.


Please closely look at the function's signature(`printSum(a, b string) (err error)`), where we use a named return value for the error(`err`).
This allows the defer statement to access and modify the error value returned by the handle function. If you want to know more about **[Named Return Values](https://go.dev/doc/effective_go#named-results)** and [**Defer**](https://go.dev/doc/effective_go#defer), please check  [here](https://go.dev/doc/effective_go#named-results).

ErrLess implementation of all example in
the [GoLang 2.0 Error Handling Proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md)
can be found in at the **[end of the file](/README.md#errless-implementation-of-examples-in-golang-20-error-handling-proposal)**.

## Features

### **Simplified Error Checking**:

With functions like Try, **Try1**, **Try2**, **Try3**, **Try4**, and **Try5**,
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
x := errless.Try1(strconv.Atoi(a))
```

### **Scoped Error Handler**: 
If you want to add context to the error for specific error check,
you can implement custom error handler. In this case, we need to use **Try1**, **Try2**, **Try3**, **Try4**,
**Try5** functions.
This allow us to reuse the error handler for multiple error checks.  

```go
convertionError:= func (err error) error {
return fmt.Errorf("failed to convert to int: %w", err)
}
x := errless.Try1(strconv.Atoi(a)).Err(convertionError)
y := errless.Try1(strconv.Atoi(a)).Err(convertionError)
z := errless.Try1(strconv.Atoi(a)).Err(convertionError)
```

### **Quickly Add Context to the Error**:
You can add additional context to the error with **ErrMessage** or **ErrWrap**  method. This will add passed message to the error message.

```go

x := errless.Try1(strconv.Atoi(a)).ErrMessage("failed to convert to int")

```
You can even implement your own message addtion and use it with **With** method.

Create A generic Handler:
```go
func message(message string,params...interfaces{}) HandlerFunc {
    return func(err error) error {
            return fmt.Errorf(message+"- error: %s",params...,err)
    }
}
```
Use the customised handler:
```go
x := errless.Try1(strconv.Atoi(a)).Err(message("failed to convert to int"))

````



###  **Scoped Error Handling in Functions**: 
If you want to add context to the error for **any** errors can happen in a function call,you can use **Handle** or **Catch** method.

Imagine you want to add "sum of values failed" to the returning error of the function. You can use below code.

```go
func sum(a, b string) (res int, err error) {
    defer errless.Catch(func (e error) {
		// assign to named error return value
        err= fmt.Errorf("sum of values failed: %w", e)
    })
    x := errless.Try1(strconv.Atoi(a))
    y := errless.Try1(strconv.Atoi(b))
    return x + y, nil
}
```
it will add "sum of values failed" to the returning error of the function not mather error is coming from `strconv.Atoi(a)` or `strconv.Atoi(b)`.

###  **Filter Error With `If` Function**:
You can filter the error before it is handled with **If** method. This will allow you to handle only specific errors.
You can use one of pre build filter functions or you can implement your own filter function.
Availabe prebuild functions are **Is**, **IsNot**, **Contains**, **NotContains**

```go
func getFromDB(id string) (res string, err error) {
    defer errless.HandleErr(&err)
    res = e.Try1(getFromDB(id)).If(e.IsNot(sql.ErrNoRows)).ErrMessage("no record found for id: "+id))
    return res, nil
}
```

### **Error Fallback with the `Fallback` Function**:
You can use **Fallback** method to provide a fallback value for executed function. 
Assume you calling a database  to get a record and 
you want to provide a default value if record is not found. You can use `Or` function

```go
func getFromDB(id string) (res string, err error) {
    defer errless.HandleErr(&err)
    res = e.Try1(getFromDB(id)).If(e.Is(sql.ErrNoRows).Fallback(func(err error) string {
            return e.Throw1(CreateNewRecord(id))
    }))
    return res, nil
}
```


### **Static Type Check**: 
Leveraging Go's generics, ErrLess provides a flexible way to work with functions that return
multiple values along with an error. Thanks to generics, **all type checking is done at compile time**.



## **Getting Started**

Using it is quite straightforward.

Just `defer errless.Handle` method beginning of the function and wrap any error returning function with `Try` ,`Try1`, `Check2`, `Check3`, `Try4`, `Try5` and remove the error checking code.

```go

import (
    "fmt"
    "log"
    "strconv"
    
    "github.com/mfatihercik/errless"
)

func sum(a, b string) (res int, err error) {
    defer errless.Handle(&err, errless.EmptyHandler)
    x := errless.Try1(strconv.Atoi(a))
    y := errless.Try1(strconv.Atoi(b))
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


## ErrLess implementation of Examples in GoLang 2.0 Error Handling Proposal

You can find the implementation of the examples in the [GoLang 2.0 Error Handling Proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md) below.
You can see that only changes are the usage of `errless.Try1` and `errless.Handle` methods instead of the `check` and `handle` keywords.

**Real Code Conversion** examples are in [here](https://gist.github.com/mfatihercik/5cf423bb720d7f7313d976b06360fdbb)

## `PrintSum` Example

 **Implementation GoLang 2.0 Error Handling Proposal**
link: [Printsum](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md#design)

```go
func printSum(a, b string) error {
    handle err { return err }
    x := check strconv.Atoi(a)
    y := check strconv.Atoi(b)
    fmt.Println("result:", x + y)
    return nil
}
```

**Implementation with ErrLess**

```go
import (
e "github.com/mfatihercik/errless"
)

func printSum(a, b string) (err error) {
    defer e.HandleReturn(func (e error){
        err = e
    }
    x := e.Try1(strconv.Atoi(a))
    y := e.Try1(strconv.Atoi(a))
    fmt.Println("result:", x + y)
    return nil
}
```



### Inline `PrintSum` Example
 **Implementation GoLang 2.0 Error Handling Proposal**
Link: [Inline Printsum](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md#checks)

```go
func printSum(a, b string) error {
	handle err { return err }
	fmt.Println("result:", check strconv.Atoi(x) + check strconv.Atoi(y))
	return nil
}
```

**Implementation with ErrLess**


```go
import (
e "github.com/mfatihercik/errless"
)

func printSum(a, b string) (err error) {
    defer e.HandleReturn(func (e error){
        err = e
    }
    fmt.Println("result:", e.Try1(strconv.Atoi(a)) + e.Try1(strconv.Atoi(b)))
    return nil
}
```


## `process` Example

**Implementation GoLang 2.0 Error Handling Proposal**
Link: [process](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md#handlers)

```go
func process(user string, files chan string) (n int, err error) {
    handle err { return 0, fmt.Errorf("process: %v", err)  }      // handler A
    for i := 0; i < 3; i++ {
        handle err { err = fmt.Errorf("attempt %d: %v", i, err) } // handler B
        handle err { err = moreWrapping(err) }                    // handler C

        check do(something())  // check 1: handler chain C, B, A
    }
    check do(somethingElse())  // check 2: handler chain A
}
```

**Implementation with ErrLess**

```go
import (
e "github.com/mfatihercik/errless"
)

func process(user string, files chan string) (n int, err error) {
	defer e.HandleReturn(func (e error){
		n= 0
		err = fmt.Errorf("process: %v", e)
   })  // handler A
    for i := 0; i < 3; i++ {
        handleB:= func (err error) { err = fmt.Errorf("attempt %d: %v", i, err) } // handler B
        handleC:=func (err error) { err = moreWrapping(err) }                    // handler C
        
        e.Try1(do(something())).Err(handleC,handleB)  // check 1: handler chain C, B, A
    }
    e.Try1(do(somethingElse()))  // check 2: handler chain A
}

```

## `TestFoo` Example

**Implementation GoLang 2.0 Error Handling Proposal**
Link: [TestFoo](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md#stack-frame-preservation)

```go
func TestFoo(t *testing.T) {
	handle err {
        t.Helper()
		t.Fatal(err) 
	}
	for _, tc := range testCases {
		x := check Foo(tc.a)
		y := check Foo(tc.b)
		if x != y {
			t.Errorf("Foo(%v) != Foo(%v)", tc.a, tc.b)
		}
	}
}
```
**Implementation with ErrLess**

```go
func TestFoo(t *testing.T) {
    defer e.HandleReturn(func (e error){
		t.Helper()
        t.Fatal(e)
   })
	for _, tc := range testCases {
		x := e.Try1(Foo(tc.a))
		y := e.Try1(Foo(tc.b))
		if x != y {
			t.Errorf("Foo(%v) != Foo(%v)", tc.a, tc.b)
		}
	}
}
```

## `SortContents` Example

**Implementation GoLang 2.0 Error Handling Proposal**
Link: [SortContents](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md#examples)

```go
func SortContents(w io.Writer, files []string) error {
    handle err {
        return fmt.Errorf("process: %v", err)             // handler A
    }

    lines := []strings{}
    for _, file := range files {
        handle err {
            return fmt.Errorf("read %s: %v ", file, err)  // handler B
        }
        scan := bufio.NewScanner(check os.Open(file))     // check runs B on error
        for scan.Scan() {
            lines = append(lines, scan.Text())
        }
        check scan.Err()                                  // check runs B on error
    }
    sort.Strings(lines)
    for _, line := range lines {
        check io.WriteString(w, line)                     // check runs A on error
    }
}
```

**Implementation with ErrLess**

```go
func SortContents(w io.Writer, files []string) (err error) {

	defer e.HandleReturn(func(e error) {
		err = fmt.Errorf("process: %v", e) // handler A
	})

	lines := []string{}
	for _, file := range files {
		handleB := func(err error) error {
			return fmt.Errorf("read %s: %v ", file, err) // handler B
		}
		scan := bufio.NewScanner(e.Try1(os.Open(file)).Err(handleB)) // handler B

		for scan.Scan() {
			lines = append(lines, scan.Text())
		}
		e.Try(scan.Err()).Err(handleB) // handler B
	}
	sort.Strings(lines)
	for _, line := range lines {
		e.Try1(io.WriteString(w, line)) // check runs A on error
	}
	return nil
}
```


## `CopyFile` Example

GoLang 1.0 Implementation
```go
func CopyFile(src, dst string) error {
	r, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("copy %s %s: %v", src, dst, err)
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("copy %s %s: %v", src, dst, err)
	}

	if _, err := io.Copy(w, r); err != nil {
		w.Close()
		os.Remove(dst)
		return fmt.Errorf("copy %s %s: %v", src, dst, err)
	}

	if err := w.Close(); err != nil {
		os.Remove(dst)
		return fmt.Errorf("copy %s %s: %v", src, dst, err)
	}
}
```

**Implementation GoLang 2.0 Error Handling Proposal**
Link: [CopyFile]https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling-overview.md#problem)

```go
func CopyFile(src, dst string) error {
	handle err {
		return fmt.Errorf("copy %s %s: %v", src, dst, err)
	}

	r := check os.Open(src)
	defer r.Close()

	w := check os.Create(dst)
	handle err {
		w.Close()
		os.Remove(dst) // (only if a check fails)
	}

	check io.Copy(w, r)
	check w.Close()
	return nil
}
```

**Implementation with ErrLess**

```go
func CopyFile(src, dst string) error {
    defer e.HandleReturn(func(e error) {
        err = fmt.Errorf("copy %s %s: %v", src, dst, e)
    })
	
	r := e.Try1(os.Open(src))
	defer r.Close()

	w := e.Try1(os.Create(dst))
	
	removeFile:= (err error) {
		w.Close()
		os.Remove(dst) // (only if a check fails)
	}

    e.Try1(io.Copy(w, r)).Err(removeFile)
    e.Try1(w.Close()).Err(removeFile)
	return nil
}
```