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
	x := -2
	for ; x < 3; x++ {
		if c.X+x >= 0 && c.X+x < 19 {
			if g.Board[c.X+x][c.Y] == 1 {
				score += 1
			} else if g.Board[c.X+x][c.Y] == 2 {
				score += 1
			}
		}
	}
	return score
}

func checkVerticalScore(c coord) int {
	score := 0
	y := -2
	for ; y < 3; y++ {
		if c.Y+y >= 0 && c.Y+y < 19 {
			if g.Board[c.X][c.Y+y] == 1 {
				score += 1
			} else if g.Board[c.X][c.Y+y] == 2 {
				score += 1
			}
		}
	}
	return score
}

func checkDiagonalScore(c coord) int {
	score := 0
	y := -2
	x := -2
	for x < 3 {
		for y < 3 {
			if c.Y+y >= 0 && c.Y+y < 19 && c.X+x >= 0 && c.X+x < 19 {
				if g.Board[c.X+x][c.Y+y] == 1 {
					score += 1
				} else if g.Board[c.X+x][c.Y+y] == 2 {
					score += 1
				}
			}
			y++
		}
		x++
	}
	return score
}

func heuristic(coords []coord) coord {
	best := 0
	bestCoord := coords[0]
	for _, coord := range coords {
		currentScore := 0
		currentScore += checkHorizontalScore(coord)
		currentScore += checkVerticalScore(coord)
		currentScore += checkDiagonalScore(coord)
		if currentScore > best {
			best = currentScore
			bestCoord = coord
			log.Println("FOUND A NEW BEST MOVE")
		}
	}
	return bestCoord

}

func isPawnNearby(xtarg, ytarg int) bool {
	x := -2
	y := -2
	for ; x < 3; x++ {
		y = -2
		for ; y < 3; y++ {
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
