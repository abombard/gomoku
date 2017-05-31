package main

func maxval(a, b, c, d int) int {
	best := 0
	if a > best {
		best = a
	}
	if b > best {
		best = b
	}
	if c > best {
		best = c
	}
	if d > best {
		best = d
	}
	return best
}

func getScore(board [][]int, player int) int {
	score := 0
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			if isMeNew(x, y, board, player) {
				tmpx := x + 1
				tmpy := y
				horScore := 0
				horSpace := 0
				verScore := 0
				verSpace := 0
				diagScore1 := 0
				diagSpace1 := 0
				diagScore2 := 0
				diagSpace2 := 0
				b := 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx++ {
					if isMeNew(tmpx, tmpy, board, player) {
						horScore += b
						b++
					}
					horSpace++
				}
				tmpx = x - 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx-- {
					if isMeNew(tmpx, tmpy, board, player) {
						horScore += b
						b++
					}
					horSpace++
				}
				tmpx = x
				tmpy = y + 1
				b = 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpy++ {
					if isMeNew(tmpx, tmpy, board, player) {
						verScore += b
						b++
					}
					verSpace++
				}
				tmpy = y - 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpy-- {
					if isMeNew(tmpx, tmpy, board, player) {
						verScore += b
						b++
					}
					verSpace++
				}
				b = 1
				tmpy = y - 1
				tmpx = x - 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx-1, tmpy-1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore1 += b
						b++
					}
					diagSpace1++
				}
				tmpx = x + 1
				tmpy = y + 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx+1, tmpy+1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore1 += b
						b++
					}
					diagSpace1++
				}
				b = 1
				tmpx = x + 1
				tmpy = y - 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx+1, tmpy-1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore2 += b
						b++
					}
					diagSpace2++
				}
				tmpx = x - 1
				tmpy = y + 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx-1, tmpy+1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore2 += b
						b++
					}
					diagSpace2++
				}
				//log.Println(x, y, "horScore= ", horScore, "verScore= ", verScore, "dia1 score= ", diagScore1, "dia2score = ", diagScore2, "horspace= ", horSpace, "verspace= ", verSpace, "dia1space= ", diagSpace1, "diagspace2= ", diagSpace2)

				if horSpace < 4 {
					horScore = 0
				}
				if verSpace < 4 {
					verScore = 0
				}
				if diagSpace1 < 4 {
					diagScore1 = 0
				}
				if diagSpace2 < 4 {
					diagScore2 = 0
				}
				score += horScore + verScore + diagScore1 + diagScore2
			}
		}
	}
	return score
}

func heuristic2(board [][]int, player int) int {
	score := getScore(board, player)
	if player == 1 {
		player = 0
	} else if player == 0 {
		player = 1
	}
	enemyScore := getScore(board, player)
	return score - enemyScore
}
