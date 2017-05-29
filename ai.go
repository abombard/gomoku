package main

import "log"

func aiPlay() {
	coords := getPossiblePlays()
	coord := maxi(coords)
	err := isValidMove(coord)
	if err != nil {
		println(err)
	}
	move(coord)
}

type fn func(x, y int) bool

func horizontalScore(me, him fn) int {
	score := 0
	finalScore := 0
	for y := 0; y < HEIGHT; y++ {
		for x := 0; isValidCoord(x, y); x++ {
			tmpx := x
			for ; me(tmpx, y); tmpx++ {
				score++
			}
			if score != 0 {
				tmpx = x
				space := 0
				spaceOk := false
				for ; isValidCoord(tmpx, y) && (me(tmpx, y) || !him(tmpx, y)); tmpx-- {
					space++
					if space >= 5 {
						spaceOk = true
						break
					}

				}
				if !spaceOk {
					tmpx = x + score
					for ; isValidCoord(tmpx, y) && (me(tmpx, y) || !him(tmpx, y)); tmpx++ {
						space++
						if space+score >= 5 {
							spaceOk = true
							break
						}

					}
				}
				if spaceOk {
					finalScore += score
				}
				score = 0
			}
		}
	}
	return finalScore

}

func verticalScore(me, him fn) int {
	score := 0
	finalScore := 0
	for x := 0; x < HEIGHT; x++ {
		for y := 0; isValidCoord(x, y); y++ {
			tmpy := y
			for ; me(x, tmpy); tmpy++ {
				score++
			}
			if score != 0 {
				tmpy = y
				space := 0
				spaceOk := false
				for ; isValidCoord(x, tmpy) && (me(x, tmpy) || !him(x, tmpy)); tmpy-- {
					space++
					if space >= 5 {
						spaceOk = true
						break
					}

				}
				if !spaceOk {
					tmpy = y + score
					for ; isValidCoord(x, tmpy) && (me(x, tmpy) || !him(x, tmpy)); tmpy++ {
						space++
						if space+score >= 5 {
							spaceOk = true
							break
						}

					}
				}
				if spaceOk {
					finalScore += score
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
	finalScore += horizontalScore(isMe, isEnemy)
	finalScore -= horizontalScore(isEnemy, isMe)
	finalScore += verticalScore(isMe, isEnemy)
	finalScore -= verticalScore(isEnemy, isMe)

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

func isPawnNearby(xtarg, ytarg int) bool {
	x := -1
	y := -1
	for ; x < 2; x++ {
		y = -1
		for ; y < 2; y++ {
			if xtarg+x >= 0 && xtarg+x < 19 && ytarg+y >= 0 && ytarg+y < 19 {
				if g.Board[xtarg+x][ytarg+y] != 0 {
					return true
				}
			}

		}
	}
	return false

}

func getPossiblePlays() []coord {
	var coords []coord

	for x := range g.Board {
		for y := range g.Board {
			if isPawnNearby(x, y) == true && g.Board[x][y] == 0 {
				coords = append(coords, coord{X: x, Y: y})
			}
		}
	}

	return coords
}
