package main

import "log"

func aiPlay() {
	coords := getPossiblePlays()
	coord := heuristic(coords)
	if g.Board[coord.X][coord.Y] != 0 {
		log.Fatal("BAD PLAY")
	}
	g.Board[coord.X][coord.Y] = 2
}

func checkHorizontalScore(c coord) int {
	score := 0
	x := -4
	ennemyScore := 0
	best := 0
	for ; x < 0; x++ {
		for ; x < 0 && isValidCoord(c.X+x, c.Y) && isMe(c.X+x, c.Y); x++ {
			score = 0
			ennemyScore += 1
			log.Println("found me")
		}
		for ; x < 0 && isValidCoord(c.X+x, c.Y) && isEnemy(c.X+x, c.Y); x++ {
			log.Println("found ennemy")
			log.Println(c.X+x, c.Y)
			ennemyScore = 0
			score += 1
		}
	}
	if score > ennemyScore {
		best = score
	} else {
		best = ennemyScore
	}
	x = 4
	score = 0
	ennemyScore = 0
	for ; x > 0; x-- {
		for ; x > 0 && isValidCoord(c.X+x, c.Y) && isMe(c.X+x, c.Y); x-- {
			log.Println("found me")
			score = 0
			ennemyScore += 1
		}
		for ; x > 0 && isValidCoord(c.X+x, c.Y) && isEnemy(c.X+x, c.Y); x-- {
			ennemyScore = 0
			score += 1
		}
	}
	if ennemyScore > best {
		return ennemyScore
	} else if score > best {
		return score
	}
	return best
}

func checkVerticalScore(c coord) int {
	score := 0
	y := -4
	ennemyScore := 0
	best := 0
	for ; y < 0; y++ {
		for ; y < 0 && isValidCoord(c.X, c.Y+y) && isMe(c.X, c.Y+y); y++ {
			score = 0
			ennemyScore += 1
		}
		for ; y < 0 && isValidCoord(c.X, c.Y+y) && isEnemy(c.X, c.Y+y); y++ {
			ennemyScore = 0
			score += 1
		}
	}
	if score > ennemyScore {
		best = score
	} else {
		best = ennemyScore
	}
	score = 0
	ennemyScore = 0
	y = 4
	for ; y > 0; y-- {
		for ; y > 0 && isValidCoord(c.X, c.Y+y) && isMe(c.X, c.Y+y); y-- {
			score = 0
			ennemyScore += 1
		}
		for ; y > 0 && isValidCoord(c.X, c.Y+y) && isEnemy(c.X, c.Y+y); y-- {
			ennemyScore = 0
			score += 1
		}
	}
	if ennemyScore > best {
		return ennemyScore
	} else if score > best {
		return score
	}
	return best
}

func checkDiagonalScore(c coord) int {
	score := 0
	y := -4
	x := -4
	ennemyScore := 0
	best := 0
	for x < 0 {
		for y < 0 {
			for ; x < 0 && y < 0 && isValidCoord(c.X+x, c.Y+y) && isMe(c.X+x, c.Y+y); x, y = x+1, y+1 {
				score = 0
				ennemyScore += 1
			}
			for ; x < 0 && y < 0 && isValidCoord(c.X+x, c.Y+y) && isEnemy(c.X+x, c.Y+y); x, y = x+1, y+1 {
				ennemyScore = 0
				score += 1
			}
			y++
		}
		x++
	}
	if score > ennemyScore {
		best = score
	} else {
		best = ennemyScore
	}
	score = 0
	ennemyScore = 0
	y = 4
	x = 4
	for x > 0 {
		for y > 0 {
			for ; x > 0 && y > 0 && isValidCoord(c.X+x, c.Y+y) && isMe(c.X+x, c.Y+y); x, y = x-1, y-1 {
				score = 0
				ennemyScore += 1
			}
			for ; x > 0 && y > 0 && isValidCoord(c.X+x, c.Y+y) && isEnemy(c.X+x, c.Y+y); x, y = x-1, y-1 {
				ennemyScore = 0
				score += 1
			}
		}
	}
	if ennemyScore > best {
		best = ennemyScore
	} else if score > best {
		best = score
	}
	score = 0
	ennemyScore = 0
	y = 4
	x = -4
	ennemyScore = 0
	score = 0
	for x < 0 {
		for y > 0 {
			for ; x < 0 && y > 0 && isValidCoord(c.X+x, c.Y+y) && isMe(c.X+x, c.Y+y); x, y = x+1, y-1 {
				score = 0
				ennemyScore += 1
			}
			for ; x < 0 && y > 0 && isValidCoord(c.X+x, c.Y+y) && isEnemy(c.X+x, c.Y+y); x, y = x+1, y-1 {
				ennemyScore = 0
				score += 1
			}
			y--
		}
		x++
	}
	if score > ennemyScore {
		best = score
	} else {
		best = ennemyScore
	}
	score = 0
	ennemyScore = 0
	y = -4
	x = 4
	for x > 0 {
		for y < 0 {
			for ; x > 0 && y < 0 && isValidCoord(c.X+x, c.Y+y) && isMe(c.X+x, c.Y+y); x, y = x-1, y+1 {
				score = 0
				ennemyScore += 1
			}
			for ; x > 0 && y < 0 && isValidCoord(c.X+x, c.Y+y) && isEnemy(c.X+x, c.Y+y); x, y = x-1, y+1 {
				ennemyScore = 0
				score += 1
			}
			y++
		}
		x--
	}
	if ennemyScore > best {
		best = ennemyScore
	} else if score > best {
		best = score
	}
	return best
}

func heuristic(coords []coord) coord {
	best := 0
	bestCoord := coords[0]
	for _, coord := range coords {
		currentScore := 0
		log.Println(coord)
		currentScore += checkHorizontalScore(coord)
		log.Println("Horizontal score: ", currentScore)
		tmp := checkVerticalScore(coord)
		if tmp > currentScore {
			currentScore = tmp
		}
		log.Println("Vertical score: ", tmp)
		tmp = checkDiagonalScore(coord)
		if tmp > currentScore {
			currentScore = tmp
		}
		log.Println("Diag score: ", tmp)
		if currentScore > best {
			best = currentScore
			bestCoord = coord
		}
		log.Println("_____________")
	}
	return bestCoord

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
