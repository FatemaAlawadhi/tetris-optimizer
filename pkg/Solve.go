package optimizer

import (
	"math"
	"fmt"
)


func Solve() [][]string {
	// Calculate the initial size of the square grid based on the number of Tetris pieces
	SquareSize := int(math.Ceil(math.Sqrt(float64(4 * len(Tetris)))))

	// Create the square grid
	Solution = CreateSquare(SquareSize)

	// Keep increasing the size of the grid until a solution is found
	for !BacktrackSolver(Tetris, 0) {
		SquareSize++
		Solution = CreateSquare(SquareSize)
	}
	return Solution
}

func BacktrackSolver(Tetris []Tetromino, n int) bool {
	// Base case: if all Tetris pieces have been placed, return true
	if n == len(Tetris) {
		return true
	}

	// Try to insert the current Tetris piece at different positions in the grid
	for hor := 0; hor < len(Solution); hor++ {
		for ver := 0; ver < len(Solution); ver++ {
			// Check if the current Tetris piece can be inserted at position (hor, ver)
			if CheckInsert(hor, ver, Tetris[n].Diagram) {
				// Insert the Tetris piece into the grid
				Insert(hor, ver, Tetris[n].Diagram)
				// Recursively try to place the next Tetris piece
				if BacktrackSolver(Tetris, n+1) {
					return true
				}
				// If the next Tetris piece cannot be placed, remove the current one and try a different position
				Remove(hor, ver, Tetris[n].Diagram)
			}
		}
	}
	return false
}

func Insert(i, j int, tetris [][]string) {
	// Count the number of cells filled in the current Tetris piece
	hor, ver, HashPlaced := 0, 0, 0

	// Iterate over the Tetris piece diagram
	for hor < 4 {
		for ver < 4 {
			// If the current cell is not empty, fill the corresponding cell in the grid
			if tetris[hor][ver] != "." {
				HashPlaced++
				Solution[i+hor][j+ver] = tetris[hor][ver]

				// Break the loop if all cells in the Tetris piece have been filled
				if HashPlaced == 4 {
					break
				}
			}
			ver++
		}
		ver = 0
		hor++
	}
}

func Remove(i, j int,tetris [][]string  ) {
	// Count the number of cells filled in the current Tetris piece
	hor, ver, c := 0, 0, 0

	// Iterate over the Tetris piece diagram
	for hor < 4 {
		for ver < 4 {
			// If the current cell is not empty, remove the corresponding cell in the grid
			if tetris[hor][ver] != "." {
				if c == 4 {
					break
				}
				Solution[i+hor][j+ver] = "."
			}
			ver++
		}
		ver = 0
		hor++
	}
}

func CheckInsert(i, j int, tetris [][]string) bool {
	// Iterate over the Tetris piece diagram
	for hor := 0; hor < 4; hor++ {
		for ver := 0; ver < 4; ver++ {
			// Check if the current cell in the Tetris piece can be placed in the grid
			if tetris[hor][ver] != "."{
				// Check if the cell is out of bounds or already filled in the grid
				if i+hor == len(Solution) || j+ver == len(Solution) || Solution[i+hor][j+ver] != "." {
					return false
				}
			}
		}
	}
	return true
}

func CreateSquare(n int) [][]string {
	var Square [][]string
	var row []string

	// Create a square grid filled with empty cells
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		Square = append(Square, row)
		row = []string{}
	}
	return Square
}

func PrintSolution() {
	// Print the solution grid
	for i := range Solution {
		for j := range Solution {
			fmt.Printf("%s ", Solution[i][j])
		}
		fmt.Printf("\n")
	}
}
