package main

import "log"

const MINMAX_ITER = 10

func getPossibleMoveList(b [][]int) []coord {
	var coords []coord

	for x := 0; x < HEIGHT; x++ {
		for y := 0; y < WIDTH; y++ {
			if isPawnNearby(b, x, y) == true && b[x][y] == 0 {
				coords = append(coords, coord{X: x, Y: y})
			}
		}
	}

	return coords
}

func isEmptyNew(x, y int, board [][]int, p int) bool {
	return board[x][y] == 0
}

func isMeNew(x, y int, board [][]int, player int) bool {
	return board[x][y] == player+1
}

func isEnemyNew(x, y int, board [][]int, player int) bool {
	return !isEmptyNew(x, y, board, player) && board[x][y] != player+1
}

func addScore(sb [][]int, x, y int, checks [][2]int) int {

	score := 0

	for i := range checks {
		x1, y1 := x+checks[i][0], y+checks[i][1]
		if isValidCoord(x1, y1) {
			score += sb[x1][y1] + 1
		}
	}

	return score
}

func assignScore(sb [][]int, x, y int, me bool, checks [][2]int) int {

	score := 0

	for i := range checks {
		x1, y1 := x+checks[i][0], y+checks[i][1]
		if isValidCoord(x1, y1) {
			if me && sb[x1][y1] > 0 {
				score += sb[x1][y1] + 1
			} else if !me && sb[x1][y1] < 0 {
				score += -sb[x1][y1]
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

type step struct {
	coord coord
	score int
}

func min(scores []step) step {
	ret := step{score: 10000}
	for i := range scores {
		if scores[i].score < ret.score {
			ret = scores[i]
		}
	}
	return ret
}

func max(scores []step) step {
	ret := step{score: -10000}
	for i := range scores {
		if scores[i].score > ret.score {
			ret = scores[i]
		}
	}
	return ret
}

func recminmax(board [][]int, pt coord, player int, iter int, ch chan step) step {

	board[pt.X][pt.Y] = player + 1

	if iter == MINMAX_ITER {
		score := heuristic(board, player)
		board[pt.X][pt.Y] = 0
		return step{pt, score}
	}

	next := getPossibleMoveList(board)

	var scores []step
	for i := range next {
		s := recminmax(board, next[i], (player+1)%2, iter+1, ch)
		scores = append(scores, step{coord: pt, score: s.score})
	}

	var ret step
	if player == current {
		ret = min(scores)
	} else {
		ret = max(scores)
	}

	board[pt.X][pt.Y] = 0

	ch <- ret

	return ret
}

func minmax(board [HEIGHT][WIDTH]int, player int) coord {

	b := make([][]int, HEIGHT)
	for i := 0; i < HEIGHT; i++ {
		b[i] = make([]int, WIDTH)
		for j := 0; j < WIDTH; j++ {
			b[i][j] = g.Board[i][j]
		}
	}

	nextMoves := getPossibleMoveList(b)
	ch := make(chan step, len(nextMoves))

	for i, _ := range nextMoves {

		// create a new slice for each go routine
		b := make([][]int, HEIGHT)
		for i := 0; i < HEIGHT; i++ {
			b[i] = make([]int, WIDTH)
			for j := 0; j < WIDTH; j++ {
				b[i][j] = g.Board[i][j]
			}
		}

		log.Println("=== send newB to go routine ===")
		log.Println(b)
		go recminmax(b, nextMoves[i], player, 0, ch)
	}

	res := make([]step, 0)
	for i := 0; i < len(nextMoves); i++ {
		tmp := <-ch
		res = append(res, tmp)
		log.Println(res)
	}

	ret := max(res)

	return ret.coord
}
