package main

import (
	"fmt"
	"strconv"
	"time"
)

var num_calls int64 = 0

const num_disks = 15

func main() {

	run_tower_hanoi()
}

func run_factorial() {
	var n int64
	for n = 0; n <= 21; n++ {
		fmt.Printf("%3d! = %20d\n", n, factorial(n))
	}
	fmt.Println()
}

func run_fibonacci() {
	for {
		// Get n as a string.
		var n_string string
		fmt.Printf("N: ")
		fmt.Scanln(&n_string)

		// If the n string is blank, break out of the loop.
		if len(n_string) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(n_string, 10, 64)
		fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))
	}
}

func run_fibonacci_on_the_fly() {
	for {
		// Get n as a string.
		var n_string string
		fmt.Printf("N: ")
		fmt.Scanln(&n_string)

		// If the n string is blank, break out of the loop.
		if len(n_string) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(n_string, 10, 64)
		fmt.Printf("fibonacci_on_the_fly(%d) = %d\n", n, fibonacci_on_the_fly(n))
	}
}

func run_fibonacci_prefilled() {
	for {
		// Get n as a string.
		var n_string string
		fmt.Printf("N: ")
		fmt.Scanln(&n_string)

		// If the n string is blank, break out of the loop.
		if len(n_string) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(n_string, 10, 64)
		fmt.Printf("fibonacci_prefilled(%d) = %d\n", n, fibonacci_prefilled(n))
	}
}

func knights_tour() {

	// Initialize the move offsets.
	initialize_offsets()

	// Try to find a tour.
	start := time.Now()
	st_rec := 1
	for st_row := 0; st_row < num_rows; st_row++ {
		for st_col := 0; st_col < num_cols; st_col++ {
			// Create the blank board.
			board := make_board(num_rows, num_cols)
			board[st_row][st_col] = 0
			start_iter := time.Now()
			fmt.Printf("[%d, %d]\t", st_row, st_col)
			if find_tour(board, num_rows, num_cols, st_row, st_col, st_rec) {
				elaps_iter := time.Since(start_iter)
				fmt.Printf("Started from %0d,%0d. Solved in %7d ms.\n", st_row, st_col, elaps_iter.Milliseconds())
				dump_board(board)
			} else {
				elaps_iter := time.Since(start_iter)
				fmt.Printf("No tour from %0d,%0d. Failed in %7d ms.\n", st_row, st_col, elaps_iter.Milliseconds())
			}
		}
	}
	elapsed := time.Since(start)
	// dump_board(board)
	fmt.Printf("\nTotal %f seconds\n", elapsed.Seconds())
	fmt.Printf("%d calls\n", num_calls)
}

func run_place_queens() {
	const num_rows = 26
	board := make_chess_board(num_rows)

	start := time.Now()
	// success := place_queens_1(board, num_rows, 0, 0)
	// success := place_queens_2(board, num_rows, 0, 0, 0)
	// success := place_queens_3(board, num_rows, 0, 0, 0)
	success := place_queens_4(board, num_rows, 0)

	elapsed := time.Since(start)
	if success {
		fmt.Println("Success!")
		dump_chess_board(board)
	} else {
		fmt.Println("No solution")
		// dump_chess_board(board)
	}
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}

func run_tower_hanoi() {

	// Make three posts.
	posts := [][]int{}

	// Push the disks onto post 0 biggest first.
	posts = append(posts, []int{})
	for disk := num_disks; disk > 0; disk-- {
		posts[0] = push(posts[0], disk)
	}

	// Make the other posts empty.
	for p := 1; p < 3; p++ {
		posts = append(posts, []int{})
	}

	// Draw the initial setup.
	draw_posts(posts)

	// Move the disks.
	move_disks(posts, num_disks, 0, 2, 1)

	fmt.Printf("Number of moves %d\n", num_calls)

}
