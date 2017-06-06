package main

import (
	"fmt"
)

func getPossibleMoveListNew(b [][]int, pt, old coord) []coord {

	var coords []coord
	for x := -4; x < 5; x++ {
		for y := -4; y < 5; y++ {
			if isValidCoord(pt.X+x, pt.Y+y) && b[pt.X+x][pt.Y+y] == 0 {
				coords = append(coords, coord{X: pt.X + x, Y: pt.Y + y})
			}

		}
	}
	for x := -4; x < 5; x++ {
		for y := -4; y < 5; y++ {
			if isValidCoord(old.X+x, old.Y+y) && b[old.X+x][old.Y+y] == 0 {
				coords = append(coords, coord{X: old.X + x, Y: old.Y + y})
			}

		}
	}
	return coords
}
func isValidCoord(x, y int) bool {
	return x >= 0 && x < HEIGHT && y >= 0 && y < WIDTH
}

func isEmpty(c int) bool {
	return c == 0
}

func isEnemy(c, p int) bool {
	return !isEmpty(c) && c != p+1
}

func isMe(c, p int) bool {
	return !isEmpty(c) && c == p+1
}

func introduceDoubleThree(b [][]int, x, y int, p int) bool {

	checks := [4][8][2]int{
		{{-4, -4}, {-3, -3}, {-2, -2}, {-1, -1}, {1, 1}, {2, 2}, {3, 3}, {4, 4}},
		{{-4, 0}, {-3, 0}, {-2, 0}, {-1, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}},
		{{-4, 4}, {-3, 3}, {-2, 2}, {-1, 1}, {1, -1}, {2, -2}, {3, -3}, {4, -4}},
		{{0, 4}, {0, 3}, {0, 2}, {0, 1}, {0, -1}, {0, -2}, {0, -3}, {0, -4}},
	}

	threeCount := 0
	for i := range checks {
		prevOk := false
		nextOk := false
		prev := -1
		count := 0
		blank := 0
		for j := 0; j < len(checks[i]); j++ {
			if j == 4 {
				count += 1
			}
			x1, y1 := x+checks[i][j][0], y+checks[i][j][1]
			if isValidCoord(x1, y1) {
				if count == 3 {
					nextOk = isEmpty(b[x][y])
					break
				}
				if isMe(b[x1][y1], p) {
					if count == 0 {
						blank = 0
					}
					count += 1
				} else if isEnemy(b[x1][y1], p) {
					if count == 0 {
						prevOk = false
					}
					blank = 0
					count = 0
				} else {
					blank += 1
					if blank > 1 && count > 0 {
						if isEmpty(prev) {
							count = 0
						} else {
							count -= 1
							blank -= 1
						}
					}
					if count == 0 {
						prevOk = true
					}
				}
				prev = b[x1][y1]
			}
		}
		if count == 3 && prevOk && nextOk {
			threeCount += 1
		}
	}

	return threeCount > 1
}

func capture(b [][]int, x, y int, p int, newBoard *[][]int) bool {

	ret := false

	checks := [8][3][2]int{
		{{-1, -1}, {-2, -2}, {-3, -3}},
		{{-1, 0}, {-2, 0}, {-3, 0}},
		{{-1, 1}, {-2, 2}, {-3, 3}},
		{{0, 1}, {0, 2}, {0, 3}},
		{{1, 1}, {2, 2}, {3, 3}},
		{{1, 0}, {2, 0}, {3, 0}},
		{{1, -1}, {2, -2}, {3, -3}},
		{{0, -1}, {0, -2}, {0, -3}},
	}
	for i := range checks {
		p1, p2, p3 := checks[i][0], checks[i][1], checks[i][2]
		x1, y1 := x+p1[0], y+p1[1]
		x2, y2 := x+p2[0], y+p2[1]
		x3, y3 := x+p3[0], y+p3[1]
		if isValidCoord(x1, y1) && isValidCoord(x2, y2) && isValidCoord(x3, y3) {
			if isEnemy(b[x1][y1], p) && isEnemy(b[x2][y2], p) && isMe(b[x3][y3], p) {
				// Successfully captured

				if len(*newBoard) == 0 {
					*newBoard = boardCopy(b)
				}
				(*newBoard)[x1][y1] = 0
				(*newBoard)[x2][y2] = 0
				ret = true
			}
		}
	}

	return ret
}

