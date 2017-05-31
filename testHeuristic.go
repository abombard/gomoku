package main

import "log"

func TESTHEURISTIC() {
	resetBoard()

	g.Board[5][5] = 1
	g.Board[5][6] = 1
	g.Board[5][7] = 1

	g.Board[7][1] = 2
	g.Board[10][2] = 2
	g.Board[15][1] = 2

	b := boardSlice()
	score := heuristic(b, 0)
	log.Println("score", score)
}
