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

func inv(x, y int) (int, int) {
	return -x, -y
}

func assignMeScore(b, sb [][]int, x, y int, checks [][2]int) int {

	score := 0

	for i := range checks {
		x1, y1 := x+checks[i][0], y+checks[i][1]
		if isValidCoord(x1, y1) {
			if sb[x1][y1] >= score {
				score = sb[x1][y1]*sb[x1][y1] + 1
			}
		}
	}

	return score
}

func assignEnemyScore(b, sb [][]int, x, y int, checks [][2]int) int {

	score := 0

	for i := range checks {
		x1, y1 := x+checks[i][0], y+checks[i][1]
		if isValidCoord(x1, y1) {
			if sb[x1][y1] <= score {
				score = -(sb[x1][y1]*sb[x1][y1] + 1)
			}
		}
	}

	return score
}

func heuristic(board [][]int, player int) int {

	// score board
	sb := make([][]int, HEIGHT)
	for x := range board {
		sb[x] = make([]int, WIDTH)
	}

	log.Println("=== the board ===")
	printBoard(board)

	// checks from left
	checksLeft := [][2]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	checksRight := [][2]int{{1, 0}, {1, 1}, {-1, 0}, {-1, 1}}

	// heuristic score
	score := 0

	// left to right
	for x := range board {
		for y := range board[x] {
			if isEmptyNew(x, y, board, player) {
				sb[x][y] = 0
			} else {
				if isMeNew(x, y, board, player) {
					sb[x][y] = assignMeScore(board, sb, x, y, checksLeft)
				} else {
					sb[x][y] = assignEnemyScore(board, sb, x, y, checksLeft)
				}
				score += sb[x][y]
			}
		}
	}

	log.Println("==== first ===")
	printBoard(sb)

	// right to left
	for x := range board {
		for y := range board[x] {
			if isEmptyNew(x, y, board, player) {
				sb[x][y] = 0
			} else {
				if isMeNew(x, y, board, player) {
					sb[x][y] = assignMeScore(board, sb, x, y, checksRight)
				} else {
					sb[x][y] = assignEnemyScore(board, sb, x, y, checksRight)
				}
				score += sb[x][y]
			}
		}
	}

	return score
}
