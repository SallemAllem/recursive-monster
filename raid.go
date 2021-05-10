// Create a program that resolves a sudoku.

// A valid sudoku has only one possible solution.

// Usage
// Example 1:
// Example of output for one valid sudoku :
// $ go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
// 3 9 6 2 4 5 7 8 1$
// 1 7 8 3 6 9 5 2 4$
// 5 2 4 8 1 7 3 9 6$
// 2 8 7 9 5 1 6 4 3$
// 9 3 1 4 8 6 2 7 5$
// 4 6 5 7 2 3 9 1 8$
// 7 1 2 6 3 8 4 5 9$
// 6 5 9 1 7 4 8 3 2$
// 8 4 3 5 9 2 1 6 7$
// $
// Example 2:
// Examples of output for invalid input or sudokus :

// $ go run . 1 2 3 4 | cat -e
// Error$
// $ go run . | cat -e
// Error$
// $ go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
// Error$
// $



package main

import (
	"os"

	"github.com/01-edu/z01"
)

func ValidInput(args []string) bool {
	if len(args) != 9 {
		z01.PrintRune('E')
		z01.PrintRune('r')
		z01.PrintRune('r')
		z01.PrintRune('o')
		z01.PrintRune('r')
		z01.PrintRune('\n')
		return false
	}
	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			z01.PrintRune('E')
			z01.PrintRune('r')
			z01.PrintRune('r')
			z01.PrintRune('o')
			z01.PrintRune('r')
			z01.PrintRune('\n')
			return false
		}
	}
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			if (args[i])[j] != '.' && (args[i])[j] < '1' || (args[i])[j] > '9' {
				z01.PrintRune('E')
				z01.PrintRune('r')
				z01.PrintRune('r')
				z01.PrintRune('o')
				z01.PrintRune('r')
				z01.PrintRune('\n')
				return false
			}
		}
	}
	return true
}

func ValidSolution(sudoku *[9][9]rune, x int, y int, nbr rune) bool {
	// ищем совпадения по горизонтали
	for j := 0; j < 9; j++ {
		if sudoku[x][j] == nbr {
			return false // если ложь, значит нашли совпадения, и цифра не подходит
		}
	}
	// ищем совпадения по вертикали
	for i := 0; i < 9; i++ {
		if sudoku[i][y] == nbr {
			return false // если ложь, значит нашли совпадения, и цифра не подходит
		}
	}
	// ищем совпадения в квадрате 3х3
	height := x / 3
	width := y / 3
	for i := 3 * height; i < 3*(height+1); i++ {
		for j := 3 * width; j < 3*(width+1); j++ {
			if sudoku[i][j] == nbr {
				return false // если ложь, значит нашли совпадения, и цифра не подходит
			}
		}
	}
	return true // если правда, значит совпадений не найдено, и цифра подходит
}

func CheckDots(sudoku *[9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == '.' {
				return true // если правда, значит в матрице есть хоть одна точка
			}
		}
	}
	return false // если ложь, значит точек в матрице нет, судоку решено
}

func Solve(sudoku *[9][9]rune) {
	if CheckDots(sudoku) == false {
		return true // правда, если в матрице нет точекб судоку решено
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == '.' {
				for nbr := '1'; nbr <= '9'; nbr++ {
					if ValidSolution(sudoku, i, j, nbr) {
						sudoku[i][j] = nbr
						if Solve(sudoku) { // рекурсиный вызов функции, если не осталось точек, значит мы нашли решение
							return true
						}
					}
					sudoku[i][j] = '.'
				}
				return false // не подходит решение, ищем следующее
			}
		}
	}
	return false // решения нет
}

func main() {
	args := os.Args[1:]
	if ValidInput(args) {
		sudoku := [9][9]rune{}
		for i := range args {
			for j := range args[i] {
				sudoku[i][j] = rune((args[i])[j])
			}
		}
		if Solve(&sudoku) {
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					if j != 8 {
						z01.PrintRune(sudoku[i][j])
						z01.PrintRune((' '))
					} else {
						z01.PrintRune(sudoku[i][j])
					}
				}
				z01.PrintRune(('\n'))
			}
		} else {
			z01.PrintRune('E')
			z01.PrintRune('r')
			z01.PrintRune('r')
			z01.PrintRune('o')
			z01.PrintRune('r')
			z01.PrintRune('\n')
		}
	}
}
