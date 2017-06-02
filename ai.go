package main

func aiPlay() coord {
	coord := minmax(g.Board, current)
	return coord
}

func isPawnNearby(b [][]int, x0, y0 int) bool {

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if isValidCoord(x0+x, y0+y) {
				if b[x0+x][y0+y] != 0 {
					return true
				}
			}

		}
	}

	return false
}

func check(c1, c2 int) bool {
	if !isEmpty(c1) && !isEmpty(c2) && c1 == c2 {
		return true
	}
	return false
}

func isPawnNearby2(b [][]int, x0, y0 int) bool {

	if isValidCoord(x0-1, y0-1) && isValidCoord(x0-2, y0-2) {
		if check(b[x0-1][y0-1], b[x0-2][y0-2]) {
			return true
		}
	}

	if isValidCoord(x0-1, y0) && isValidCoord(x0-2, y0) {
		if check(b[x0-1][y0], b[x0-2][y0]) {
			return true
		}
	}

	if isValidCoord(x0-1, y0+1) && isValidCoord(x0-2, y0+2) {
		if check(b[x0-1][y0+1], b[x0-2][y0+2]) {
			return true
		}
	}

	if isValidCoord(x0, y0-1) && isValidCoord(x0, y0-2) {
		if check(b[x0][y0-1], b[x0][y0-2]) {
			return true
		}
	}

	if isValidCoord(x0, y0+1) && isValidCoord(x0, y0+2) {
		if check(b[x0][y0+1], b[x0][y0+2]) {
			return true
		}
	}

	if isValidCoord(x0+1, y0-1) && isValidCoord(x0+2, y0-2) {
		if check(b[x0+1][y0-1], b[x0+2][y0-2]) {
			return true
		}
	}

	if isValidCoord(x0+1, y0) && isValidCoord(x0+2, y0) {
		if check(b[x0+1][y0], b[x0+2][y0]) {
			return true
		}
	}

	if isValidCoord(x0+1, y0+1) && isValidCoord(x0+2, y0+2) {
		if check(b[x0+1][y0+1], b[x0+2][y0+2]) {
			return true
		}
	}

	return false
}
