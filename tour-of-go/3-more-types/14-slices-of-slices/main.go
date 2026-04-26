package main

 import (
	"fmt" 
	"strings"
)

func main() {
	// create tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// players can take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		// strings.Join joins items in slice into a string with a separator
		fmt.Printf("%s\n\n", strings.Join(board[i], "     "))
	}
}

