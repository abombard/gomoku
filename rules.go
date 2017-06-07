package main

import (
	"fmt"
)

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
				if count == 0 {
					blank = 0
				}
				count += 1
			}
			x1, y1 := x+checks[i][j][0], y+checks[i][j][1]
			if count == 3 {
				nextOk = isValidCoord(x1, y1) && isEmpty(b[x1][y1])
				break
			}
			if isValidCoord(x1, y1) {
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

func capture(b [][]int, x, y int, p int, newBoard *[][]int, pCapturedCount *int) bool {

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

				*pCapturedCount += 2
				ret = true
			}
		}
	}

	return ret
}

func canBeCaptured(b [][]int, x, y int) bool {

	p := b[x][y] - 1

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

func isGameOver(b [][]int, p int) bool {

	for x0 := 0; x0 < HEIGHT; x0++ {
		for y0 := 0; y0 < WIDTH; y0++ {

			// horizontal
			count := 0
			for x, y := x0, y0; x < x0+5; x++ {
				if isValidCoord(x, y) && isMe(b[x][y], p) && !canBeCaptured(b, x, y) {
					count += 1
					if count == 5 {
						return true
					}
				} else {
					count = 0
				}
			}

			// vertical
			count = 0
			for x, y := x0, y0; y < y0+5; y++ {
				if isValidCoord(x, y) && isMe(b[x][y], p) && !canBeCaptured(b, x, y) {
					count += 1
					if count == 5 {
						return true
					}
				} else {
					count = 0
				}
			}

			// diagonal
			count = 0
			for x, y := x0, y0; x < x0+5 && y < y0+5; x, y = x+1, y+1 {
				if isValidCoord(x, y) && isMe(b[x][y], p) && !canBeCaptured(b, x, y) {
					count += 1
					if count == 5 {
						return true
					}
				} else {
					count = 0
				}
			}

			// diagonal
			count = 0
			for x, y := x0, y0; x < x0+5 && y > y0-5; x, y = x+1, y-1 {
				if isValidCoord(x, y) && isMe(b[x][y], p) && !canBeCaptured(b, x, y) {
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

func isValidMove(b [][]int, x, y int) error {
	if !isValidCoord(x, y) {
		return fmt.Errorf("Invalid coordonate")
	} else if !isEmpty(b[x][y]) {
		return fmt.Errorf("Case already taken")
	}
	return nil
}

// the new move is added to both b and newBoard
// but the capture is only applied to newBoard
// if newBoard is nil
//   and a capture occure,
//     a copy of board is made and newBoard will point on it
//   if no capture occure,
//     newBoard will point on b
func move(b [][]int, c coord, p int, newBoard *[][]int, pCapturedCount *int) error {
	x, y := c.X, c.Y
	err := isValidMove(b, x, y)
	if err != nil {
		return err
	}
	if !capture(b, x, y, p, newBoard, pCapturedCount) && introduceDoubleThree(b, x, y, p) {
		return fmt.Errorf("Illegal move")
	}
	b[x][y] = p + 1
	if len(*newBoard) == 0 {
		*newBoard = b
	} else {
		(*newBoard)[x][y] = p + 1
	}
	if isGameOver(*newBoard, p) {
		return fmt.Errorf("Game Over")
	}
	return nil
}
