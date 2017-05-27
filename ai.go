package main

import (
	"log"
)

func aiPlay() {
	coords := getPossiblePlays()
	for _, coord := range coords {
		if g.Board[coord.X][coord.Y] == 0 {
			g.Board[coord.X][coord.Y] = 2
			return
		}

	}
	log.Println("COULDNT FIND PLAY")
}

func isPawnNearby(xtarg, ytarg int) bool {
	x := -2
	y := -2
	for ; x < 3; x++ {
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
			if isPawnNearby(x, y) == true {
				coords = append(coords, coord{X: x, Y: y})
			}
		}
	}

	return coords
}