func canBeCaptured(b [][]int, x, y int, p int) bool {

	checks := [][4][2]int{
		{{-2, -2}, {-1, -1}, {1, 1}, {2, 2}},
		{{-2, 0}, {-1, 0}, {1, 0}, {2, 0}},
		{{-2, 2}, {-1, 1}, {1, -1}, {2, -2}},
		{{0, 2}, {0, 1}, {0, -1}, {0, -2}},
	}
	for i := range checks {
		p1, p2, p3, p4 := checks[i][0], checks[i][1], checks[i][2], checks[i][3]
		x1, y1 := x+p1[0], y+p1[1]
		x2, y2 := x+p2[0], y+p2[1]
		x3, y3 := x+p3[0], y+p3[1]
		x4, y4 := x+p4[0], y+p4[1]
		if isValidCoord(x1, y1) && isValidCoord(x2, y2) && isValidCoord(x3, y3) {
			if isMe(b[x2][y2], p) && ((isEmpty(b[x1][y1]) && isEnemy(b[x3][y3], p)) || (isEnemy(b[x1][y1], p) && isEmpty(b[x3][y3]))) {
				return true
			}
		}
		if isValidCoord(x2, y2) && isValidCoord(x3, y3) && isValidCoord(x4, y4) {
			if isMe(b[x3][y3], p) && ((isEmpty(b[x2][y2]) && isEnemy(b[x4][y4], p)) || (isEnemy(b[x2][y2], p) && isEmpty(b[x4][y4]))) {
				return true
			}
		}
	}
	return false
}

func isGameOver(b [][]int, c coord, p int) bool {

	for x0 := 0; x0 < WIDTH; x0++ {
		for y0 := 0; y0 < WIDTH; y0++ {

			count := 0
			for x, y := x0-4, y0; x <= x0+4; x++ {
				if isValidCoord(x, y) && isMe(b[x][y], p) && !canBeCaptured(b, x, y, p) {
					count += 1
					if count == 5 {
						return true
					}
				} else {
					count = 0
				}
			}

			count = 0
			for x, y := x0, y0-4; y <= y0+4; y++ {
				if isValidCoord(x, y) && isMe(b[x][y], p) && !canBeCaptured(b, x, y, p) {
					count += 1
					if count == 5 {
						return true
					}
				} else {
					count = 0
				}
			}

			count = 0
			for x, y := x0-4, y0-4; x <= x0+4 && y <= y0+4; x, y = x+1, y+1 {
				if isValidCoord(x, y) && isMe(b[x][y], p) && !canBeCaptured(b, x, y, p) {
					count += 1
					if count == 5 {
						return true
					}
				} else {
					count = 0
				}
			}

			count = 0
			for x, y := x0-4, y0+4; x <= x0+4 && y >= y0-4; x, y = x+1, y-1 {
				if isValidCoord(x, y) && isMe(b[x][y], p) && !canBeCaptured(b, x, y, p) {
					count += 1
					if count == 5 {
						return true
					}
				} else {
					count = 0
				}
			}
		}
	}

	return false
}

func isValidMove(b [][]int, x, y int, p int, newBoard *[][]int) error {
	if !isValidCoord(x, y) {
		return fmt.Errorf("Invalid coordonate")
	} else if !isEmpty(b[x][y]) {
		return fmt.Errorf("Case already taken")
	} else {
		if !capture(b, x, y, p, newBoard) && introduceDoubleThree(b, x, y, p) {
			return fmt.Errorf("Illegal move")
		}
	}
	return nil
}

func move(b [][]int, c coord, p int, newBoard *[][]int) error {
	x, y := c.X, c.Y
	err := isValidMove(b, x, y, p, newBoard)
	if err != nil {
		return err
	}
	if len(*newBoard) == 0 {
		*newBoard = b
	}
	b[x][y] = p + 1
	if isGameOver(b, c, p) {
		return fmt.Errorf("Game Over")
	}
	return nil
}
