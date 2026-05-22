package main

import (
	"io"
	"os"
	"strings"
)

// Common pattern: io.Reader that wraps another io.Reader
// example:
// gzip reader takes in an io.Reader
// returns a "*gzip.Reader" (also implements io.Reader)
// zr, err := gzip.NewReader(req.Body)

// The takeaway is the Decorator pattern: 
// you can wrap an io.Reader with another io.Reader to transparently transform data on-the-fly 
// without changing the caller's code (io.Copy doesn't know or care rot13 is happening). 
// This is how gzip.Reader, tls.Reader, base64.Reader, etc. all work — composition via interface satisfaction.
// ** The transformation is invisible to the caller **

// will make rot13Reader an io.Reader (along with its reader function)
type rot13Reader struct {
    reader io.Reader
}

// var rot13Map map[rune]rune = map[rune]rune {
//     'A': 'N', 'B': 'O', 'C': 'P', 'D': 'Q', 'E': 'R', 'F': 'S', 'G': 'T', 'H': 'U', 'I': 'V', 'J': 'W', 'K': 'X', 'L': 'Y', 'M': 'Z',
//     'N': 'A', 'O': 'B', 'P': 'C', 'Q': 'D', 'R': 'E', 'S': 'F', 'T': 'G', 'U': 'H', 'V': 'I', 'W': 'J', 'X': 'K', 'Y': 'L', 'Z': 'M',
//     'a': 'n', 'b': 'o', 'c': 'p', 'd': 'q', 'e': 'r', 'f': 's', 'g': 't', 'h': 'u', 'i': 'v', 'j': 'w', 'k': 'x', 'l': 'y', 'm': 'z',
//     'n': 'a', 'o': 'b', 'p': 'c', 'q': 'd', 'r': 'e', 's': 'f', 't': 'g', 'u': 'h', 'v': 'i', 'w': 'j', 'x': 'k', 'y': 'l', 'z': 'm',
// }

// rot13Reader implements io.Reader (read method)
// applies rot 13 cipher - replaces each char with the 13th letter after it in the alphabet
// Option A: Use lookup map:
// func (rot rot13Reader) Read(buffer []byte) (int,  error) {
//     // pass the buffer to our rot13 reader
//     n, err := rot.reader.Read(buffer)
//
//     // loop over number of bytes read
//     // lookup each one and replace buffer with decrypted rune
//     for i := 0; i < n; i++ {
//         if r, ok := rot13Map[rune(buffer[i])]; ok {
//             buffer[i] = byte(r)
//         }
//     }
//
//     return n, err
// }


// bite - 'A' -> gets index (0 for A)
// + 13 -> applies rot13 shift
// % 26  -> wraps around (n itself or how many nums past 26)
// 'A' + -> converts index back to an ASCII letter
// 'A' is a byte (rune or a uint8)
// byte + (byte - byte + byte) % byte --> they are untyped constants - the flex to fit what is needed
func rot13(bite byte) byte {
    switch {
    case 'A' <= bite && bite <= 'Z': // btw A and Z
        return 'A' + (bite - 'A' + 13) % 26
    case 'a' <= bite && bite <= 'z': // btw a and z
        return 'a' + (bite - 'a' + 13) % 26
    default:
        return bite
    }
}

// Option B: 
func (rot rot13Reader) Read(buffer []byte) (int, error) {
    n, err := rot.reader.Read(buffer)

    for i := 0; i < n; i++ {
        buffer[i] = rot13((buffer[i]))
    }

    return n, err
}


func main() {
    stringsReader := strings.NewReader("Lbh penpxrq gur pbqr!")

    // rot13Reader wraps the strings reader
    r13Reader := rot13Reader{
        reader: stringsReader,
    }

    // io.Copy just knows, "I need to call a .Read() func"
    // the rot13Reader does the transformation
    // io.Copy doesn't know anything about a transformation, just how to call the .Read()
    io.Copy(os.Stdout, r13Reader)
}

