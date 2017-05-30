package main

import "github.com/twmb/algoimpl/go/graph"

type step struct {
	board [HEIGHT][WIDTH]int
}

func minmax() {
	gr := graph.New(graph.Directed)
	nodes := make(map[step]graph.Node, 0)

	nextSteps := getPossiblePlays()
	for i := range nextSteps {
		nodes[nextSteps[i]] = gr.MakeNode()
	}
}
