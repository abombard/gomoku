package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	WIDTH  = 19
	HEIGHT = 19
)

type Gomoku struct {
	Board    [][]int
	Current  int
	Mode     string
	Time     time.Duration
	Players  [2]player
	GameOver bool
	Winner   int
}

var g Gomoku
var iaPlaying = false

type player struct {
	Name  string
	Score int
}
type coord struct {
	X      int
	Y      int
	Player string
}

func sendState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(g)
}

func resetBoard() {
	g.Board = make([][]int, HEIGHT)
	for x := 0; x < HEIGHT; x++ {
		(g.Board)[x] = make([]int, WIDTH)
		for y := 0; y < WIDTH; y++ {
			(g.Board)[x][y] = 0
		}
	}
}

func resetGame(gc *Gomoku) {
	resetBoard()
	gc.Current = 0
	gc.Mode = ""
	gc.Time = 0
	gc.Players[0].Name = ""
	gc.Players[0].Score = 0
	gc.Players[1].Name = ""
	gc.Players[1].Score = 0
	gc.GameOver = false
	gc.Winner = 0
}

func reset(w http.ResponseWriter, r *http.Request) {
	resetBoard()
	g.Current = 0
	g.Time = 0
	g.Players[0].Score = 0
	g.Players[1].Score = 0
	g.GameOver = false
	g.Winner = 0
	sendState(w, r)
}

func setWinner(gc *Gomoku) {
	gc.GameOver = true
	gc.Winner = gc.Current
}

// Init
type start struct {
	Mode   string
	Player string
}

func startGame(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t start
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	if g.Mode != t.Mode {
		resetGame(&g)
		g.Mode = t.Mode
	}
	if g.Mode == "solo" {
		if g.Players[0].Name == "" || (g.Players[0].Name != "" && t.Player != g.Players[0].Name) {
			g.Players[0].Name = t.Player
			g.Players[1].Name = "AI"
		}
	} else if g.Mode == "multi" {
		if g.Players[0].Name == "" {
			g.Players[0].Name = t.Player
		} else if g.Players[1].Name == "" {
			g.Players[1].Name = t.Player
		} else {
			resetGame(&g)
			g.Players[0].Name = t.Player
		}
	}
	sendState(w, r)
}

func getBoard(w http.ResponseWriter, r *http.Request) {
	if iaPlaying {
		http.Error(w, "AI PLAYING", 400)
		return

	}
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
	if t.Player != g.Players[g.Current].Name {
		http.Error(w, "Not your turn bitch", 400)
		return
	}
	pCapturedCount := 0
	err = move(g.Board, t, g.Current, &g.Board, &pCapturedCount)
	if err != nil {
		log.Println(err)
		if err.Error() == "Game Over" {
			setWinner(&g)
			sendState(w, r)
		} else {
			http.Error(w, err.Error(), 400)
		}
		return
	}
	g.Players[g.Current].Score += pCapturedCount
	if g.Players[g.Current].Score >= 10 {
		setWinner(&g)
		sendState(w, r)
		return
	}
	g.Current = (g.Current + 1) % 2
	sendState(w, r)
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
	if g.Mode == "solo" && !g.GameOver {
		iaPlaying = true
		start := time.Now()
		t = aiPlay()
		g.Time = (time.Since(start) / 1000000)

		pCapturedCount := 0
		err = move(g.Board, t, g.Current, &g.Board, &pCapturedCount)
		if err != nil {
			if err.Error() == "Game Over" {
				setWinner(&g)
				sendState(w, r)
			} else {
				http.Error(w, err.Error(), 400)
			}
			iaPlaying = false
			return
		}
		g.Players[g.Current].Score += pCapturedCount
		if g.Players[g.Current].Score >= 10 {
			setWinner(&g)
			sendState(w, r)
			iaPlaying = false
			return
		}
		g.Current = (g.Current + 1) % 2
	}
	iaPlaying = false
	sendState(w, r)
}

func hint(w http.ResponseWriter, r *http.Request) {
	if g.Mode == "" {
		http.Error(w, "No mode selected yet", 400)
		return
	}
	var t coord
	if !g.GameOver {
		iaPlaying = true
		t = aiPlay()
		g.Board[t.X][t.Y] = -1
	}
	g.Board[t.X][t.Y] = 0
	iaPlaying = false
	sendState(w, r)
}

func main() {
	/*
		TESTHEURISTIC()
		return
	*/

	log.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))

	resetGame(&g)

	var entry string
	var static string
	var port string

	flag.StringVar(&entry, "entry", "index.html", "the entrypoint to serve.")
	flag.StringVar(&static, "static", ".", "the directory to serve static files from.")
	flag.StringVar(&port, "port", "3004", "the `port` to listen on.")
	flag.Parse()

	r := mux.NewRouter()

	// Note: In a larger application, we'd likely extract our route-building logic into our handlers
	// package, given the coupling between them.

	// It's important that this is before your catch-all route ("/")
	//	api := r.PathPrefix("/api/v1/").Subrouter()
	r.HandleFunc("/startgame", startGame).Methods("POST")
	r.HandleFunc("/play", play).Methods("POST")
	r.HandleFunc("/getboard", getBoard).Methods("POST")
	r.HandleFunc("/reset", reset).Methods("GET")
	r.HandleFunc("/board", sendState).Methods("GET")
	r.HandleFunc("/hint", hint).Methods("GET")
	// Optional: Use a custom 404 handler for our API paths.
	// api.NotFoundHandler = JSONNotFound

	// Serve static assets directly.
	r.PathPrefix("/dist").Handler(http.FileServer(http.Dir(static)))

	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	r.HandleFunc("/", IndexHandler(entry))

	srv := &http.Server{
		Handler: handlers.LoggingHandler(os.Stdout, r),
		Addr:    "0.0.0.0:" + port,
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
