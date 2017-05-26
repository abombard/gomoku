package main

import (
	"log"
	"math/rand"
)

func aiPlay() {
	log.Println("AI PLAY")
	g.board[rand.Intn(19)][rand.Intn(19)] = 1
}
