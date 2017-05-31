package main

import "fmt"

var WIDTH = 19
var HEIGHT = 19

type fnc func(x, y int, tmp [][]int, p int) bool

func isEmptyNew(x, y int, board [][]int, p int) bool {
	return board[x][y] == 0
}

func isMeNew(x, y int, board [][]int, player int) bool {
	return board[x][y] == player+1
}

func isEnemyNew(x, y int, board [][]int, player int) bool {
	return !isEmptyNew(x, y, board, player) && board[x][y] != player+1
}
func isValidCoord(x, y int) bool {
	return x >= 0 && x < HEIGHT && y >= 0 && y < WIDTH
}

func getScore(board [][]int, player int) int {
	score := 0
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			if isMeNew(x, y, board, player) {
				tmpx := x + 1
				tmpy := y
				horScore := 0
				horSpace := 0
				verScore := 0
				verSpace := 0
				diagScore1 := 0
				diagSpace1 := 0
				diagScore2 := 0
				diagSpace2 := 0
				b := 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx++ {
					if isMeNew(tmpx, tmpy, board, player) {
						horScore += b
						b++
					}
					horSpace++
				}
				tmpx = x - 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx-- {
					if isMeNew(tmpx, tmpy, board, player) {
						horScore += b
						b++
					}
					horSpace++
				}
				tmpx = x
				tmpy = y + 1
				b = 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpy++ {
					if isMeNew(tmpx, tmpy, board, player) {
						verScore += b
						b++
					}
					verSpace++
				}
				tmpy = y - 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpy-- {
					if isMeNew(tmpx, tmpy, board, player) {
						verScore += b
						b++
					}
					verSpace++
				}
				b = 1
				tmpy = y - 1
				tmpx = x - 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx-1, tmpy-1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore1 += b
						b++
					}
					diagSpace1++
				}
				tmpx = x + 1
				tmpy = y + 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx+1, tmpy+1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore1 += b
						b++
					}
					diagSpace1++
				}
				b = 1
				tmpx = x + 1
				tmpy = y - 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx+1, tmpy-1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore2 += b
						b++
					}
					diagSpace2++
				}
				tmpx = x - 1
				tmpy = y + 1
				for ; isValidCoord(tmpx, tmpy) && (isMeNew(tmpx, tmpy, board, player) || isEmptyNew(tmpx, tmpy, board, player)); tmpx, tmpy = tmpx-1, tmpy+1 {
					if isMeNew(tmpx, tmpy, board, player) {
						diagScore2 += b
						b++
					}
					diagSpace2++
				}
				//log.Println(x, y, "horScore= ", horScore, "verScore= ", verScore, "dia1 score= ", diagScore1, "dia2score = ", diagScore2, "horspace= ", horSpace, "verspace= ", verSpace, "dia1space= ", diagSpace1, "diagspace2= ", diagSpace2)

				if horSpace < 4 {
					horScore = 0
				}
				if verSpace < 4 {
					verScore = 0
				}
				if diagSpace1 < 4 {
					diagScore1 = 0
				}
				if diagSpace2 < 4 {
					diagScore2 = 0
				}
				score += horScore + verScore + diagScore1 + diagScore2
			}
		}
	}
	return score

}
func main() {
	board := make([][]int, 19)
	for i := 0; i < HEIGHT; i++ {
		board[i] = make([]int, 19)
		for j := 0; j < WIDTH; j++ {
			board[i][j] = 0
		}
	}

	board[6][4] = 1
	board[6][5] = 1
	board[6][6] = 1
	board[6][7] = 1
	fmt.Println(getScore(board, 1))
	fmt.Println(getScore(board, 0))
}
