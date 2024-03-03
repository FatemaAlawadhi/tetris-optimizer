package optimizer

func OptimizeTetris(tetris [][]string) [][]string {
	i := 0

	// Shift the tetris piece vertically until the top row contains at least one filled cell
	for {
		zeroes := 0
		for j := 0; j < 4; j++ {
			if tetris[i][j] == "." {
				zeroes++
			}
		}
		if zeroes == 4 {
			tetris = ShiftVertical(tetris)
			continue
		}
		break
	}

	// Shift the tetris piece horizontally until the leftmost column contains at least one filled cell
	for {
		zeroes := 0
		for j := 0; j < 4; j++ {
			if tetris[j][i] == "." {
				zeroes++
			}
		}
		if zeroes == 4 {
			tetris = ShiftHorizontal(tetris)
			continue
		}
		break
	}

	return tetris
}

func ShiftVertical( tetris [][]string) [][]string {
	// Shift the tetris piece vertically by swapping rows
	temp := tetris[0]
	tetris[0] = tetris[1]
	tetris[1] = tetris[2]
	tetris[2] = tetris[3]
	tetris[3] = temp
	return tetris
}

func ShiftHorizontal(tetris [][]string) [][]string {
	// Shift the tetris piece horizontally by transposing, shifting vertically, and transposing back
	tetris = Transpose(tetris)
	tetris = ShiftVertical(tetris)
	tetris = Transpose(tetris)
	return tetris
}

func Transpose(tetris [][]string) [][]string {
	// Transpose the tetris piece matrix by swapping rows and columns
	xl := len(tetris[0])
	yl := len(tetris)
	result := make([][]string, xl)
	for i := 0; i < xl; i++ {
		result[i] = make([]string, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = tetris[j][i]
		}
	}
	return result
}