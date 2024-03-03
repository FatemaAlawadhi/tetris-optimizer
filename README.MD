## tetris-optimizer

### Objectives

Develop a program that receives only one argument, a path to a text file which will contain a list of [tetrominoes](https://en.wikipedia.org/wiki/Tetromino) and assemble them in order to create the smallest square possible.

### Instructions

The program must :

- Compile successfully
- Assemble all of the tetrominoes in order to create the smallest square possible
- Identify each tetromino in the solution by printing them with uppercase latin letters (`A` for the first one, `B` for the second, etc)
- Expect at least one tetromino in the text file
- In case of bad format on the tetrominoes or bad file format it should print `ERROR`
- The project must be written in **Go**.
- The code must respect the [**good practices**](../good-practices/README.md).

#### Example of a text File

```
#...
#...
#...
#...

....
....
..##
..##
```

- If it isn't possible to form a complete square, the program should leave spaces between the tetrominoes. For example:

```
ABB.
ABB.
A...
A...
```

### USAGE

```
go run . FileName.txt
```

to test the programm with given bad and good input files, you can run test.sh by
`bash test.sh`

### Example of the output

```console
$ cat -e sample.txt
...#$
...#$
...#$
...#$
$
....$
....$
....$
####$
$
.###$
...#$
....$
....$
$
....$
..##$
.##.$
....$
$
....$
.##.$
.##.$
....$
$
....$
....$
##..$
.##.$
$
##..$
.#..$
.#..$
....$
$
....$
###.$
.#..$
....$

$ go run . sample.txt | cat -e
ABBBB.$
ACCCEE$
AFFCEE$
A.FFGG$
HHHDDG$
.HDD.G$
$
```

### Logic used to implement the project

1. Scan the file and save each tetromino in  an [][]string (ReadFile.go)

2. Check the below conditions (CheckTetris.go)
   - A tetromino has an area of 4 squares (4 Hash)
   - Each tetromino is represented in a 4*4 square
   - Each tetromino has 6 or 8 connections. Otherwise it is not a valid tetromino shape.
   - Each tetromino is represented using only "#" and "."

3. Optimize the tetromino until it becomes at the top left of the square (Optimize.go)
   - By shifting vertically
   - By shifting horizontally (Transpose,shift vertical and transpose again)


3. Add the tetromino to Tetris []Tetromino

4. Solve (Solve.go):
   - Calculate the size of the square and create it.
   - Solve using Back Track solver technique 
   - If there isnt any possible solution increase the size of the square and repeat

### AUTHORS

- Fatema Alawadhi

### LICENSES

This program developed within the scope of Reboot.

