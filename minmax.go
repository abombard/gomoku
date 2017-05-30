package main

import "log"

func getPossibleMoveList(b [][]int) []coord {
	var coords []coord

	for x := range b {
		for y := range b[0] {
			if isPawnNearby(x, y) == true && b[x][y] == 0 {
				coords = append(coords, coord{X: x, Y: y})
			}
		}
	}

	return coords
}

type fnptr func(x, y int, board [][]int, p int) bool

func isEmptyNew(x, y int, board [][]int, p int) bool {
	return board[x][y] == 0
}

func isMeNew(x, y int, board [][]int, player int) bool {
	return board[x][y] == player+1
}

func isEnemyNew(x, y int, board [][]int, player int) bool {
	return !isEmptyNew(x, y, board, player) && board[x][y] != player+1
}

func horizontalScoreNew(board [][]int, player int) int {
	score := 0
	finalScore := 0
	for y := 0; y < HEIGHT; y++ {
		for x := 0; isValidCoord(x, y); x++ {
			tmpx := x
			for ; isMeNew(tmpx, y, board, player); tmpx++ {
				score++
			}
			if score != 0 {
				tmpx = x
				space := 0
				spaceOk := false
				for ; isValidCoord(tmpx, y) && (isMeNew(tmpx, y, board, player) || !isEnemyNew(tmpx, y, board, player)); tmpx-- {
					space++
					if space+score >= 5 {
						spaceOk = true
						break
					}

				}
				if !spaceOk {
					tmpx = x + score
					for ; isValidCoord(tmpx, y) && (isMeNew(tmpx, y, board, player) || !isEnemyNew(tmpx, y, board, player)); tmpx++ {
						space++
						if space+score >= 5 {
							spaceOk = true
							break
						}

					}
				}
				if spaceOk {
					finalScore += score
					if score == 4 {
						if player == current {
							finalScore += 100
						} else {
							finalScore -= 200
						}
					}
				}
				score = 0
			}
		}
	}
	return finalScore
}

func verticalScoreNew(board [][]int, player int) int {
	score := 0
	finalScore := 0
	for x := 0; x < HEIGHT; x++ {
		for y := 0; isValidCoord(x, y); y++ {
			tmpy := y
			for ; isMeNew(x, tmpy, board, player); tmpy++ {
				score++
			}
			if score != 0 {
				tmpy = y
				space := 0
				spaceOk := false
				for ; isValidCoord(x, tmpy) && (isMeNew(x, tmpy, board, player) || !isEnemyNew(x, tmpy, board, player)); tmpy-- {
					space++
					if space+score >= 5 {
						spaceOk = true
						break
					}

				}
				if !spaceOk {
					tmpy = y + score
					for ; isValidCoord(x, tmpy) && (isMeNew(x, tmpy, board, player) || !isEnemyNew(x, tmpy, board, player)); tmpy++ {
						space++
						if space+score >= 5 {
							spaceOk = true
							break
						}

					}
				}
				if spaceOk {
					finalScore += score
					if score == 4 {
						if player == current {
							finalScore += 100
						} else {
							finalScore -= 200
						}
					}
				}
				score = 0
			}
		}
	}
	return finalScore
}

func heuristic(board [][]int, player int) int {

	score := 0
	score += horizontalScoreNew(board, player)
	score += horizontalScoreNew(board, player)
	score += verticalScoreNew(board, player)
	score += verticalScoreNew(board, player)

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

	if iter == 3 {
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
	if player == current {
		ret = min(scores)
	} else {
		ret = max(scores)
	}

	board[pt.X][pt.Y] = 0

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

	res := make([]step, 0)
	for i, _ := range nextMoves {
		pt := nextMoves[i]
		tmp := recminmax(b, pt, player, 0)
		res = append(res, tmp)
		log.Println(res)
	}

	ret := max(res)

	return ret.coord
}
