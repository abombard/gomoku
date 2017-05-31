package main

import "log"

func aiPlay() coord {
	coord := minmax(g.Board, current)
	return coord
}

type fn func(x, y int, tmp [19][19]int) bool

func isEnemyTmp(x, y int, tmp [19][19]int) bool {
	return !isEmptyTmp(x, y, tmp) && tmp[x][y] != current+1
}

func isMeTmp(x, y int, tmp [19][19]int) bool {
	return !isEmptyTmp(x, y, tmp) && tmp[x][y] == current+1
}
func isEmptyTmp(x, y int, tmp [19][19]int) bool {
	return tmp[x][y] == 0
}

func horizontalScore(who string, me, him fn, tmp [19][19]int) int {
	score := 0
	finalScore := 0
	for y := 0; y < HEIGHT; y++ {
		for x := 0; isValidCoord(x, y); x++ {
			tmpx := x
			for ; me(tmpx, y, tmp); tmpx++ {
				score++
			}
			if score != 0 {
				tmpx = x
				space := 0
				spaceOk := false
				for ; isValidCoord(tmpx, y) && (me(tmpx, y, tmp) || !him(tmpx, y, tmp)); tmpx-- {
					space++
					if space+score >= 5 {
						spaceOk = true
						break
					}

				}
				if !spaceOk {
					tmpx = x + score
					for ; isValidCoord(tmpx, y) && (me(tmpx, y, tmp) || !him(tmpx, y, tmp)); tmpx++ {
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
						if who == "me" {
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

func verticalScore(who string, me, him fn, tmp [19][19]int) int {
	score := 0
	finalScore := 0
	for x := 0; x < HEIGHT; x++ {
		for y := 0; isValidCoord(x, y); y++ {
			tmpy := y
			for ; me(x, tmpy, tmp); tmpy++ {
				score++
			}
			if score != 0 {
				tmpy = y
				space := 0
				spaceOk := false
				for ; isValidCoord(x, tmpy) && (me(x, tmpy, tmp) || !him(x, tmpy, tmp)); tmpy-- {
					space++
					if space+score >= 5 {
						spaceOk = true
						break
					}

				}
				if !spaceOk {
					tmpy = y + score
					for ; isValidCoord(x, tmpy) && (me(x, tmpy, tmp) || !him(x, tmpy, tmp)); tmpy++ {
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
						if who == "me" {
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

type resp struct {
	C     coord
	Score int
}

func evaluate(c coord, ch chan resp) {
	var tmp [19][19]int
	for i, row := range g.Board {
		tmp[i] = row
	}
	tmp[c.X][c.Y] = 2
	finalScore := 0
	finalScore += horizontalScore("me", isMeTmp, isEnemyTmp, tmp)
	finalScore += horizontalScore("ennemy", isEnemyTmp, isMeTmp, tmp)
	finalScore += verticalScore("me", isMeTmp, isEnemyTmp, tmp)
	finalScore += verticalScore("ennemy", isEnemyTmp, isMeTmp, tmp)

	ch <- resp{C: c, Score: finalScore}
}

func maxi(coords []coord) coord {
	var tmp [19][19]int
	ch := make(chan resp, len(coords))
	for i, row := range g.Board {
		tmp[i] = row
	}
	for i, _ := range coords {
		val := coords[i]
		go evaluate(val, ch)
	}
	ret := resp{Score: -1000}
	for i := 0; i < len(coords); i++ {
		tmp := <-ch
		log.Println(tmp)
		if tmp.Score > ret.Score {
			ret = tmp
		}
	}

	return coord{X: ret.C.X, Y: ret.C.Y}
}

func isPawnNearby(b [][]int, xtarg, ytarg int) bool {
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if isValidCoord(xtarg+x, ytarg+y) {
				if b[xtarg+x][ytarg+y] != 0 {
					return true
				}
			}

		}
	}
	return false
}
