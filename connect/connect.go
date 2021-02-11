package connect

type coordinates struct {
	Y int
	X int
}

// players 1 and 2
const (
	P1 = "X"
	P2 = "O"
)

// ResultOf calculates the result of a
// game of connect. Two loops are needed
// incase the board is not square (rows == cols)
func ResultOf(board []string) (string, error) {
	// Check the left-to-right win
	numRows := len(board)
	for i := 0; i < numRows; i++ {
		if findMatch(board, &coordinates{i, 0}, P1) {
			return P1, nil
		}
	}

	// Check the top-to-bottom win
	cellsPerRow := len(board[0])
	for i := 0; i < cellsPerRow; i++ {
		if findMatch(board, &coordinates{0, i}, P2) {
			return P2, nil
		}
	}

	return "", nil
}

// findMatch takes a coordinate and identifies any surrounding
// matches, it then follows the matches it finds, and marks
// any cells that have been checked. If it makes it's way across
// the board, 999 is returned
func findMatch(
	board []string,
	currCoord *coordinates,
	expectedSymbol string,
) (winner bool) {
	// if symbol is not a match, this isn't a win
	if !coordMatchSymbol(board, currCoord, expectedSymbol) {
		return false
	}

	// Check if at end of row
	// or column, return 999 for win
	switch expectedSymbol {
	case P2:
		rowsInBoard := len(board)
		if currCoord.Y+2 > rowsInBoard {
			return true
		}
	case P1:
		cellsInRow := len(board[currCoord.Y])
		if currCoord.X+2 > cellsInRow {
			return true
		}
	}

	markLookedAt(board, currCoord)

	// coordinates of all surrounding cells
	// on the board
	surroundingCells := []*coordinates{
		{Y: currCoord.Y - 1, X: currCoord.X},     // cell above right
		{Y: currCoord.Y - 1, X: currCoord.X + 1}, // cell above left
		{Y: currCoord.Y, X: currCoord.X + 1},     // cell to the right
		{Y: currCoord.Y, X: currCoord.X - 1},     // cell to the left
		{Y: currCoord.Y + 1, X: currCoord.X},     // cell below left
		{Y: currCoord.Y + 1, X: currCoord.X - 1}, // cell below right
	}

	for _, coord := range surroundingCells {
		match := findMatch(board, coord, expectedSymbol)
		if match {
			return true
		}
	}

	return false
}

// make sure we're on the board, and the symbol at the given
// coordinates matches the expected symbol
func coordMatchSymbol(board []string, coord *coordinates, symbol string) bool {
	return coord.Y >= 0 && coord.Y < len(board) && coord.X >= 0 && coord.X < len(board[0]) && string(board[coord.Y][coord.X]) == symbol
}

// Set the cell to '.' so we don't
// re-check it, thus preventing
// an infinite loop
func markLookedAt(board []string, coord *coordinates) {
	row := []byte(board[coord.Y])
	row[coord.X] = '.'
	board[coord.Y] = string(row)
}
