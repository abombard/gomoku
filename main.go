package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Gomoku struct {
	board [19][19]int
}

var g Gomoku

type coord struct {
	X      int
	Y      int
	Player string
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "GOMOKU")
}

func reset(w http.ResponseWriter, r *http.Request) {
	resetBoard()
	w.WriteHeader(200)
}

func play(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t coord
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	if g.board[t.X][t.Y] == 2 || g.board[t.X][t.Y] == 1 {
		http.Error(w, "Already played here", 400)
	}
	g.board[t.X][t.Y] = 2
	aiPlay()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(g.board)

}

func resetBoard() {
	for x := range g.board {
		for y := range g.board {
			g.board[x][y] = 0
		}
	}
}

func main() {
	log.Println("Start")
	resetBoard()
	http.HandleFunc("/", hello)
	http.HandleFunc("/play", play)
	http.HandleFunc("/reset", reset)
	http.ListenAndServe(":8000", nil)
}
