package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Gomoku struct {
	Board [19][19]int
	Mode  string
}

var g Gomoku
var player string
var c int

type coord struct {
	X      int
	Y      int
	Player string
}

type start struct {
	Mode string
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "GOMOKU")
}

func reset(w http.ResponseWriter, r *http.Request) {
	resetBoard()
	c = 0
	w.WriteHeader(200)
}

func startGame(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t start
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	g.Mode = t.Mode

}

func play(w http.ResponseWriter, r *http.Request) {
	if g.Mode == "" {
		http.Error(w, "No mode selected yet", 400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var t coord
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	if t.X < 0 || t.X > 18 || t.Y < 0 || t.Y > 18 {
		http.Error(w, "Out of board range", 400)
		return
	}
	if g.Board[t.X][t.Y] == 2 || g.Board[t.X][t.Y] == 1 {
		http.Error(w, "Already played here", 400)
		return
	}
	if player == t.Player {
		http.Error(w, "Already Played", 400)
		return
	}
	player = t.Player
	if g.Mode == "multi" {
		if c%2 == 0 {
			g.Board[t.X][t.Y] = 2
		} else {
			g.Board[t.X][t.Y] = 1
		}
		c++
	} else if g.Mode == "solo" {
		g.Board[t.X][t.Y] = 2
		aiPlay()
		player = "AI"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(g.Board)

}

func resetBoard() {
	for x := range g.Board {
		for y := range g.Board {
			g.Board[x][y] = 0
		}
	}
}

func main() {
	log.Println("Start")
	resetBoard()
	http.HandleFunc("/", hello)
	http.HandleFunc("/play", play)
	http.HandleFunc("/startgame", startGame)
	http.HandleFunc("/reset", reset)
	http.ListenAndServe(":8000", nil)
}
