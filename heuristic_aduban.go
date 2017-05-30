package main

import "log"

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
				horSpace := 0
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx++ {
					if isMeNew(tmpx, tmpy, board, player) {
						horScore++
						horSpace++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						horSpace++
					}
				}
				tmpx = x
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx-- {
					if isMeNew(tmpx, tmpy, board, player) {
						horScore++
						horSpace++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						horSpace++
					}
				}
				tmpx = x
				verScore := 0
				verSpace := 0
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpy++ {
					if isMeNew(tmpx, tmpy, board, player) {
						verScore++
						verSpace++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						verSpace++
					}
				}
				tmpy = y
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpy-- {
					if isMeNew(tmpx, tmpy, board, player) {
						verScore++
						verSpace++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						verSpace++
					}
				}
				tmpy = y
				diagScore1 := 0
				diagSpace1 := 0
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx-1, tmpy-1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore1++
						diagSpace1++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						diagSpace1++
					}
				}
				tmpx = x
				tmpy = y
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx+1, tmpy+1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore1++
						diagSpace1++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						diagSpace1++
					}
				}
				tmpx = x
				tmpy = y
				diagScore2 := 0
				diagSpace2 := 0
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx+1, tmpy-1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore2++
						diagSpace2++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						diagSpace2++
					}
				}
				tmpx = x
				tmpy = y
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx-1, tmpy+1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore2++
						diagSpace2++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						diagSpace2++
					}
				}
				/*
					if verScore > horScore && verSpace >= 4 {
						b[x][y] = verScore
					} else if horScore > verScore && horSpace >= 4 {
						b[x][y] = horScore
					}*/
				log.Println(x, y, "horScore= ", horScore, "verScore= ", verScore, "dia1 score= ", diagScore1, "dia2score = ", diagScore2, "horspace= ", horSpace, "verspace= ", verSpace, "dia1space= ", diagSpace1, "diagspace2= ", diagSpace2)
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
				b[x][y] = maxval(horScore, verScore, diagScore1, diagScore2)
			}
		}
	}
	log.Println("_________")
	for x := 0; x < WIDTH; x++ {
		log.Println(board[x])
		for y := 0; y < HEIGHT; y++ {
			score += b[x][y]
			b[x][y] = 0
		}
	}
	log.Println("_________")
	log.Fatal("")
	//ENNEMY
	enemyScore := 0
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			if isMeNew(x, y, board, player) {
				tmpx := x
				tmpy := y
				horScore := 0
				horSpace := 0
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx++ {
					if isMeNew(tmpx, tmpy, board, player) {
						horScore++
						horSpace++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						horSpace++
					}
				}
				tmpx = x
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx-- {
					if isMeNew(tmpx, tmpy, board, player) {
						horScore++
						horSpace++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						horSpace++
					}
				}
				tmpx = x
				verScore := 0
				verSpace := 0
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpy++ {
					if isMeNew(tmpx, tmpy, board, player) {
						verScore++
						verSpace++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						verSpace++
					}
				}
				tmpy = y
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpy-- {
					if isMeNew(tmpx, tmpy, board, player) {
						verScore++
						verSpace++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						verSpace++
					}
				}
				tmpy = y
				diagScore1 := 0
				diagSpace1 := 0
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx-1, tmpy-1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore1++
						diagSpace1++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						diagSpace1++
					}
				}
				tmpx = x
				tmpy = y
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx+1, tmpy+1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore1++
						diagSpace1++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						diagSpace1++
					}
				}
				tmpx = x
				tmpy = y
				diagScore2 := 0
				diagSpace2 := 0
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx+1, tmpy-1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore2++
						diagSpace2++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						diagSpace2++
					}
				}
				tmpx = x
				tmpy = y
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx-1, tmpy+1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore2++
						diagSpace2++
					}
					if isEmptyNew(tmpx, tmpy, board, player) {
						diagSpace2++
					}
				}
				/*
					if verScore > horScore && verSpace >= 4 {
						b[x][y] = verScore
					} else if horScore > verScore && horSpace >= 4 {
						b[x][y] = horScore
					}*/
				if horSpace < 4 {
					log.Fatal("NO SPACE")
					horScore = 0
				}
				if verSpace < 4 {
					log.Fatal("NO SPACE")
					verScore = 0
				}
				if diagSpace1 < 4 {
					log.Fatal("NO SPACE")
					diagScore1 = 0
				}
				if diagSpace2 < 4 {
					log.Fatal("NO SPACE")
					diagScore2 = 0
				}
				b[x][y] = maxval(horScore, verScore, diagScore1, diagScore2)
			}
		}
	}
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			enemyScore += b[x][y]
		}
	}
	return score - enemyScore
}
