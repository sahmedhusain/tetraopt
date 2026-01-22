package main

import (
	"fmt"
	"os"
	"tetris-opt/packages"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Incorrect number of arguments\nUsage: go run main.go <filename>")
		return
	}

	pieces, err := packages.ParseFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: Invalid file", err)
		return
	}

	board := packages.Solve(pieces)
	if board == nil {
		fmt.Println("ERROR: No solution found")
		return
	}

	fmt.Print(board.String())
}
