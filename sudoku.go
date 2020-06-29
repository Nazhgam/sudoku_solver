package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sudoku [9][9]int
var a [9][9]int = sudoku

func collectSudokuFromFile() {
	var i, j int
	file, err := os.Open("sudoku3.txt")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a, _ := strconv.Atoi(scanner.Text())
		if j == 9 {
			i++
			j = 0
		}
		sudoku[i][j] = a
		j++

	}
}

func square(n, i, j int) bool {
	modI := i - i%3
	modJ := j - j%3
	for x := modI; x <= modI+2; x++ {
		for k := modJ; k <= modJ+2; k++ {
			if (x != i && k != j) && sudoku[x][k] == n {
				return false
			}
		}
	}
	return true
}
func column(pos, n, i, j int) bool {
	if pos == 9 {
		return true
	}
	if pos == i {
		return column(pos+1, n, i, j)
	}
	if n != sudoku[pos][j] {
		return column(pos+1, n, i, j)
	}
	return false
}
func row(pos, n, i, j int) bool {
	if pos == 9 {
		return true
	}
	if pos == j {
		return row(pos+1, n, i, j)
	}
	if n != sudoku[i][pos] {
		return row(pos+1, n, i, j)
	}
	return false
}
func main() {

	collectSudokuFromFile()
	fillTheSudoku()
	printSudoku()
}
func isItFilled() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
func findZeroCase() (int, int) {
	var a, b int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == 0 {
				a = i
				b = j
				return i, j
			}
		}
	}
	return a, b
}
func fillTheSudoku() bool {
	if isItFilled() {
		return true
	}
	i, j := findZeroCase()
	for n := 1; n < 10; n++ {
		if isItPossibleToPut(n, i, j) {
			sudoku[i][j] = n
			if fillTheSudoku() {
				return true
			}
		}
		sudoku[i][j] = 0
	}
	return false
}
func isItPossibleToPut(n, i, j int) bool {

	if !row(0, n, i, j) || !column(0, n, i, j) || !square(n, i, j) {
		return false
	}
	return true
}
func printSudoku() {
	fmt.Println("\nanswer of Sudoku by Great and Powerful MAGZHAN\n")
	niz1 := 0
	niz2 := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if i == 3 || i == 6 {
				if i == 3 {
					niz1++
				}
				if i == 6 {
					niz2++
				}
				if niz1 == 1 || niz2 == 1 {
					fmt.Println()
					fmt.Println("_________________________________")
					fmt.Println()
				}
			}
			fmt.Print(" ", sudoku[i][j], " ")
			if j == 2 || j == 5 {
				fmt.Print(" | ")
			}

		}
		fmt.Println()
	}
}
