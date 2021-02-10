package connect

import (
	"fmt"
)

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
	// 997 as an arbitrary success value
	cellsInRow := len(board[rowIndex])
	if cellIndex+2 > cellsInRow {
		// we made it to the end!
		return 999
	}

	markLookedAt(board, rowIndex, cellIndex)

	fmt.Printf("searching for: %s, row: %d, cell: %d\n", currSymbol, rowIndex, cellIndex)

	// check the next cell of this row
	if match := checkRow(board, rowIndex, cellIndex+1, currSymbol, expectedSymbol); match != -1 {
		return match
	}

	// check the next cell of the row above
	previousRowIndex := rowIndex - 1
	if match := checkRow(board, previousRowIndex, cellIndex+1, currSymbol, expectedSymbol); match != -1 {
		return match
	}

	// check the next cell of the row below
	nextRowIndex := rowIndex + 1
	// note we do not increment cellIndex
	// because of the diagonal shape of the board
	if match := checkRow(board, nextRowIndex, cellIndex, currSymbol, expectedSymbol); match != -1 {
		return match
	}

	// check the previous cell of the row below
	// note we do not increment cellIndex
	// because of the diagonal shape of the board
	if match := checkRow(board, nextRowIndex, cellIndex-1, currSymbol, expectedSymbol); match != -1 {
		return match
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

	fmt.Printf("searching for: %s, row: %d, cell: %d\n", currSymbol, rowIndex, cellIndex)

	// check the row of this col
	if match := checkCol(board, rowIndex+1, cellIndex, currSymbol, expectedSymbol); match != -1 {
		return match
	}

	// check the next row of the col to the left
	previousCellIndex := cellIndex - 1
	if match := checkCol(board, rowIndex+1, previousCellIndex, currSymbol, expectedSymbol); match != -1 {
		return match
	}

	// check the same row, cell to the right
	nextCellIndex := cellIndex + 1
	if match := checkCol(board, rowIndex, nextCellIndex, currSymbol, expectedSymbol); match != -1 {
		return match
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
