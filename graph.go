package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Graph struct {
	Name        string
	TotalNodes  int
	TotalEdges  int
	Nodes       []string
	LengthEdges []EdgeLength
}

type EdgeLength struct {
	Nodes []string
	// hasn't been implemented for use yet
	Length float64
}

func randomPaths(myGraph Graph, pathLength int, ch chan []string, chFinished chan bool) {

	// stores our path that is being generated
	path := []string{}
	possibleNext := []string{}

	// generate first node as starting point
	currentPoint := myGraph.Nodes[rand.Intn(len(myGraph.Nodes))]
	path = append(path, currentPoint)

	// send finished message to channel
	defer func() {
		chFinished <- true
	}()

	// generate subsequent nodes up to len(path) == pathLength
	for i := pathLength; i > 1; i-- {
		possibleNext = possibleNext[:0]
		for _, value := range myGraph.LengthEdges {
			if value.Nodes[0] == currentPoint {
				possibleNext = append(possibleNext, value.Nodes[1])
			} else if value.Nodes[1] == currentPoint {
				possibleNext = append(possibleNext, value.Nodes[0])
			}
		}
		nextPoint := possibleNext[rand.Intn(len(possibleNext))]
		path = append(path, nextPoint)
		currentPoint = nextPoint
	}

	// send path to channel
	ch <- path
}

func main() {
	rand.Seed(time.Now().UnixNano())
	timeStart := time.Now()
	ourGraph := Graph{
		Name:       "thisGraph",
		TotalNodes: 11,
		TotalEdges: 17,
		Nodes:      []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"},
		LengthEdges: []EdgeLength{
			{[]string{"A", "B"}, 20.0},
			{[]string{"A", "C"}, 32.0},
			{[]string{"B", "C"}, 14.0},
			{[]string{"C", "D"}, 47.0},
			{[]string{"C", "E"}, 13.0},
			{[]string{"B", "F"}, 19.0},
			{[]string{"D", "G"}, 38.0},
			{[]string{"F", "G"}, 11.0},
			{[]string{"E", "H"}, 18.0},
			{[]string{"G", "H"}, 15.0},
			{[]string{"E", "K"}, 11.0},
			{[]string{"F", "J"}, 18.0},
			{[]string{"D", "H"}, 15.0},
			{[]string{"C", "H"}, 15.0},
			{[]string{"D", "K"}, 11.0},
			{[]string{"F", "I"}, 18.0},
			{[]string{"G", "K"}, 15.0},
		},
	}

	// channels
	chPaths := make(chan []string)
	chFinished := make(chan bool)

	// pre-set values
	lengthOfPaths := 100
	numPaths := 100000

	// where we'll store our generated paths
	paths := [][]string{}

	// generate paths via goroutines
	for i := 0; i < numPaths; i++ {
		go randomPaths(ourGraph, lengthOfPaths, chPaths, chFinished)
	}

	// wait for finished messages in channel, also append path to paths
	for c := 0; c < numPaths; {
		select {
		case thisPath := <-chPaths:
			paths = append(paths, thisPath)
		case <-chFinished:
			c++
		}
	}

	// benchmarking purposes
	elapsed := time.Since(timeStart)

	// print paths
	for _, value := range paths {
		fmt.Println("Path: ", value)
	}
	fmt.Println("Computation took: ", elapsed)
}
