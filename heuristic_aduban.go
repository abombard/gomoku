package main

func heuristic2(board [][]int, player int) int {
	b := make([][]int, HEIGHT)
	for i := 0; i < HEIGHT; i++ {
		b[i] = make([]int, WIDTH)
		for j := 0; j < WIDTH; j++ {
			b[i][j] = 0
		}
	}
	score := 0
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			if isMeNew(x, y, board, player) {
				tmpx := x
				tmpy := y
				horScore := 0
				for ; isValidCoord(tmpx, tmpy) && isMeNew(tmpx, tmpy, board, player); tmpx++ {
					horScore++
				}
				tmpx = x
				for ; isValidCoord(tmpx, tmpy) && isMeNew(tmpx, tmpy, board, player); tmpx-- {
					horScore++
				}
				tmpx = x
				verScore := 0
				for ; isValidCoord(tmpx, tmpy) && isMeNew(tmpx, tmpy, board, player); tmpy++ {
					verScore++
				}
				tmpy = y
				for ; isValidCoord(tmpx, tmpy) && isMeNew(tmpx, tmpy, board, player); tmpy-- {
					verScore++
				}
				if verScore > horScore {
					b[x][y] = verScore
				} else {
					b[x][y] = horScore
				}
			}
		}
	}
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			score += b[x][y]
			b[x][y] = 0
		}
	}
	enemyScore := 0
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			if isEnemyNew(x, y, board, player) {
				tmpx := x
				tmpy := y
				horScore := 0
				for ; isValidCoord(tmpx, tmpy) && isEnemyNew(tmpx, tmpy, board, player); tmpx++ {
					horScore++
				}
				tmpx = x
				for ; isValidCoord(tmpx, tmpy) && isEnemyNew(tmpx, tmpy, board, player); tmpx-- {
					horScore++
				}
				tmpx = x
				verScore := 0
				for ; isValidCoord(tmpx, tmpy) && isEnemyNew(tmpx, tmpy, board, player); tmpy++ {
					verScore++
				}
				tmpy = y
				for ; isValidCoord(tmpx, tmpy) && isEnemyNew(tmpx, tmpy, board, player); tmpy-- {
					verScore++
				}
				if verScore > horScore {
					b[x][y] = verScore
				} else {
					b[x][y] = horScore
				}
			}
		}
	}
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			enemyScore += b[x][y]
		}
	}
	return enemyScore - score
}
