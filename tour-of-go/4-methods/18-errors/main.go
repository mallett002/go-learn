package main

import (
    "fmt"
    "time"
    "strconv"
)

// error interface is built-in
// type error interface {
//     Error() string
// }
// fmt package looks for error interface when printing values

// functions should return errors and calling code should check if err != nil

type MyError struct {
    When time.Time
    What string
}

// Implement Error interface
func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
    return &MyError{
        When: time.Now(),
        What: "it didn't work",
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Println(err)
    }

    // other example:
    // functions return errors, callers handle
    // strconv.Atoi() is equivalent to ParseInt(s, 10, 0)
    num, err := strconv.Atoi("27")
    if err != nil {
        fmt.Printf("couldn't convert the number %v\n", err)
        return
    }

    fmt.Printf("converted the number %d\n", num)
}
