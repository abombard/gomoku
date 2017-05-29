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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func checkHorizontalScore(c coord) int {
	score := 0
	ennemyScore := 0
	bestRightScore := 0
	bestLeftScore := 0
	bestRightEnnemyScore := 0
	bestLeftEnnemyScore := 0
	for x := -4; x < 0 && isValidCoord(c.X+x, c.Y); x++ {
		if isEnemy(c.X+x, c.Y) {
			ennemyScore += 2
			score = 0
		} else if isMe(c.X+x, c.Y) {
			score += 2
			ennemyScore = 0
		}
	}
	bestLeftScore = score
	bestLeftEnnemyScore = ennemyScore
	score = 0
	ennemyScore = 0
	for x := 4; x > 0 && isValidCoord(c.X+x, c.Y); x-- {
		if isEnemy(c.X+x, c.Y) {
			ennemyScore += 2
			score = 0
		} else if isMe(c.X+x, c.Y) {
			score += 2
			ennemyScore = 0
		}
	}
	bestRightScore = score
	bestRightEnnemyScore = ennemyScore
	if bestLeftScore > bestLeftEnnemyScore && bestRightScore > bestRightEnnemyScore {
		return bestLeftScore + bestRightScore
	}
	if bestLeftScore < bestLeftEnnemyScore && bestRightScore < bestRightEnnemyScore {
		return bestLeftEnnemyScore + bestRightEnnemyScore
	}
	maxScore := max(bestLeftScore, bestRightScore)
	maxEnnemyScore := max(bestLeftEnnemyScore, bestRightEnnemyScore)
	return max(maxScore, maxEnnemyScore)
}

func checkVerticalScore(c coord) int {
	score := 0
	ennemyScore := 0
	bestRightScore := 0
	bestLeftScore := 0
	bestRightEnnemyScore := 0
	bestLeftEnnemyScore := 0
	for y := -4; y < 0 && isValidCoord(c.X, c.Y+y); y++ {
		if isEnemy(c.X, c.Y+y) {
			ennemyScore += 2
			score = 0
		} else if isMe(c.X, c.Y+y) {
			score += 2
			ennemyScore = 0
		}
	}
	bestLeftScore = score
	bestLeftEnnemyScore = ennemyScore
	score = 0
	ennemyScore = 0
	for y := 4; y > 0 && isValidCoord(c.X, c.Y+y); y-- {
		if isEnemy(c.X, c.Y+y) {
			ennemyScore += 2
			score = 0
		} else if isMe(c.X, c.Y+y) {
			score += 2
			ennemyScore = 0
		}
	}
	bestRightScore = score
	bestRightEnnemyScore = ennemyScore
	if bestLeftScore > bestLeftEnnemyScore && bestRightScore > bestRightEnnemyScore {
		return bestLeftScore + bestRightScore
	}
	if bestLeftScore < bestLeftEnnemyScore && bestRightScore < bestRightEnnemyScore {
		return bestLeftEnnemyScore + bestRightEnnemyScore
	}
	maxScore := max(bestLeftScore, bestRightScore)
	maxEnnemyScore := max(bestLeftEnnemyScore, bestRightEnnemyScore)
	return max(maxScore, maxEnnemyScore)
}

func checkDiagonalScore(c coord) int {
	score := 0
	ennemyScore := 0
	bestRightScore := 0
	bestLeftScore := 0
	bestRightEnnemyScore := 0
	bestLeftEnnemyScore := 0
	for x, y := -4, -4; y < 0 && x < 0 && isValidCoord(c.X+x, c.Y+y); x, y = x+1, y+1 {
		if isEnemy(c.X+x, c.Y+y) {
			ennemyScore += 2
			score = 0
		} else if isMe(c.X+x, c.Y+y) {
			score += 2
			ennemyScore = 0
		}
	}
	bestLeftScore = score
	bestLeftEnnemyScore = ennemyScore
	score = 0
	ennemyScore = 0
	for x, y := 4, 4; y > 0 && x > 0 && isValidCoord(c.X+x, c.Y+y); x, y = x-1, y-1 {
		if isEnemy(c.X+x, c.Y+y) {
			ennemyScore += 2
			score = 0
		} else if isMe(c.X+x, c.Y+y) {
			score += 2
			ennemyScore = 0
		}
	}
	bestRightScore = score
	bestRightEnnemyScore = ennemyScore
	if bestLeftScore > bestLeftEnnemyScore && bestRightScore > bestRightEnnemyScore {
		return bestLeftScore + bestRightScore
	}
	if bestLeftScore < bestLeftEnnemyScore && bestRightScore < bestRightEnnemyScore {
		return bestLeftEnnemyScore + bestRightEnnemyScore
	}
	maxScore := max(bestLeftScore, bestRightScore)
	maxEnnemyScore := max(bestLeftEnnemyScore, bestRightEnnemyScore)
	return max(maxScore, maxEnnemyScore)
}
func checkDiagonalScore2(c coord) int {
	score := 0
	ennemyScore := 0
	bestRightScore := 0
	bestLeftScore := 0
	bestRightEnnemyScore := 0
	bestLeftEnnemyScore := 0
	for x, y := -4, 4; y > 0 && x < 0 && isValidCoord(c.X+x, c.Y+y); x, y = x+1, y-1 {
		if isEnemy(c.X+x, c.Y+y) {
			ennemyScore += 2
			score = 0
		} else if isMe(c.X+x, c.Y+y) {
			score += 2
			ennemyScore = 0
		}
	}
	bestLeftScore = score
	bestLeftEnnemyScore = ennemyScore
	score = 0
	ennemyScore = 0
	for x, y := 4, -4; y < 0 && x > 0 && isValidCoord(c.X+x, c.Y+y); x, y = x-1, y+1 {
		if isEnemy(c.X+x, c.Y+y) {
			ennemyScore += 2
			score = 0
		} else if isMe(c.X+x, c.Y+y) {
			score += 2
			ennemyScore = 0
		}
	}
	bestRightScore = score
	bestRightEnnemyScore = ennemyScore
	if bestLeftScore > bestLeftEnnemyScore && bestRightScore > bestRightEnnemyScore {
		return bestLeftScore + bestRightScore
	}
	if bestLeftScore < bestLeftEnnemyScore && bestRightScore < bestRightEnnemyScore {
		return bestLeftEnnemyScore + bestRightEnnemyScore
	}
	maxScore := max(bestLeftScore, bestRightScore)
	maxEnnemyScore := max(bestLeftEnnemyScore, bestRightEnnemyScore)
	return max(maxScore, maxEnnemyScore)
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
