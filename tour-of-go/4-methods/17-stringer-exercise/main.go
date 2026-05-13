package main

import (
    "fmt"
    // "strings"
)

type IPAddr [4]byte

// TODO: add a "String() string" method to IPAddr so it implements the fmt.Stringer interface
// make IPAddr{1, 2, 3, 4} print as "1.2.3.4"

// my implementation
// func (addr IPAddr) String() string {
//     bldr := strings.Builder{}
//
//     for i, digit := range addr {
//         if i == 3 { // last item
//             bldr.WriteString(fmt.Sprintf("%d", digit))
//         } else {
//             bldr.WriteString(fmt.Sprintf("%d.", digit))
//         }
//     }
//
//     return bldr.String()
// }

// AI's solution (simpler)
func (addr IPAddr) String() string {
    return fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
}


func main() {
    hosts := map[string]IPAddr{
        "loopback": {127, 0, 0, 0},
        "googleDNS": {8, 8, 8, 8},
    }

    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
