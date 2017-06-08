package main

import "log"

func printBoard(b [][]int) {
	for x := 0; x < HEIGHT; x++ {
		log.Println(b[x])
	}
}

func (room *Room) p1(x, y int) {
	room.Board[x][y] = 1
}

func (room *Room) p2(x, y int) {
	room.Board[x][y] = 2
}

func (room *Room) test() {
	score := heuristic2(room.Board, 0)
	log.Println("score", score)
	printBoard(room.Board)
}

func TESTHEURISTIC() {

	room := NewRoom()

	room.RestartGame()

	room.p1(5, 5)
	room.p1(5, 6)

	room.p2(7, 1)
	room.p2(7, 7)

	room.test()

	room.RestartGame()

	room.p1(5, 5)
	room.p1(5, 6)
	room.p1(5, 7)

	room.p2(6, 5)
	room.p2(6, 6)
	room.p2(6, 7)

	room.test()

	room.RestartGame()

	room.p1(5, 5)
	room.p1(5, 6)
	room.p1(5, 7)

	room.p2(5, 8)
	room.p2(6, 6)
	room.p2(6, 7)

	room.test()

	room.RestartGame()

	room.p1(5, 5)
	room.p1(5, 6)
	room.p1(7, 7)

	room.p2(7, 4)
	room.p2(7, 5)
	room.p2(7, 6)

	room.test()

	room.RestartGame()

	room.p1(5, 5)
	room.p1(5, 6)
	room.p1(7, 10)

	room.p2(7, 4)
	room.p2(7, 5)
	room.p2(7, 6)

	room.test()

	room.RestartGame()

	room.p1(5, 5)
	room.p1(6, 5)
	room.p1(7, 5)

	room.p2(7, 6)
	room.p2(7, 7)
	room.p2(8, 7)
	room.p2(9, 7)

	room.p1(6, 7)
	room.p1(10, 7)

	room.p2(9, 8)
	room.p2(9, 9)

	room.p1(9, 10)

	room.test()

}
