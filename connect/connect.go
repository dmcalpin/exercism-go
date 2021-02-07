package connect

import "fmt"

// ResultOf calculates the result of a
// game of connect
func ResultOf(board []string) (string, error) {
	// Iterate over each starting row
	for i := 0; i < len(board); i++ {
		cellIndex := 0
		rowIndex := i
		// loop over any susequent rows
		// where a match is found
		fmt.Printf("Starting row: %d\n", rowIndex)
		nextRow := findMatch(board, rowIndex, cellIndex)
		if nextRow != -1 {
			return string(board[rowIndex][cellIndex]), nil
		}
	}

	if board[0] == "O" {
		return "O", nil
	}
	if board[0] == "X" {
		return "X", nil
	}
	return "", nil
}

func findMatch(board []string, rowIndex int, cellIndex int) (nextRow int) {
	// get the current symbol
	currRow := board[rowIndex]
	currSymbol := string(currRow[cellIndex])

	fmt.Printf("searching for: %s, row: %d, cell: %d\n", currSymbol, rowIndex, cellIndex)

	// if we're at the end we return
	// 999 as an arbitrary success value
	if cellIndex == len(currRow)-1 {
		// we made it to the end!
		return 999
	}
	// if row doesn't start with a symbol
	// can't have a win
	if currSymbol == "." {
		return -1
	}

	// check the next cell of this row
	if match := checkRow(board, rowIndex, cellIndex+1, currSymbol); match != -1 {
		return match
	}

	// check the next cell of the row above
	previousRowIndex := rowIndex - 1
	if previousRowIndex >= 0 {
		if match := checkRow(board, previousRowIndex, cellIndex+1, currSymbol); match != -1 {
			return match
		}
	}

	// check the next cell of the row below
	nextRowIndex := rowIndex + 1
	if nextRowIndex < len(board) {
		// note we do not increment cellIndex
		// because of the diagonal shape of the board
		if match := checkRow(board, nextRowIndex, cellIndex, currSymbol); match != -1 {
			return match
		}
	}

	return -1
}

func checkRow(board []string, rowIndex int, cellIndex int, symbol string) int {
	// if match, move to that row
	if symbol == string(board[rowIndex][cellIndex]) {
		return findMatch(board, rowIndex, cellIndex)
	}
	return -1
}
