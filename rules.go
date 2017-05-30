package main

import "fmt"

func isValidCoord(x, y int) bool {
	return x >= 0 && x < HEIGHT && y >= 0 && y < WIDTH
}

func isEmpty(x, y int) bool {
	return g.Board[x][y] == 0
}

func isEnemy(x, y int) bool {
	return !isEmpty(x, y) && g.Board[x][y] != current+1
}

func isMe(x, y int) bool {
	return !isEmpty(x, y) && g.Board[x][y] == current+1
}

func capture(x, y int, p int) {

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
			if isEnemy(x1, y1) && isEnemy(x2, y2) && isMe(x3, y3) {
				// Successfully captured
				g.Board[x1][y1] = 0
				g.Board[x2][y2] = 0
			}
		}
	}
}

func introduceDoubleThree(x, y int) bool {

	checks := [4][6][2]int{
		{{-3, -3}, {-2, -2}, {-1, -1}, {1, 1}, {2, 2}, {3, 3}},
		{{-3, 0}, {-2, 0}, {-1, 0}, {1, 0}, {2, 0}, {3, 0}},
		{{-3, 3}, {-2, 2}, {-1, 1}, {1, -1}, {2, -2}, {3, -3}},
		{{0, 3}, {0, 2}, {0, 1}, {0, -1}, {0, -2}, {0, -3}},
	}

	threeCount := 0
	for i := range checks {
		count := 0
		blank := 0
		for j := range checks[i] {
			if j == 3 {
				count += 1
				if count == 3 {
					break
				}
			}
			x1, y1 := x+checks[i][j][0], y+checks[i][j][1]
			if isValidCoord(x1, y1) {
				if isMe(x1, y1) {
					blank = 0
					count += 1
					if count == 3 {
						break
					}
				} else if isEnemy(x1, y1) {
					blank = 0
					count = 0
				} else {
					blank += 1
					if blank > 1 {
						count = 0
					}
				}
			}
		}
		if count == 3 {
			threeCount += 1
		}
	}

	return threeCount > 1
}

func isValidMove(c coord) error {
	x, y := c.X, c.Y
	if !isValidCoord(x, y) {
		return fmt.Errorf("Invalid coordonate")
	} else if !isEmpty(x, y) {
		return fmt.Errorf("Case already taken")
	} else if introduceDoubleThree(x, y) {
		return fmt.Errorf("Illegal move")
	}
	return nil
}

func tryCapture(x, y int) {

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
			if isEnemy(x1, y1) && isEnemy(x2, y2) && isMe(x3, y3) {
				// Successfully captured
				g.Board[x1][y1] = 0
				g.Board[x2][y2] = 0
			}
		}
	}
}

func move(c coord) {
	err := isValidMove(c)
	if err != nil {
		return
	}
	x, y := c.X, c.Y
	tryCapture(x, y)
	g.Board[x][y] = current + 1
	current = (current + 1) % 2
}

func canBeCaptured(x, y int) bool {

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
			if isMe(x2, y2) && ((isEmpty(x1, y1) && isEnemy(x3, y3)) || (isEnemy(x1, y1) && isEmpty(x3, y3))) {
				return true
			}
		}
		if isValidCoord(x2, y2) && isValidCoord(x3, y3) && isValidCoord(x4, y4) {
			if isMe(x3, y3) && ((isEmpty(x2, y2) && isEnemy(x4, y4)) || (isEnemy(x2, y2) && isEmpty(x4, y4))) {
				return true
			}
		}
	}
	return false
}

func isGameOver(c coord) bool {
	x, y := c.X, c.Y

	if canBeCaptured(x, y) {
		println("New move can be captured")
		return false
	}

	checks := [4][8][2]int{
		{{-4, -4}, {-3, -3}, {-2, -2}, {-1, -1}, {1, 1}, {2, 2}, {3, 3}, {4, 4}},
		{{-4, 0}, {-3, 0}, {-2, 0}, {-1, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}},
		{{-4, 4}, {-3, 3}, {-2, 2}, {-1, 1}, {1, -1}, {2, -2}, {3, -3}, {4, -4}},
		{{0, -4}, {0, -3}, {0, -2}, {0, -1}, {0, 1}, {0, 2}, {0, 3}, {0, 4}},
	}
	for i := range checks {
		count := 0
		for j := range checks[i] {
			x1, y1 := x+checks[i][j][0], y+checks[i][j][1]
			if isValidCoord(x1, y1) && isMe(x1, y1) && !canBeCaptured(x1, y1) {
				count += 1
				if count == 4 {
					return true
				}
			} else {
				count = 0
			}
		}
	}
	return false
}
