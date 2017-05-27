package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"time"
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(g.Board)

}

func checkDoubleThree(t coord) error {

	return nil
}

func processPlay(t coord) error {
	ret := checkDoubleThree(t)
	if ret != nil {
		return ret
	}

	return ret
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
	processPlay(t)
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
	resetBoard()
	var entry string
	var static string
	var port string

	flag.StringVar(&entry, "entry", "index.html", "the entrypoint to serve.")
	flag.StringVar(&static, "static", ".", "the directory to serve static files from.")
	flag.StringVar(&port, "port", "8000", "the `port` to listen on.")
	flag.Parse()

	r := mux.NewRouter()

	// Note: In a larger application, we'd likely extract our route-building logic into our handlers
	// package, given the coupling between them.

	// It's important that this is before your catch-all route ("/")
	//	api := r.PathPrefix("/api/v1/").Subrouter()
	r.HandleFunc("/startgame", startGame).Methods("POST")
	r.HandleFunc("/play", play).Methods("POST")
	r.HandleFunc("/reset", reset).Methods("GET")
	// Optional: Use a custom 404 handler for our API paths.
	// api.NotFoundHandler = JSONNotFound

	// Serve static assets directly.
	r.PathPrefix("/dist").Handler(http.FileServer(http.Dir(static)))

	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	r.HandleFunc("/", IndexHandler(entry))

	srv := &http.Server{
		Handler: handlers.LoggingHandler(os.Stdout, r),
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}
