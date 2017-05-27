package main

import (
	"log"
	"math/rand"
	"time"
)

func aiPlay() {
	log.Println("AI PLAY")
	g.Board[rand.Intn(19)][rand.Intn(19)] = current + 1
	time.Sleep(10000 * time.Millisecond)
}
