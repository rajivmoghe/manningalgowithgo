package main

import "fmt"

// Make a board filled with periods.
func make_chess_board(num_rows int) [][]string {
	num_cols := num_rows
	board := make([][]string, num_rows)
	for r := range board {
		board[r] = make([]string, num_cols)
		for c := 0; c < num_cols; c++ {
			board[r][c] = "."
		}
	}
	return board
}

// Display the board.
func dump_chess_board(board [][]string) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			fmt.Printf("%s ", board[r][c])
		}
		fmt.Println()
	}
}

// Return true if this series of squares contains at most one queen.
func series_is_legal(board [][]string, num_rows, r0, c0, dr, dc int) bool {
	num_cols := num_rows
	has_queen := false

	r := r0
	c := c0
	for {
		if board[r][c] == "Q" {
			// If we already have a queen on this row,
			// then this board is not legal.
			if has_queen {
				return false
			}

			// Remember that we have a queen on this row.
			has_queen = true
		}

		// Move to the next square in the series.
		r += dr
		c += dc

		// If we fall off the board, then the series is legal.
		if r >= num_rows ||
			c >= num_cols ||
			r < 0 ||
			c < 0 {
			return true
		}
	}
}

// Return true if the board is legal.
func board_is_legal(board [][]string, num_rows int) bool {
	// See if each row is legal.
	for r := 0; r < num_rows; r++ {
		if !series_is_legal(board, num_rows, r, 0, 0, 1) {
			return false
		}
	}

	// See if each column is legal.
	for c := 0; c < num_rows; c++ {
		if !series_is_legal(board, num_rows, 0, c, 1, 0) {
			return false
		}
	}

	// See if diagonals down to the right are legal.
	for r := 0; r < num_rows; r++ {
		if !series_is_legal(board, num_rows, r, 0, 1, 1) {
			return false
		}
	}
	for c := 0; c < num_rows; c++ {
		if !series_is_legal(board, num_rows, 0, c, 1, 1) {
			return false
		}
	}

	// See if diagonals down to the left are legal.
	for r := 0; r < num_rows; r++ {
		if !series_is_legal(board, num_rows, r, num_rows-1, 1, -1) {
			return false
		}
	}
	for c := 0; c < num_rows; c++ {
		if !series_is_legal(board, num_rows, 0, c, 1, -1) {
			return false
		}
	}

	// If we survived this long, then the board is legal.
	return true
}

// Return true if the board is legal and a solution.
func board_is_a_solution(board [][]string, num_rows int) bool {
	// See if the board contains exactly num_rows queens.
	num_queens := 0
	for r := 0; r < num_rows; r++ {
		for c := 0; c < num_rows; c++ {
			if board[r][c] == "Q" {
				num_queens += 1
			}
		}
	}

	// See if it is legal.
	if !board_is_legal(board, num_rows) {
		return false
	}
	return num_queens == num_rows
}

// Try placing a queen at position [r][c].
// Return true if we find a legal board.
func place_queens_1(board [][]string, num_rows, r, c int) bool {
	if r >= num_rows {
		// found := board_is_a_solution(board, num_rows)
		// if !found {
		// 	fmt.Println("Foll is Not a solution")
		// } else {
		// 	fmt.Println("Foll is Not a solution")
		// }
		// dump_chess_board(board)
		return board_is_a_solution(board, num_rows)
	}

	// Find the next square.
	next_r := r
	next_c := c + 1
	if next_c >= num_rows {
		next_r += 1
		next_c = 0
	}

	if place_queens_1(board, num_rows, next_r, next_c) {
		return true
	}

	board[r][c] = "Q"
	if place_queens_1(board, num_rows, next_r, next_c) {
		return true
	}

	board[r][c] = "."
	return false
}

// Try placing a queen at position [r][c].
// Return true if we find a legal board.
func place_queens_2(board [][]string, num_rows, r, c, num_placed int) bool {
	if r >= num_rows {
		if num_placed > num_rows {
			// fmt.Printf("Number of queens (%d) for rows (%d). Returning false.\n", num_placed, num_rows)
			return false
		}
		// found := board_is_a_solution(board, num_rows)
		// if !found {
		// 	fmt.Println("Foll is Not a solution")
		// } else {
		// 	fmt.Println("Foll is valid solution")
		// }
		// dump_chess_board(board)
		return board_is_a_solution(board, num_rows)
	}

	// Find the next square.
	next_r := r
	next_c := c + 1
	if next_c >= num_rows {
		next_r += 1
		next_c = 0
	}

	if place_queens_2(board, num_rows, next_r, next_c, num_placed) {
		return true
	}

	board[r][c] = "Q"
	if place_queens_2(board, num_rows, next_r, next_c, num_placed+1) {
		return true
	}

	board[r][c] = "."
	return false
}

// Set up and call place_queens_3.
func place_queens_3(board [][]string, num_rows, num_placed, r, c int) bool {
    // Make the num_attacking array.
    // The value num_attacking[r][c] is the number
    // of queens that can attack square (r, c).
    num_cols := num_rows
    num_attacking := make([][]int, num_rows)
    for r := range board {
        num_attacking[r] = make([]int, num_cols)
    }

    // Call do_place_queens_3.
    return do_place_queens_3(board, num_rows, num_placed, 0, 0, num_attacking)
}

// Try placing a queen at position [r][c].
// Keep track of the number of queens placed.
// Keep running totals of the number of queens attacking a square.
// Return true if we find a legal board.
func do_place_queens_3(board [][]string, num_rows, num_placed, r, c int, num_attacking [][]int) bool {
    return false
}