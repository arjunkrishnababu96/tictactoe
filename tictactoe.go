package tictactoe

import (
	"fmt"
	"strings"
)

func IsValidBoard(b string) bool {
	board := strings.ToUpper(b)

	if !hasValidCharsOnly(board) {
		return false
	}

	if len(board) != 11 {
		return false
	}

	return true
}

func IsWinnable(board string, my_symbol byte, indices [3]int) (bool, []int, error) {
	var count_mysym, count_oppsym int

	opp_symbol := getOpponentSymbol(my_symbol)

	for _, index := range indices {
		if !isValidIndex(index) {
			return false, nil, fmt.Errorf("IsWinnable: index %d is outside the board")
		}

		switch board[index] {
		case my_symbol:
			count_mysym++
		case opp_symbol:
			count_oppsym++
		}
	}

	if count_mysym == 2 && count_oppsym == 0 {
		return true, indices[:], nil
	}
	return false, nil, nil
}

func HasWon(b string, symbol byte) bool {
	//check if somebody has won

	patterns := [][]int{
		// horizontals
		{0, 1, 2},
		{4, 5, 6},
		{8, 9, 10},

		// verticals
		{0, 4, 8},
		{1, 5, 9},
		{2, 6, 10},

		// diagonals
		{0, 5, 10},
		{2, 5, 8},
	}

	for _, index := range patterns {
		if b[index[0]] == symbol && b[index[1]] == symbol && b[index[2]] == symbol {
			// won!
			return true
		}
	}

	return false
}

func getOpponentSymbol(ch byte) byte {
	switch ch {
	case 'X':
		return 'O'
	case 'O':
		return 'X'
	default:
		return '-'
	}
}

func hasValidCharsOnly(board string) bool {
	for i, ch := range board {
		if i == 3 || i == 7 {
			if ch != '|' {
				return false
			}
		} else {
			if ch != 'X' && ch != 'O' && ch != '-' {
				return false
			}
		}
	}
	return true
}

func isValidIndex(i int) bool {
	// 0123456789X
	// ---|---|---
	if i >= 0 && i <= 10 {
		return true
	} else {
		return false
	}
}

func isIndicesDifferent(idx [3]int) bool {
	if idx[0] == idx[1] || idx[0] == idx[2] || idx[1] == idx[2] {
		return false
	} else {
		return true
	}
}
