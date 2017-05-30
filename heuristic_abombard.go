package main

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
				score += addScore(sb, x, y, checksLeft)
			} else {
				sb[x][y] = assignScore(sb, x, y,
					isMeNew(x, y, board, player),
					checksLeft)
			}
		}
	}

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

	return score
}
