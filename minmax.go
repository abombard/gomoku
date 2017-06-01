package main

const MAXDEPTH = 5

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

func recminmax(board [][]int, pt coord, player int, depth int, alpha, beta int, ch chan step) step {

	next := getPossibleMoveList(board)
	// ERROR depth == MAXDEPTH && len(next) == 0 -> pt = shit
	if depth == 0 || len(next) == 0 {
		if ch != nil {
			ch <- step{pt, heuristic2(board)}
		}
		return step{pt, heuristic2(board)}
	}

	var v step
	if player == current {
		v = step{score: -10000}
	} else {
		v = step{score: 10000}
	}

	if depth == MAXDEPTH {
		newch := make(chan step, len(next))

		k := 0
		for i := range next {

			b := boardCopy(board)

			err := move(b, next[i], player, &b)
			if err != nil {
				if err.Error() == "Game Over" {
					depth = 1
				} else {
					k++
					continue
				}
			}

			go recminmax(b, next[i], (player+1)%2, depth-1, alpha, beta, newch)
		}

		for i := 0; i < len(next)-k; i++ {
			tmp := <-newch
			if player == current {
				if tmp.score > v.score {
					v.coord = tmp.coord
					v.score = tmp.score
				}
				if v.score > alpha {
					alpha = v.score
					if alpha >= beta {
						break
					}
				}
			} else {
				if tmp.score < v.score {
					v.coord = tmp.coord
					v.score = tmp.score
				}
				if v.score < beta {
					beta = v.score
					if alpha >= beta {
						break
					}
				}
			}
		}

	} else {

		for i := range next {

			var newBoard [][]int
			err := move(board, next[i], player, &newBoard)
			if err != nil {
				if err.Error() == "Game Over" {
					depth = 1
				} else {
					continue
				}
			}

			tmp := recminmax(newBoard, next[i], (player+1)%2, depth-1, alpha, beta, nil)
			if player == current {
				if tmp.score > v.score {
					v.coord = pt
					v.score = tmp.score
				}
				if v.score > alpha {
					alpha = v.score
					if alpha >= beta {
						board[next[i].X][next[i].Y] = 0
						break
					}
				}
			} else {
				if tmp.score < v.score {
					v.coord = pt
					v.score = tmp.score
				}
				if v.score < beta {
					beta = v.score
					if alpha >= beta {
						board[next[i].X][next[i].Y] = 0
						break
					}
				}
			}

			board[next[i].X][next[i].Y] = 0
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

func minmaxRoutine(board [][]int, pt coord, player int, ch chan step) {

	// create a new slice for each go routine
	//b := boardCopy(board)

	//ch <- recminmax(b, pt, player, MAXDEPTH, -10000, 10000)
}

func minmax(board [][]int, player int) coord {

	//nextMoves := getPossibleMoveList(board)

	b := boardCopy(board)

	v := recminmax(b, coord{0, 0, ""}, player, MAXDEPTH, -10000, 10000, nil)

	//log.Println("THE CHOOSEN ONE : ", v)

	return v.coord
}
