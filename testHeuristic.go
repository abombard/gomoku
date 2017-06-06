package main

import "log"

func printBoard(b [][]int) {
	for x := 0; x < HEIGHT; x++ {
		log.Println(b[x])
	}
}

func p1(x, y int) {
	g.Board[x][y] = 1
}

func p2(x, y int) {
	g.Board[x][y] = 2
}

func test() {
	score := heuristic2(g.Board, 0)
	log.Println("score", score)
	printBoard(g.Board)
}

func TESTHEURISTIC() {

	resetBoard()

	p1(5, 5)
	p1(5, 6)

	p2(7, 1)
	p2(7, 7)

	test()

	resetBoard()

	p1(5, 5)
	p1(5, 6)
	p1(5, 7)

	p2(6, 5)
	p2(6, 6)
	p2(6, 7)

	test()

	resetBoard()

	p1(5, 5)
	p1(5, 6)
	p1(5, 7)

	p2(5, 8)
	p2(6, 6)
	p2(6, 7)

	test()

	resetBoard()

	p1(5, 5)
	p1(5, 6)
	p1(7, 7)

	p2(7, 4)
	p2(7, 5)
	p2(7, 6)

	test()

	resetBoard()

	p1(5, 5)
	p1(5, 6)
	p1(7, 10)

	p2(7, 4)
	p2(7, 5)
	p2(7, 6)

	test()

	resetBoard()

	p1(5, 5)
	p1(6, 5)
	p1(7, 5)

	p2(7, 6)
	p2(7, 7)
	p2(8, 7)
	p2(9, 7)

	p1(6, 7)
	p1(10, 7)

	p2(9, 8)
	p2(9, 9)

	p1(9, 10)

	test()

}
