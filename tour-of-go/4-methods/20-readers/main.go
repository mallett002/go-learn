package main

import (
    "fmt"
    "io"
    "strings"
)

// io.Reader interface used to read many things (files, network connections, ciphers)
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// Read populates slice with data and returns number of bytes populated and an error

func main() {
    str := "Hello Reader!"
    r  := strings.NewReader(str)
    buffer  := make([]byte, 8)
    
    for {
        n, err := r.Read(buffer) // reads 8 bytes at a time (size of buffer)

        fmt.Printf("n = %d, err = %v, buffer = %v\n", n, err, buffer)
        fmt.Printf("buffer[:n] = %q\n", buffer[:n]) // notice: buffer's bytes get replaced from 0 - n

        fmt.Println()

        if err == io.EOF {
            break;
        }
    }
}
