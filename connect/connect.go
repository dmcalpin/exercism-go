package connect

import (
	"fmt"
)

type coordinates struct {
	X int
	Y int
}

// ResultOf calculates the result of a
// game of connect
func ResultOf(board []string) (string, error) {
	// Check the left-to-right win
	numRows := len(board)
	for i := 0; i < numRows; i++ {
		fmt.Printf("Starting row: %d\n", i)
		nextRow := findLRMatch(board, i, 0, "X")
		if nextRow != -1 {
			return "X", nil
		}
	}

	// Check the top-to-bottom win
	cellsPerRow := len(board[0])
	for i := 0; i < cellsPerRow; i++ {
		fmt.Printf("Starting cell: %d\n", i)
		nextCol := findVertMatch(board, 0, i, "O")
		if nextCol != -1 {
			return "O", nil
		}
	}

	return "", nil
}

func findLRMatch(board []string, rowIndex int, cellIndex int, expectedSymbol string) (nextRow int) {
	// get the current symbol
	currSymbol := string(board[rowIndex][cellIndex])

	// If not expected symbol, fail
	if currSymbol != expectedSymbol {
		return -1
	}

	// if we're at the end we return
	// 999 as an arbitrary success value
	cellsInRow := len(board[rowIndex])
	if cellIndex+2 > cellsInRow {
		// we made it to the end!
		return 999
	}

	markLookedAt(board, rowIndex, cellIndex)

	surroundingCells := []coordinates{
		{Y: rowIndex - 1, X: cellIndex},     // cell above right
		{Y: rowIndex - 1, X: cellIndex + 1}, // cell above left
		{Y: rowIndex, X: cellIndex + 1},     // cell to the right
		{Y: rowIndex, X: cellIndex - 1},     // cell to the left
		{Y: rowIndex + 1, X: cellIndex},     // cell below left
		{Y: rowIndex + 1, X: cellIndex - 1}, // cell below right
	}

	for _, coord := range surroundingCells {
		if match := checkRow(board, coord.Y, coord.X, currSymbol, expectedSymbol); match != -1 {
			return match
		}
	}

	return -1
}

func checkRow(board []string, rowIndex int, cellIndex int, symbol string, expectedSymbol string) int {
	// make sure we're on the board
	if rowIndex >= 0 && rowIndex < len(board) && cellIndex >= 0 && cellIndex < len(board[0]) {
		// if match, move to that row
		if symbol == string(board[rowIndex][cellIndex]) {
			return findLRMatch(board, rowIndex, cellIndex, expectedSymbol)
		}
	}
	return -1
}

func findVertMatch(board []string, rowIndex int, cellIndex int, expectedSymbol string) (nextRow int) {
	// get the current symbol
	currSymbol := string(board[rowIndex][cellIndex])

	// if row doesn't start with a symbol
	// can't have a win
	if currSymbol != expectedSymbol {
		return -1
	}

	// if we're at the end we return
	// 999 as an arbitrary success value
	rowsInBoard := len(board)
	if rowIndex+2 > rowsInBoard {
		// we made it to the end!
		return 999
	}

	markLookedAt(board, rowIndex, cellIndex)

	surroundingCells := []coordinates{
		{Y: rowIndex - 1, X: cellIndex},     // cell above right
		{Y: rowIndex - 1, X: cellIndex + 1}, // cell above left
		{Y: rowIndex, X: cellIndex + 1},     // cell to the right
		{Y: rowIndex, X: cellIndex - 1},     // cell to the left
		{Y: rowIndex + 1, X: cellIndex},     // cell below left
		{Y: rowIndex + 1, X: cellIndex - 1}, // cell below right
	}

	for _, coord := range surroundingCells {
		if match := checkCol(board, coord.Y, coord.X, currSymbol, expectedSymbol); match != -1 {
			return match
		}
	}

	return -1
}

func checkCol(board []string, rowIndex int, cellIndex int, symbol string, expectedSymbol string) int {
	// make sure we're on the board
	if rowIndex >= 0 && rowIndex < len(board) && cellIndex >= 0 && cellIndex < len(board[0]) {
		// if match, move to that row
		if rowIndex < len(board) {
			if symbol == string(board[rowIndex][cellIndex]) {
				return findVertMatch(board, rowIndex, cellIndex, expectedSymbol)
			}
		}
	}
	return -1
}

func markLookedAt(board []string, rowIndex, cellIndex int) {
	row := []byte(board[rowIndex])
	row[cellIndex] = '.'
	board[rowIndex] = string(row)
}
