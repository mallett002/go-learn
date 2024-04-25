package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0]

	fmt.Printf("%v, %T\n", indexed, indexed)

	// go represents strings as utf-8 encoding, (array of bytes)
	// notice the "é" is represented as to bytes:
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// easier way to iterate and indexing strings is to cast to array of runes:------------------

	var myRune = []rune("résumé")
	fmt.Println("myRune: ")

	var myIndexedRune = myRune[1]

	// runes are an alias for int32:
	fmt.Printf("%v, %T\n", myIndexedRune, myIndexedRune)

	// Runes are just unicode point #s that represent the character
	for i, v := range myRune {
		fmt.Println(i, v)
	}

	// declare rune type:
	var newRune = 'a'
	fmt.Printf("%v, %T\n", newRune, newRune)

	// Concatenating strings: ----------------------------------------------------------------
	var strSlice = []string{"h", "e", "l", "l", "o"}

	var catStr = ""

	// inefficient: creates a new string each iteration: (should use string builder like below instead)
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)

	// strings are immutable in go (error):
	// catStr[0] = 'a'

	// use string builder from "strings" package
	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	var builtStr = strBuilder.String()

	fmt.Printf("\n%v", builtStr)
}
