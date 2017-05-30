package main

import "log"

const MINMAX_ITER = 3

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

func recminmax(board [][]int, pt coord, player int, iter int) step {

	board[pt.X][pt.Y] = player + 1

	if iter == MINMAX_ITER {
		score := heuristic(board, player)
		board[pt.X][pt.Y] = 0
		return step{pt, score}
	}

	next := getPossibleMoveList(board)

	var scores []step
	for i := range next {
		s := recminmax(board, next[i], (player+1)%2, iter+1)
		scores = append(scores, step{coord: pt, score: s.score})
	}

	var ret step
	if iter%2 == 0 {
		ret = max(scores)
	} else {
		ret = min(scores)
	}

	board[pt.X][pt.Y] = 0

	return ret
}

func minmaxRoutine(board [][]int, pt coord, player int, ch chan step) {
	ret := recminmax(board, pt, player, 0)
	ch <- ret
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

	for x := 0; x < len(nextMoves); x++ {

		// create a new slice for each go routine
		b := make([][]int, HEIGHT)
		for i := 0; i < HEIGHT; i++ {
			b[i] = make([]int, WIDTH)
			for j := 0; j < WIDTH; j++ {
				b[i][j] = g.Board[i][j]
			}
		}

		go minmaxRoutine(b, nextMoves[x], player, ch)
	}

	res := make([]step, 0)
	for i := 0; i < len(nextMoves); i++ {
		tmp := <-ch
		res = append(res, tmp)
		log.Println(res)
	}

	ret := min(res)

	return ret.coord
}
