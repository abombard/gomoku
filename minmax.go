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

	var newBoard [][]int

	err := move(board, pt, &player, &newBoard)
	if err != nil {
		var score int
		if err.Error() == "Game Over" {
			if player == current {
				score = (MAXDEPTH - depth) * 500
			} else {
				score = -(MAXDEPTH - depth) * 500
			}
		} else {
			score = -(MAXDEPTH + 1) * 500
		}
		return step{pt, score}
	}

	if len(newBoard) == 0 {
		newBoard = board
	}

	next := getPossibleMoveList(newBoard)
	if depth == 0 || len(next) == 0 {
		score := heuristic2(newBoard, player)
		board[pt.X][pt.Y] = 0
		return step{pt, score}
	}

	var v step
	if player != current {
		v = step{score: -10000}
		for i := range next {
			v = max(v, recminmax(newBoard, next[i], player, depth-1, alpha, beta))
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
			v = min(v, recminmax(newBoard, next[i], player, depth-1, alpha, beta))
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

func boardCopy(board [][]int) [][]int {
	b := make([][]int, HEIGHT)
	for i := 0; i < HEIGHT; i++ {
		b[i] = make([]int, WIDTH)
		for j := 0; j < WIDTH; j++ {
			b[i][j] = board[i][j]
		}
	}
	return b
}

func minmaxRoutine(board [][]int, pt coord, player int, ch chan step) {

	// create a new slice for each go routine
	b := boardCopy(board)

	ch <- recminmax(b, pt, player, MAXDEPTH, -10000, 10000)
}

func minmax(board [][]int, player int) coord {

	nextMoves := getPossibleMoveList(board)
	ch := make(chan step, len(nextMoves))

	for x := 0; x < len(nextMoves); x++ {

		go minmaxRoutine(board, nextMoves[x], player, ch)
	}

	v := step{score: -10000}
	for i := 0; i < len(nextMoves); i++ {
		tmp := <-ch
		v = max(v, tmp)
	}

	return v.coord
}
