package main

import (
	"fmt"
)

// The board dimensions.
const num_rows = 8
const num_cols = num_rows

// Whether we want an open or closed tour.
const require_closed_tour = false

// Value to represent a square that we have not visited.
const unvisited = -1

// Define offsets for the knight's movement.
type Offset struct {
	dr, dc int
}

var move_offsets []Offset

// var num_calls int64

// Fill the Offset slice.
func initialize_offsets() {
	move_offsets = []Offset{
		{-2, -1},
		{-1, -2},
		{+2, -1},
		{+1, -2},
		{-2, +1},
		{-1, +2},
		{+2, +1},
		{+1, +2},
	}
}

// Make a board filled with -1s.
func make_board(num_rows, num_cols int) [][]int {
	board := make([][]int, num_rows)
	for r := range board {
		board[r] = make([]int, num_cols)
		for c := 0; c < num_cols; c++ {
			board[r][c] = unvisited
		}
	}
	return board
}

func dump_board(board [][]int) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			fmt.Printf("%02d ", board[r][c])
		}
		fmt.Println()
	}
}

func is_allowed(move Offset) bool {
	return move.dc > 0 && move.dc < num_cols && move.dr > 0 && move.dr < num_rows
}

// Try to extend a knight's tour starting at (start_row, start_col).
// Return true or false to indicate whether we have found a solution.
func find_tour(board [][]int, num_rows, num_cols, cur_row, cur_col, num_visited int) bool {
	num_calls += 1

	if num_visited == num_rows*num_cols { // visited all squares, check
		if !require_closed_tour { // no closed tour - all is well
			return true
		} else { // check if next valid step leads to first square numbered 0
			// make a list of possible next moves
			next_moves := make([]Offset, 0)
			for _, nxt_stp := range move_offsets {
				next_moves = append(next_moves, Offset{cur_row + nxt_stp.dr, cur_col + nxt_stp.dc})
			}
			for _, move := range next_moves {
				if !(move.dc > 0 && move.dc < num_cols && move.dr > 0 && move.dr < num_rows) {
					continue
				}
				if board[move.dr][move.dc] == 0 {
					return true
				}
			}

			return false
		}
	} else {
		// for each move, check if it is a solution or leads to one
		// 1. out of board - ditch the move
		// 2. landing on already filled square - ditch the move
		for _, offset := range move_offsets {
			nxt_row := cur_row + offset.dr
			nxt_col := cur_col + offset.dc

			// within bounds of board AND is unvisited
			if nxt_row < 0 || nxt_row >= num_rows || nxt_col < 0 || nxt_col >= num_cols || board[nxt_row][nxt_col] != unvisited {
				continue
			}
			// 3. landed on a new square -
			//   a. mark spot visited
			//   b. recurse find_tour from that spot!
			board[nxt_row][nxt_col] = num_visited
			if find_tour(board, num_rows, num_cols, nxt_row, nxt_col, num_visited+1) {
				return true
			}

			// all forwards are exhausted - this is a dead end.
			board[nxt_row][nxt_col] = unvisited

		}
		return false
	}
}
