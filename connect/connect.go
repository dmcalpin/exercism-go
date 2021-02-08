package connect

import "fmt"

// ResultOf calculates the result of a
// game of connect
func ResultOf(board []string) (string, error) {
	// Check the left-to-right win
	for i := 0; i < len(board); i++ {
		cellIndex := 0
		rowIndex := i
		// loop over any susequent rows
		// where a match is found
		fmt.Printf("Starting row: %d\n", rowIndex)
		nextRow := findLRMatch(board, rowIndex, cellIndex)
		if nextRow != -1 {
			return string(board[rowIndex][cellIndex]), nil
		}
	}

	// Check the top-to-bottom win
	for i := 0; i < len(board[0]); i++ {
		cellIndex := i
		rowIndex := 0
		// loop over any susequent rows
		// where a match is found
		fmt.Printf("Starting cell: %d\n", cellIndex)
		nextCol := findVertMatch(board, rowIndex, cellIndex)
		if nextCol != -1 {
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

func findLRMatch(board []string, rowIndex int, cellIndex int) (nextRow int) {
	// if we're at the end we return
	// 999 as an arbitrary success value
	if cellIndex >= len(board[rowIndex])-1 {
		// we made it to the end!
		return 997
	}

	// get the current symbol
	currSymbol := string(board[rowIndex][cellIndex])

	// if row doesn't start with a symbol
	// can't have a win
	if currSymbol != "X" {
		return -1
	}

	fmt.Printf("searching for: %s, row: %d, cell: %d\n", currSymbol, rowIndex, cellIndex)

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
		return findLRMatch(board, rowIndex, cellIndex)
	}
	return -1
}

func findVertMatch(board []string, rowIndex int, cellIndex int) (nextRow int) {
	// if we're at the end we return
	// 999 as an arbitrary success value
	if rowIndex == len(board) {
		// we made it to the end!
		return 999
	}

	if cellIndex == len(board[0]) {
		// we made it to the end!
		return 998
	}

	// get the current symbol
	currSymbol := string(board[rowIndex][cellIndex])

	// if row doesn't start with a symbol
	// can't have a win
	if currSymbol != "O" {
		return -1
	}

	fmt.Printf("searching for: %s, row: %d, cell: %d\n", currSymbol, rowIndex, cellIndex)

	// check the row of this col
	if match := checkCol(board, rowIndex+1, cellIndex, currSymbol); match != -1 {
		return match
	}

	// check the next row of the col to the left
	previousCellIndex := cellIndex - 1
	if previousCellIndex >= 0 {
		if match := checkCol(board, rowIndex+1, previousCellIndex, currSymbol); match != -1 {
			return match
		}
	}

	// check the same row, cell to the right
	nextCellIndex := cellIndex + 1
	if nextCellIndex < len(board[0]) {
		// note we do not increment cellIndex
		// because of the diagonal shape of the board
		if match := checkCol(board, rowIndex, nextCellIndex, currSymbol); match != -1 {
			return match
		}
	}

	// we're at the end!
	if rowIndex == len(board)-1 {
		return 997
	}

	return -1
}

func checkCol(board []string, rowIndex int, cellIndex int, symbol string) int {
	// if match, move to that row
	if rowIndex < len(board) {
		if symbol == string(board[rowIndex][cellIndex]) {
			return findVertMatch(board, rowIndex, cellIndex)
		}
	}
	return -1
}
