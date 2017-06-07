package main

import "fmt"

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

func getScore1(board [][]int, player int) int {
	score := 0
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			if isMe(board[x][y], player) {
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
				stop := false
				for ; isValidCoord(tmpx, tmpy) && (isMe(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpx++ {
					if !stop && isMe(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						horScore += b
						b++
					} else {
						stop = true
					}

					horSpace++
					if horSpace > 5 {
						break
					}
				}
				stop = false
				tmpx = x - 1
				for ; isValidCoord(tmpx, tmpy) && (isMe(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpx-- {
					if !stop && isMe(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						horScore += b
						b++
					} else {
						stop = true
					}

					horSpace++
					if horSpace > 5 {
						break
					}
				}
				stop = false
				tmpx = x
				tmpy = y + 1
				b = 1
				for ; isValidCoord(tmpx, tmpy) && (isMe(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpy++ {
					if !stop && isMe(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						verScore += b
						b++
					} else {
						stop = true
					}
					verSpace++
					if verSpace > 5 {
						break
					}
				}
				stop = false
				tmpy = y - 1
				for ; isValidCoord(tmpx, tmpy) && (isMe(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpy-- {
					if !stop && isMe(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						verScore += b
						b++
					} else {
						stop = true
					}

					verSpace++
					if verSpace > 5 {
						break
					}
				}
				stop = false
				b = 1
				tmpy = y - 1
				tmpx = x - 1
				for ; isValidCoord(tmpx, tmpy) && (isMe(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpx, tmpy = tmpx-1, tmpy-1 {
					if !stop && isMe(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						diagScore1 += b
						b++
					} else {
						stop = true
					}

					diagSpace1++
					if diagSpace1 > 5 {
						break
					}
				}
				stop = false
				tmpx = x + 1
				tmpy = y + 1
				for ; isValidCoord(tmpx, tmpy) && (isMe(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpx, tmpy = tmpx+1, tmpy+1 {
					if !stop && isMe(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						diagScore1 += b
						b++
					} else {
						stop = true
					}

					diagSpace1++
					if diagSpace1 > 5 {
						break
					}
				}
				stop = false
				b = 1
				tmpx = x + 1
				tmpy = y - 1
				for ; isValidCoord(tmpx, tmpy) && (isMe(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpx, tmpy = tmpx+1, tmpy-1 {
					if !stop && isMe(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						diagScore2 += b
						b++
					} else {
						stop = true
					}

					diagSpace2++
					if diagSpace2 > 5 {
						break
					}
				}
				stop = false
				tmpx = x - 1
				tmpy = y + 1
				for ; isValidCoord(tmpx, tmpy) && (isMe(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpx, tmpy = tmpx-1, tmpy+1 {
					if !stop && isMe(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						diagScore2 += b
						b++
					} else {
						stop = true
					}

					diagSpace2++
					if diagSpace2 > 5 {
						break
					}
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
			} else if isEnemy(board[x][y], player) {
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
				stop := false
				for ; isValidCoord(tmpx, tmpy) && (isEnemy(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpx++ {
					if !stop && isEnemy(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						horScore += b
						b++
					} else {
						stop = true
					}

					horSpace++
					if horSpace > 5 {
						break
					}
				}
				stop = false
				tmpx = x - 1
				for ; isValidCoord(tmpx, tmpy) && (isEnemy(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpx-- {
					if !stop && isEnemy(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						horScore += b
						b++
					} else {
						stop = true
					}

					horSpace++
					if horSpace > 5 {
						break
					}
				}
				stop = false
				tmpx = x
				tmpy = y + 1
				b = 1
				for ; isValidCoord(tmpx, tmpy) && (isEnemy(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpy++ {
					if !stop && isEnemy(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						verScore += b
						b++
					} else {
						stop = true
					}
					verSpace++
					if verSpace > 5 {
						break
					}
				}
				stop = false
				tmpy = y - 1
				for ; isValidCoord(tmpx, tmpy) && (isEnemy(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpy-- {
					if !stop && isEnemy(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						verScore += b
						b++
					} else {
						stop = true
					}

					verSpace++
					if verSpace > 5 {
						break
					}
				}
				stop = false
				b = 1
				tmpy = y - 1
				tmpx = x - 1
				for ; isValidCoord(tmpx, tmpy) && (isEnemy(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpx, tmpy = tmpx-1, tmpy-1 {
					if !stop && isEnemy(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						diagScore1 += b
						b++
					} else {
						stop = true
					}

					diagSpace1++
					if diagSpace1 > 5 {
						break
					}
				}
				stop = false
				tmpx = x + 1
				tmpy = y + 1
				for ; isValidCoord(tmpx, tmpy) && (isEnemy(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpx, tmpy = tmpx+1, tmpy+1 {
					if !stop && isEnemy(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						diagScore1 += b
						b++
					} else {
						stop = true
					}

					diagSpace1++
					if diagSpace1 > 5 {
						break
					}
				}
				stop = false
				b = 1
				tmpx = x + 1
				tmpy = y - 1
				for ; isValidCoord(tmpx, tmpy) && (isEnemy(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpx, tmpy = tmpx+1, tmpy-1 {
					if !stop && isEnemy(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						diagScore2 += b
						b++
					} else {
						stop = true
					}

					diagSpace2++
					if diagSpace2 > 5 {
						break
					}
				}
				stop = false
				tmpx = x - 1
				tmpy = y + 1
				for ; isValidCoord(tmpx, tmpy) && (isEnemy(board[tmpx][tmpy], player) || isEmpty(board[tmpx][tmpy])); tmpx, tmpy = tmpx-1, tmpy+1 {
					if !stop && isEnemy(board[tmpx][tmpy], player) {
						if canBeCaptured(board, tmpx, tmpy) {
							b = 1
						}
						diagScore2 += b
						b++
					} else {
						stop = true
					}

					diagSpace2++
					if diagSpace2 > 5 {
						break
					}
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

				score -= horScore + verScore + diagScore1 + diagScore2
			}
		}
	}
	return score
}

func getScore(board [][]int, player int) int {

	score := 0

	curP := -1
	curScore := 0
	curSpacePrev := 0
	curSpaceNext := 0

	reset := func() {
		curP = -1
		curScore = 0
		curSpacePrev = 0
		curSpaceNext = 0
	}

	calcScore := func(curScore int) int {
		s := curScore
		if s+curSpacePrev+curSpaceNext >= 5 {
			s *= s
			if curScore >= 5 {
				s *= 5
			}
			if isEnemy(curP, player) {
				s = -s
			}
			return s
		}
		return 0
	}

	updateScore := func(x, y int) {
		score += calcScore(curScore)
		if curSpacePrev+curScore+curSpaceNext >= 5 {
			if curSpacePrev > 0 {
				curSpacePrev -= 1
			} else if curScore > 0 {
				curScore -= 1
			} else if curSpaceNext > 0 {
				curSpaceNext -= 1
			}
		}
		if isEmpty(board[x][y]) {
			curSpaceNext += 1
		} else {
			if board[x][y] == curP {
				curSpacePrev += curSpaceNext
				curSpaceNext = 0
			} else {
				curScore = 0
				curP = board[x][y]
				curSpacePrev = curSpaceNext
				curSpaceNext = 0
			}
			if canBeCaptured(board, x, y) {
				curScore = 0
				curSpacePrev = 0
				curSpaceNext = 0
			} else {
				curScore += 1
			}
		}
	}

	// horizontal
	for x := 0; x < HEIGHT; x++ {
		reset()
		for y := 0; y < WIDTH; y++ {
			updateScore(x, y)
		}
	}

	// vertical
	for y := 0; y < WIDTH; y++ {
		reset()
		for x := 0; x < HEIGHT; x++ {
			updateScore(x, y)
		}
	}

	// diagonal 1 from left
	for x0, y0 := 0, 0; x0 < HEIGHT-4; x0++ {
		reset()
		for x, y := x0, y0; x < HEIGHT; x, y = x+1, y+1 {
			updateScore(x, y)
		}
	}

	// diagonal 1 from top
	for x0, y0 := 0, 1; y0 < WIDTH-4; y0++ {
		reset()
		for x, y := x0, y0; y < WIDTH; x, y = x+1, y+1 {
			updateScore(x, y)
		}
	}

	// diagonal 2 from left
	for x0, y0 := 4, 0; x0 < HEIGHT; x0++ {
		reset()
		for x, y := x0, y0; x >= 0; x, y = x-1, y+1 {
			updateScore(x, y)
		}
	}

	// diagonal 2 from bot
	for x0, y0 := HEIGHT-1, 1; y0 < WIDTH-4; y0++ {
		reset()
		for x, y := x0, y0; y < WIDTH; x, y = x-1, y+1 {
			updateScore(x, y)
		}
	}

	return score
}

var trans = make(map[string]int)

func heuristic2(board [][]int, player int) int {
	if val, ok := trans[fmt.Sprintln(board)]; ok {
		return val
	}
	score := getScore(board, player)
	trans[fmt.Sprintln(board)] = score
	return score
}
