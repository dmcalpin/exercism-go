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
	// if row doesn't start with a symbol
	// can't have a win
	if currSymbol == "." {
		return -1
	}

	// check the next cell of this row
	if checkRow(currRow, cellIndex, currSymbol) {
		match := findMatch(board, rowIndex, cellIndex+1)
		if match != -1 {
			return match
		}
	}

	// check the next cell of the row above
	previousRowIndex := rowIndex - 1
	if previousRowIndex >= 0 {
		if checkRow(board[previousRowIndex], cellIndex, currSymbol) {
			match := findMatch(board, previousRowIndex, cellIndex+1)
			if match != -1 {
				return match
			}
		}
	}

	// check the next cell of the row below
	nextRowIndex := rowIndex + 1
	if nextRowIndex < len(board) {
		if checkRow(board[nextRowIndex], cellIndex, currSymbol) {
			match := findMatch(board, nextRowIndex, cellIndex)
			if match != -1 {
				return match
			}
		}
	}

	if cellIndex == len(currRow)-1 {
		// we made it to the end!
		return 999
	}
	return -1
}

func checkRow(row string, index int, symbol string) bool {
	if index == len(row) {
		return true
	} else if index < len(row)-1 {
		nextSymbol := string(row[index+1])
		// if match, move to that row
		if symbol == nextSymbol {
			return true
		}
	}
	return false
}
