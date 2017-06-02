package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"
	//"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	WIDTH  = 19
	HEIGHT = 19
)

type Gomoku struct {
	Board [][]int
	Mode  string
}

var g Gomoku
var players [2]string
var current int = 0

type coord struct {
	X      int
	Y      int
	Player string
}

// Init
type start struct {
	Mode   string
	Player string
}

func resetBoard() {
	g.Board = make([][]int, HEIGHT)
	for x := 0; x < HEIGHT; x++ {
		g.Board[x] = make([]int, WIDTH)
		for y := 0; y < WIDTH; y++ {
			g.Board[x][y] = 0
		}
	}
}

func reset(w http.ResponseWriter, r *http.Request) {
	resetBoard()
	current = 0
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(g.Board)
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
	if g.Mode == "solo" {
		if players[0] == "" {
			players[0] = t.Player
			players[1] = "AI"
		}
	} else if g.Mode == "multi" {
		if players[0] == "" {
			players[0] = t.Player
		} else if players[1] == "" {
			players[1] = t.Player
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(g.Board)

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
	/*
		if t.Player != players[current] {
			http.Error(w, "Not your turn bitch", 400)
			return
		}
	*/
	err = move(g.Board, t, current, &g.Board)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	} else {
		current = (current + 1) % 2
		if g.Mode == "solo" {
			t = aiPlay()
			err = move(g.Board, t, current, &g.Board)
			current = (current + 1) % 2
			if err != nil {
				log.Fatal("AI", err)
				return
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(g.Board)
}

func main() {
	/*
		TESTHEURISTIC()
		return
	*/

	log.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))

	resetBoard()
	var entry string
	var static string
	var port string

	flag.StringVar(&entry, "entry", "index.html", "the entrypoint to serve.")
	flag.StringVar(&static, "static", ".", "the directory to serve static files from.")
	flag.StringVar(&port, "port", "8080", "the `port` to listen on.")
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
		//WriteTimeout: 15 * time.Second,
		//ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}
