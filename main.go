package main

import (
	"fmt"
	"os"
	"tetris-optimizer/pkg"
)

func main() {
	arg := os.Args
	if len(arg) != 2 {
		fmt.Println("Please write go run . FileName.txt")
	} else {
		FilePath := "Examples/" + arg[1]
		optimizer.ReadFile(FilePath)
		optimizer.Solve()
		optimizer.PrintSolution() 
	}
}
