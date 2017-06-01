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

func recminmax(board [][]int, player int, depth int, alpha, beta int, ch chan step) step {

	next := getPossibleMoveList(board)
	if depth == 0 || len(next) == 0 {
		return step{score: heuristic2(board)}
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

			var newBoard [][]int
			err := move(board, next[i], player, &newBoard)
			if err != nil {
				if err.Error() == "Game Over" {
					depth = 1
				} else {
					continue
				}
			}
			k++

			if len(newBoard) == 0 {
				newBoard = board
			}

			go recminmax(newBoard, (player+1)%2, depth-1, alpha, beta, newch)
			tmp := <-newch
			if player == current {
				if tmp.score > v.score {
					v.coord = next[i]
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
					v.coord = next[i]
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
		ch <- v

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

			if len(newBoard) == 0 {
				newBoard = board
			}

			tmp := recminmax(newBoard, (player+1)%2, depth-1, alpha, beta, nil)
			if player == current {
				if tmp.score > v.score {
					v.coord = next[i]
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
					v.coord = next[i]
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

	v := recminmax(board, player, MAXDEPTH, -10000, 10000, nil)

	//log.Println("THE CHOOSEN ONE : ", v)

	return v.coord
}
