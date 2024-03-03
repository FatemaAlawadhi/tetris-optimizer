package optimizer

import (
	"fmt"
	"os"
)

func CheckTetris(Tetris [][]string, Letter string) [][]string {
	//Check Vertical length
	Hash := 0
	Connections := 0

	if len(Tetris) > 4 {
		fmt.Println("ERROR")
		os.Exit(0)
	}

	for hor := 0; hor < len(Tetris); hor++ {
		//Check Horizontal Length
		if len(Tetris[hor]) > 4 {
			fmt.Println("ERROR")
			os.Exit(0)
		}

		for ver := 0; ver < len(Tetris[hor]); ver++ {
			if Tetris[hor][ver] == "#" {
				Hash++
				//Valid shapes has 6 or 8 conncections
				Connections = Connections + CheckConnections(hor, ver, Tetris)
				//Convert # to a letter
				Tetris[hor][ver] = Letter
				//Check used char
			} else if Tetris[hor][ver] != "#" && Tetris[hor][ver] != "." {
				fmt.Println("ERROR")
				os.Exit(0)
			}
		}
	}

	//Check Num of Hash
	if Hash != 4 {
		fmt.Println("ERROR")
		os.Exit(0)
	}

	//Check Shape Validity
	if Connections != 6 && Connections != 8 {
		fmt.Println("ERROR")
		os.Exit(0)
	}

	return Tetris
}

func CheckConnections(HorInd int, VerInd int, Tetris [][]string) int {
	Connections := 0
	if HorInd < 3 && Tetris[HorInd+1][VerInd] != "." {
		Connections++
	}
	if HorInd > 0 && Tetris[HorInd-1][VerInd] != "." {
		Connections++
	}
	if VerInd < 3 && Tetris[HorInd][VerInd+1] != "." {
		Connections++
	}
	if VerInd > 0 && Tetris[HorInd][VerInd-1] != "." {
		Connections++
	}
	return Connections
}
