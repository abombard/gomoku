package main

const MAXDEPTH = 5

func getPossibleMoveList(b [][]int) []coord {

	var coords []coord

	for x := 0; x < HEIGHT; x++ {
		for y := 0; y < WIDTH; y++ {
			if isEmpty(b[x][y]) && (isPawnNearby2(b, x, y) || isBetweenPawn(b, x, y)) {
				coords = append(coords, coord{X: x, Y: y})
			}
		}
	}

	if len(coords) == 0 {
		for x := 0; x < HEIGHT; x++ {
			for y := 0; y < WIDTH; y++ {
				if isEmpty(b[x][y]) && isPawnNearby(b, x, y) {
					coords = append(coords, coord{X: x, Y: y})
				}
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

func recminmax(board [][]int, pt coord, player int, depth int, alpha, beta int, ch chan step, gameOver bool) step {

	next := getPossibleMoveList(board)
	// ERROR depth == MAXDEPTH && len(next) == 0 -> pt = shit
	if depth == MAXDEPTH && len(next) == 0 {
		return step{coord: coord{X: 10, Y: 10}}
		//	log.Fatal("depth == MAXDEPTH && len(next) == 0: GAME OVER")
	}
	if depth == 0 || gameOver || len(next) == 0 {
		if gameOver {
			depth += 1
		}
		score := (depth + 1) * heuristic2(board)
		ret := step{pt, score}
		if ch != nil {
			ch <- ret
		}
		return ret
	}

	var v step
	if player == current {
		v = step{score: -10000000}
	} else {
		v = step{score: 10000000}
	}

	addMove := func(b [][]int, nb *[][]int, c coord, gameOver *bool) error {
		err := move(b, c, player, nb)
		if err != nil {
			if err.Error() == "Game Over" {
				*gameOver = true
			} else {
				return err
			}
		}
		return nil
	}

	updateVScore := func(score int, c coord, done *bool) {
		*done = false
		if player == current {
			if score > v.score {
				v.coord = c
				v.score = score
			}
			if score > alpha {
				alpha = score
				if alpha >= beta {
					*done = true
				}
			}
		} else {
			if score < v.score {
				v.coord = c
				v.score = score
			}
			if score < beta {
				beta = score
				if alpha >= beta {
					*done = true
				}
			}
		}
	}

	if depth == MAXDEPTH {
		newch := make(chan step, len(next))

		k := 0
		for i := range next {

			gameOver := false

			b := boardCopy(board)

			err := addMove(b, &b, next[i], &gameOver)
			if err != nil {
				k++
				continue
			}

			go recminmax(b, next[i], (player+1)%2, depth-1, alpha, beta, newch, gameOver)
		}

		for i := 0; i < len(next)-k; i++ {
			tmp := <-newch

			var done bool
			updateVScore(tmp.score, tmp.coord, &done)
			if done {
				break
			}

		}

	} else {

		for i := range next {

			gameOver := false

			var newBoard [][]int
			err := addMove(board, &newBoard, next[i], &gameOver)
			if err != nil {
				continue
			}

			tmp := recminmax(newBoard, next[i], (player+1)%2, depth-1, alpha, beta, nil, gameOver)

			var done bool
			updateVScore(tmp.score, pt, &done)
			board[next[i].X][next[i].Y] = 0

			if done {
				break
			}
		}
	}

	if ch != nil {
		ch <- v
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

func countEnnemyPawns(b [][]int, player int) int {
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

	nb := countEnnemyPawns(b, player)

	v := recminmax(b, coord{0, 0, ""}, player, MAXDEPTH, -10000, 10000, nil, false)
	end := countEnnemyPawns(b, player)
	if end == nb+2 {
		g.Players[player].Score += 2
	}

	//log.Println("THE CHOOSEN ONE : ", v)

	return v.coord
}
