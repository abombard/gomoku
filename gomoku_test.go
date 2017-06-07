package main

import (
	"log"
	"testing"
)

func TestTimeConsuming(t *testing.T) {

	resetBoard()
	for {
		t := aiPlay()
		pCapturedCount := 0
		err := move(g.Board, t, g.Current, &g.Board, &pCapturedCount)
		if err != nil {
			log.Println(err)
			break
		}
		g.Current = (g.Current + 1) % 2
		if g.GameOver {
			break
		}
	}
}
