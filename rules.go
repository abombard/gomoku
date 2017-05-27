package main

const (
	WIDTH = 19
	HEIGHT = 19
)

func isValidCoord(x, y int) bool {
	return x >= 0 && x < HEIGHT && y >= 0 && y < WIDTH
}

func isEnemy(x, y int, p int) bool {
	return g.Board[x][y] != p
}

func isMe(x, y int, p int) bool {
	return g.Board[x][y] == p
}

func Move(c coord) error {
	x := c.X
	y := c.Y

	p := 1
	if c.Player != players[0] {
		p = 2
	}

	// Captured
	checks := [][3][2]int{
		{ { -1, -1 }, { -2, -2 }, { -3, -3 } },
		{ { -1,  0 }, { -2,  0 }, { -3,  0 } },
		{ { -1,  1 }, { -2,  2 }, { -3,  3 } },
		{ {  0,  1 }, {  0,  2 }, {  0,  3 } },
		{ {  1,  1 }, {  2,  2 }, {  3,  3 } },
		{ {  1,  0 }, {  2,  0 }, {  3,  0 } },
		{ {  1, -1 }, {  2, -2 }, {  3, -3 } },
		{ {  0, -1 }, {  0, -2 }, {  0, -3 } },
	}
	for i := range checks {
		p1, p2, p3 := checks[i][0], checks[i][1], checks[i][2]
		x1, y1 := x + p1[0], y + p1[1]
		x2, y2 := x + p2[0], y + p2[1]
		x3, y3 := x + p3[0], y + p3[1]
		if isValidCoord(x1, y1) && isValidCoord(x2, y2) && isValidCoord(x3, y3) {
			if isEnemy(x1, y1, p) && isEnemy(x2, y2, p) && isMe(x3, y3, p) {
				// Successfully captured
				g.Board[x1][y1] = 0
				g.Board[x2][y2] = 0
			}
		}
	}



	return nil
}
