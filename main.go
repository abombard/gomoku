package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

var hub map[string]*Room = make(map[string]*Room)

type player struct {
	Name  string
	Score int
	Index int
}

type coord struct {
	X      int
	Y      int
	Player string
}

func intSendGameState(w http.ResponseWriter, r *http.Request, state GameState) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(state)

}

type GameStateRequest struct {
	UserName string
	RoomName string
}

func unpackGameStateRequest(r *http.Request) (GameStateRequest, error) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var req GameStateRequest
	err := decoder.Decode(&req)

	return req, err
}

func getRoom(name string) (*Room, error) {

	if room, ok := hub[name]; ok {
		return room, nil
	}

	return nil, fmt.Errorf("Invalid room name")
}

func sendGameState(w http.ResponseWriter, r *http.Request) {

	req, err := unpackGameStateRequest(r)
	if err != nil {
		http.Error(w, "Invalid packet request", 400)
		return
	}

	room, err := getRoom(req.RoomName)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	roomState := room.GetState(req.UserName)
	intSendGameState(w, r, roomState)
}

func startGame(w http.ResponseWriter, r *http.Request) {

	req, err := unpackGameStateRequest(r)
	if err != nil {
		http.Error(w, "Invalid packet request", 400)
		return
	}

	room, ok := hub[req.RoomName]
	if !ok {
		http.Error(w, "Invalid room name", 400)
		return
	}

	if room.GameState == GameStateInit {
		room.StartGame()
	} else {
		room.RestartGame()
	}

	roomState := room.GetState(req.UserName)
	intSendGameState(w, r, roomState)
}

type PlayerMoveRequest struct {
	UserName string
	RoomName string
	X        int
	Y        int
}

func playerMove(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var req PlayerMoveRequest
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	room, ok := hub[req.RoomName]
	if !ok {
		http.Error(w, "Invalid room name", 400)
		return
	}

	if room.GameState != GameStateRunning {
		http.Error(w, "Game is not running", 400)
		return
	}

	if req.UserName != room.Players[room.Current].Name {
		http.Error(w, "Not your turn bitch", 400)
		return
	}

	if room.Players[room.Current].Index != room.HistoryLen-1 {
		http.Error(w, "You can't play in History", 400)
		return
	}

	mv := coord{X: req.X, Y: req.Y}

	pCapturedCount := 0
	err = move(room.Board, mv, room.Current, &room.Board, &pCapturedCount)
	if err != nil && err.Error() != "Game Over" {
		http.Error(w, err.Error(), 400)
		return
	}

	room.Players[room.Current].Score += pCapturedCount
	if (err != nil && err.Error() == "Game Over") ||
		room.Players[room.Current].Score >= 10 {
		room.SetWinner()
	} else {
		room.History = append(room.History, boardCopy(room.Board))
		room.HistoryLen += 1
		room.Players[room.Current].Index = room.HistoryLen - 1
		room.SwitchPlayer()
		room.Players[room.Current].Index = room.HistoryLen - 1
	}

	roomState := room.GetState(req.UserName)
	intSendGameState(w, r, roomState)
}

func IAMove(w http.ResponseWriter, r *http.Request) {

	req, err := unpackGameStateRequest(r)
	if err != nil {
		http.Error(w, "Invalid packet request", 400)
		return
	}

	room, err := getRoom(req.RoomName)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if room.Mode != "solo" {
		return
	}

	if room.GameState != GameStateRunning {
		http.Error(w, "Game is not running", 400)
		return
	}

	start := time.Now()
	mv := room.AIGetMove()
	room.Time = (time.Since(start) / 1000000)

	pCapturedCount := 0
	err = move(room.Board, mv, room.Current, &room.Board, &pCapturedCount)
	if err != nil && err.Error() != "Game Over" {
		http.Error(w, err.Error(), 400)
		return
	}

	room.Players[room.Current].Score += pCapturedCount
	if (err != nil && err.Error() == "Game Over") ||
		room.Players[room.Current].Score >= 10 {
		room.SetWinner()
	} else {
		room.History = append(room.History, boardCopy(room.Board))
		room.HistoryLen += 1
		room.Players[room.Current].Index = room.HistoryLen - 1
		room.SwitchPlayer()
		room.Players[room.Current].Index = room.HistoryLen - 1
	}

	roomState := room.GetState(req.UserName)
	intSendGameState(w, r, roomState)
}

func hint(w http.ResponseWriter, r *http.Request) {

	req, err := unpackGameStateRequest(r)
	if err != nil {
		http.Error(w, "Invalid packet request", 400)
		return
	}

	room, err := getRoom(req.RoomName)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if room.GameState != GameStateRunning {
		http.Error(w, "Waiting for the game to start", 400)
		return
	}

	mv := room.AIGetMove()
	room.Board[mv.X][mv.Y] = -1

	roomState := room.GetState(req.UserName)
	intSendGameState(w, r, roomState)

	room.Board[mv.X][mv.Y] = 0
}

func historyPrev(w http.ResponseWriter, r *http.Request) {

	req, err := unpackGameStateRequest(r)
	if err != nil {
		http.Error(w, "Invalid packet request", 400)
		return
	}

	room, err := getRoom(req.RoomName)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if room.GameState != GameStateRunning {
		http.Error(w, "Waiting for the game to start", 400)
		return
	}

	room.HistoryPrev(req.UserName)

	roomState := room.GetState(req.UserName)
	intSendGameState(w, r, roomState)
}

func historyNext(w http.ResponseWriter, r *http.Request) {

	req, err := unpackGameStateRequest(r)
	if err != nil {
		http.Error(w, "Invalid packet request", 400)
		return
	}

	room, err := getRoom(req.RoomName)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if room.GameState != GameStateRunning {
		http.Error(w, "Waiting for the game to start", 400)
		return
	}

	room.HistoryNext(req.UserName)

	roomState := room.GetState(req.UserName)
	intSendGameState(w, r, roomState)
}

func main() {
	/*
		TESTHEURISTIC()
		return
	*/

	log.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))

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
	r.HandleFunc("/GameStateRequest", sendGameState).Methods("POST")
	r.HandleFunc("/RoomsRequest", sendRooms).Methods("GET")
	r.HandleFunc("/CreateRoomRequest", CreateRoom).Methods("POST")
	r.HandleFunc("/JoinRoomRequest", JoinRoom).Methods("POST")
	r.HandleFunc("/StartGameRequest", startGame).Methods("POST")
	r.HandleFunc("/PlayerMoveRequest", playerMove).Methods("POST")
	r.HandleFunc("/IAMoveRequest", IAMove).Methods("POST")
	r.HandleFunc("/HintRequest", hint).Methods("POST")
	r.HandleFunc("/HistoryPrevRequest", historyPrev).Methods("POST")
	r.HandleFunc("/HistoryNextRequest", historyNext).Methods("POST")
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
