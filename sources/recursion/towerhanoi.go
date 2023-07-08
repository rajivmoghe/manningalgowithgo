package main

import "fmt"

// Add a disk to beginning of the post.
func push(post []int, disk int) []int {
	return append([]int{disk}, post...)
}

// Remove the first disk from the post.
// Return that disk and the revised post.
func pop(post []int) (int, []int) {
	return post[0], post[1:]
}

// Move one disk from from_post to to_post.
func move_disk(posts [][]int, from_post, to_post int) {
	var disk int
	num_calls++
	disk, posts[from_post] = pop(posts[from_post])
	posts[to_post] = push(posts[to_post], disk)
}

// Draw the posts by showing the size of the disk at each level.
func draw_posts(posts [][]int) {
    // Add 0s to the end of each post so they all have num_disks entries.
    for p := 0; p < 3; p++ {
        for len(posts[p]) < num_disks {
            posts[p] = push(posts[p], 0)
        }
    }

    // Draw the posts.
    for row := 0; row < num_disks; row++ {
        // Draw this row.
        for p := 0; p < 3; p++ {
            // Draw the disk on post p's row.
            fmt.Printf("\t%d ", posts[p][row])
        }
        fmt.Println()
    }
    // Draw a line between moves.
    fmt.Println("\t\t-----")

    // Remove the 0s.
    for p := 0; p < 3; p++ {
        for len(posts[p]) > 0 && posts[p][0] == 0 {
            _, posts[p] = pop(posts[p])
        }
    }
}

// Move the disks from from_post to to_post
// using temp_post as temporary storage.
func move_disks(posts [][]int, numToMove, from, to, via int) {

	if numToMove > 1 {
		move_disks(posts, numToMove-1, from, via, to)
	}

	move_disk(posts, from, to)
	draw_posts(posts)

	if numToMove > 1 {
		move_disks(posts, numToMove-1, via, to, from)
	}

}
