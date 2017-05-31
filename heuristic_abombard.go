package main

import "log"

func printBoard(b [][]int) {
	for x := 0; x < HEIGHT; x++ {
		log.Println(b[x])
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func addScore(sb [][]int, x, y int, checks [][2]int) int {

	score := 0

	for i := range checks {
		x1, y1 := x+checks[i][0], y+checks[i][1]
		if isValidCoord(x1, y1) {
			if sb[x1][y1] > 0 {
				score += sb[x1][y1] + 1
			} else {
				score += sb[x1][y1] - 1
			}
		}
	}

	return score
}

func assignScore(sb [][]int, x, y int, me bool, checks [][2]int) int {

	score := 0

	for i := range checks {
		x1, y1 := x+checks[i][0], y+checks[i][1]
		if isValidCoord(x1, y1) {
			if abs(sb[x1][y1]) >= abs(score) {
				if me && sb[x1][y1] >= 0 {
					score = sb[x1][y1] + 1
				} else if !me && sb[x1][y1] <= 0 {
					score = sb[x1][y1] - 1
				}
			}
		}
	}

	return score
}

func heuristic(board [][]int, player int) int {

	// heuristic score
	score := 0

	// score board
	sb := make([][]int, HEIGHT)
	for x := range board {
		sb[x] = make([]int, WIDTH)
	}

	// checks from left
	checksLeft := [][2]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	checksRight := [][2]int{{1, 0}, {1, 1}, {-1, 0}, {-1, 1}}

	// left to right
	for x := range board {
		for y := range board[x] {
			if isEmptyNew(x, y, board, player) {
				sb[x][y] = 0
			} else {
				sb[x][y] = assignScore(sb, x, y,
					isMeNew(x, y, board, player),
					checksLeft)
			}
		}
	}

	log.Println("==== left to right ===")
	printBoard(sb)

	// right to left
	for x := range board {
		for y := range board[x] {
			if isEmptyNew(x, y, board, player) {
				sb[x][y] = 0
				score += addScore(sb, x, y, checksRight)
			} else {
				sb[x][y] = assignScore(sb, x, y,
					isMeNew(x, y, board, player),
					checksRight)
			}
		}
	}

	log.Println("==== right to ileft ===")
	printBoard(sb)

	return score
}
