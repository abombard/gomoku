package main

const MAXDEPTH = 3

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

type step struct {
	coord coord
	score int
}

func min(s1, s2 step) step {
	if s1.score <= s2.score {
		return s1
	}
	return s2
}

func max(s1, s2 step) step {
	if s1.score >= s2.score {
		return s1
	}
	return s2
}

func recminmax(board [][]int, pt coord, player int, depth int, alpha, beta int) step {

	player = (player + 1) % 2

	board[pt.X][pt.Y] = player + 1

	next := getPossibleMoveList(board)
	if depth == 0 || len(next) == 0 {
		score := heuristic2(board, player)
		board[pt.X][pt.Y] = 0
		return step{pt, score}
	}

	var v step
	if player != current {
		v = step{score: -10000}
		for i := range next {
			v = max(v, recminmax(board, next[i], player, depth-1, alpha, beta))
			if v.score > alpha {
				alpha = v.score
			}
			if alpha >= beta {
				break
			}
		}
	} else {
		v = step{score: 10000}
		for i := range next {
			v = min(v, recminmax(board, next[i], player, depth-1, alpha, beta))
			if v.score < beta {
				beta = v.score
			}
			if alpha >= beta {
				break
			}
		}
	}

	board[pt.X][pt.Y] = 0

	return step{pt, v.score}
}

func minmaxRoutine(board [][]int, pt coord, player int, ch chan step) {

	// create a new slice for each go routine
	b := make([][]int, HEIGHT)
	for i := 0; i < HEIGHT; i++ {
		b[i] = make([]int, WIDTH)
		for j := 0; j < WIDTH; j++ {
			b[i][j] = g.Board[i][j]
		}
	}

	ch <- recminmax(b, pt, player, MAXDEPTH, -10000, 10000)
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

		go minmaxRoutine(b, nextMoves[x], player, ch)
	}

	v := step{score: -10000}
	for i := 0; i < len(nextMoves); i++ {
		tmp := <-ch
		v = max(v, tmp)
	}

	return v.coord
}
