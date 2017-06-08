package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	GameStateInit    = iota
	GameStateRunning = iota
	GameStateOver    = iota
)

type Room struct {
	GameState int
	Mode      string
	Board     [][]int

	History [][][]int

	Players        [2]player
	PlayerCountMax int
	PlayerCount    int
	Current        int

	Time time.Duration
}

func newBoard() [][]int {
	board := make([][]int, HEIGHT)
	for x := 0; x < HEIGHT; x++ {
		board[x] = make([]int, WIDTH)
		for y := 0; y < WIDTH; y++ {
			board[x][y] = 0
		}
	}
	return board
}

func NewRoom() *Room {

	room := &Room{
		GameState:      GameStateInit,
		Mode:           "",
		Board:          newBoard(),
		History:        make([][][]int, 0),
		Players:        [2]player{{Name: "", Score: 0, Index: 0}, {Name: "", Score: 0, Index: 0}},
		PlayerCount:    0,
		PlayerCountMax: 2,
		Current:        0,
		Time:           0,
	}

	return room
}

func (room *Room) SetMode(mode string) error {

	if mode != "solo" && mode != "multi" {
		return fmt.Errorf("Invalid mode")
	}

	room.Mode = mode

	if mode == "solo" {
		room.PlayerCountMax = 1
		room.Players[1].Name = "IA"
	} else if mode == "multi" {
		room.PlayerCountMax = 2
	}

	return nil
}

func (room *Room) AddPlayer(name string) error {

	if room.GameState != GameStateInit {
		return fmt.Errorf("Another game is running")
	}

	if room.PlayerCount == room.PlayerCountMax {
		return fmt.Errorf("Room is full")
	}

	room.Players[room.PlayerCount].Name = name
	room.PlayerCount += 1

	return nil
}

func (room *Room) StartGame() error {

	if room.GameState != GameStateInit {
		return fmt.Errorf("Another game is running")
	}

	if room.PlayerCount != room.PlayerCountMax {
		return fmt.Errorf("Waiting for player to connect")
	}

	room.Board = newBoard()
	room.Current = 0

	room.History = make([][][]int, 0)
	room.History = append(room.History, boardCopy(room.Board))

	room.Players[0].Index = 0
	room.Players[1].Index = 0

	room.Players[0].Score = 0
	room.Players[1].Score = 0

	room.GameState = GameStateRunning

	return nil
}

func (room *Room) RestartGame() error {

	room.Board = newBoard()
	room.Current = 0

	room.History = make([][][]int, 0)
	room.History = append(room.History, boardCopy(room.Board))

	room.Players[0].Index = 0
	room.Players[1].Index = 0

	room.Players[0].Score = 0
	room.Players[1].Score = 0

	room.GameState = GameStateRunning

	return nil
}

type GameState struct {
	Board    [][]int
	Players  [2]player
	Current  int
	GameOver bool
	Winner   int
	Time     time.Duration
}

func (room *Room) GetState(UserName string) GameState {

	state := GameState{
		Board:    room.Board,
		Players:  room.Players,
		Current:  room.Current,
		GameOver: false,
		Winner:   0,
		Time:     room.Time,
	}

	if room.Players[0].Name == UserName {
		state.Board = room.History[room.Players[0].Index]
	} else if room.Players[1].Name == UserName {
		state.Board = room.History[room.Players[1].Index]
	}

	if room.GameState == GameStateOver {
		state.GameOver = true
		state.Winner = room.Current
	}

	return state
}

func (room *Room) SetWinner() {

	room.GameState = GameStateOver

}

func (room *Room) AIGetMove() coord {

	coord := minmax(room.Board, room.Current)

	return coord
}

func (room *Room) SwitchPlayer() {

	room.Current = (room.Current + 1) % 2

}

func (room *Room) MakeMove(mv coord) {

	room.History = append(room.History, boardCopy(room.Board))
	room.Board[mv.X][mv.Y] = room.Current
	room.SwitchPlayer()
}

func (room *Room) HistoryPrev(UserName string) {

	if room.Players[0].Name == UserName && room.Players[0].Index > 0 {
		room.Players[0].Index -= 1
	} else if room.Players[1].Name == UserName && room.Players[1].Index > 0 {
		room.Players[1].Index -= 1
	}

}

func (room *Room) HistoryNext(UserName string) {

	if room.Players[0].Name == UserName && room.Players[0].Index < len(room.History) {
		room.Players[0].Index += 1
	} else if room.Players[1].Name == UserName && room.Players[0].Index < len(room.History) {
		room.Players[1].Index += 1
	}

}

type create struct {
	RoomName string
	RoomMode string
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var req create
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	room, ok := hub[req.RoomName]
	if ok {
		http.Error(w, "Room already ereqist", 400)
		return
	}

	room = NewRoom()

	err = room.SetMode(req.RoomMode)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	hub[req.RoomName] = room

	log.Println("Room " + req.RoomName + " created")
}

type join struct {
	RoomName string
	UserName string
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var req join
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	room, ok := hub[req.RoomName]
	if !ok {
		http.Error(w, "This room "+req.RoomName+" does not exist", 400)
		return
	}

	room.AddPlayer(req.UserName)

	roomState := room.GetState(req.UserName)
	intSendGameState(w, r, roomState)

	log.Println(req.UserName + " joined room " + req.RoomName)
}

func sendRooms(w http.ResponseWriter, r *http.Request) {

	var rooms []string

	for roomName, _ := range hub {
		rooms = append(rooms, roomName)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(rooms)

}
