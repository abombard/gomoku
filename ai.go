package main

func aiPlay() coord {
	coord := minmax(g.Board, current)
	return coord
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
