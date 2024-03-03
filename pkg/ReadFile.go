package optimizer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(filePath string) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tetromino Tetromino
	Char := int('A')

	for scanner.Scan() {
		if Char > 'Z' {
			fmt.Println("ERROR")
			os.Exit(0)
		}
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" && len(tetromino.Diagram) > 0 {
			tetromino.Diagram = CheckTetris(tetromino.Diagram, string(Char))
			Char++
			tetromino.Diagram = OptimizeTetris(tetromino.Diagram)
			Tetris = append(Tetris, tetromino)
			tetromino.Diagram = [][]string{}
		} else {
			tetromino.Diagram = append(tetromino.Diagram, strings.Split(line, ""))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	//Add last Tetris
	if len(tetromino.Diagram) > 0 {
		tetromino.Diagram = CheckTetris(tetromino.Diagram, string((Char)))
		tetromino.Diagram = OptimizeTetris(tetromino.Diagram)
		Tetris = append(Tetris, tetromino)
	}

	if len(Tetris) == 0 {
		fmt.Println("Enter at least one tetromino in the text file")
		os.Exit(0)
	}
}
