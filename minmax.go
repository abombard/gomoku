package main

const MAXDEPTH = 5

func getPossibleMoveList(b [][]int) []coord {

	var coords []coord

	// Get all pawns after or between 2 pawns of same color
	for x := 0; x < HEIGHT; x++ {
		for y := 0; y < WIDTH; y++ {
			if isEmpty(b[x][y]) && (isPawnNearby2(b, x, y) || isBetweenPawn(b, x, y)) {
				coords = append(coords, coord{X: x, Y: y})
			}
		}
	}

	// Get all pawns near another pawn
	if len(coords) == 0 {
		for x := 0; x < HEIGHT; x++ {
			for y := 0; y < WIDTH; y++ {
				if isEmpty(b[x][y]) && isPawnNearby(b, x, y) {
					coords = append(coords, coord{X: x, Y: y})
				}
			}
		}
	}

	// First move
	if len(coords) == 0 {
		coords = append(coords, coord{X: 10, Y: 10})
	}

	return coords
}

type step struct {
	coord coord
	score int
}

func maxstep(s1, s2 step) step {
	if s1.score >= s2.score {
		return s1
	}
	return s2
}

func minstep(s1, s2 step) step {
	if s1.score <= s2.score {
		return s1
	}
	return s2
}

func max(n1, n2 int) int {
	if n1 >= n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 <= n2 {
		return n1
	}
	return n2
}

func recminmax(board [][]int, player, depth, alpha, beta int, gameOver bool) step {

	if depth == 0 || gameOver {
		score := (depth + 1) * heuristic2(board, g.Current)
		return step{score: score}
	}

	addMove := func(b [][]int, nb *[][]int, c coord, gameOver *bool) error {
		pCapturedCount := 0
		err := move(b, c, player, nb, &pCapturedCount)
		if err != nil {
			if err.Error() == "Game Over" {
				*gameOver = true
			} else {
				return err
			}
		}
		return nil
	}

	var v step
	if player == g.Current {
		v = step{score: -10000000}
	} else {
		v = step{score: 10000000}
	}

	nextMoves := getPossibleMoveList(board)
	for i := range nextMoves {

		var newBoard [][]int
		gameOver := false

		err := addMove(board, &newBoard, nextMoves[i], &gameOver)
		if err != nil {
			continue
		}

		s := recminmax(newBoard, (player+1)%2, depth-1, alpha, beta, gameOver)
		s.coord = nextMoves[i]

		board[nextMoves[i].X][nextMoves[i].Y] = 0

		if player == g.Current {
			v = maxstep(v, s)
			alpha = max(alpha, s.score)
		} else {
			v = minstep(v, s)
			beta = min(beta, s.score)
		}

		if alpha >= beta {
			break
		}
	}

	return v
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

func countEnemyPawns(b [][]int, player int) int {
	count := 0
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			if isEnemy(b[x][y], player) {
				count++
			}
		}
	}
	return count
}

func minmax(board [][]int, player int) coord {

	b := boardCopy(board)

	v := recminmax(b, player, MAXDEPTH, -10000, 10000, false)

	//log.Println("THE CHOOSEN ONE : ", v)

	return v.coord
}
