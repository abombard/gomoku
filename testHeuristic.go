package main

import "log"

func p1(x, y int) {
	g.Board[x][y] = 1
}

func p2(x, y int) {
	g.Board[x][y] = 2
}

func TESTHEURISTIC() {

	resetBoard()

	p1(5, 5)
	p1(5, 6)
	p1(5, 7)
	p1(5, 8)

	p2(6, 5)
	p2(6, 6)
	p2(6, 7)
	p2(6, 8)

	score := heuristic2(g.Board, 1)
	log.Println("score", score)

	resetBoard()

	p1(5, 5)
	p1(5, 6)
	p1(5, 7)

	p2(5, 8)
	p2(6, 6)
	p2(6, 7)

	score = heuristic2(g.Board, 1)
	log.Println("score", score)
}
