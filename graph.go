package  main

import (
  "math/rand"
  "time"
  "fmt"
)

type Graph struct {
  Name string
  TotalNodes int
  TotalEdges int
  Nodes []string
  LengthEdges []EdgeLength
}

type EdgeLength struct {
  Nodes []string
  Length float64
}

func randomPaths(selectedGraph Graph, pathLength int) []string {
  path := []string{}
  possibleNext := []string{}
  currentPoint := selectedGraph.Nodes[rand.Intn(len(selectedGraph.Nodes))]
  path = append(path, currentPoint)
  for i := pathLength; i > 1; i-- {
    possibleNext = possibleNext[:0]
    for _, value := range selectedGraph.LengthEdges {
      if (value.Nodes[0] == currentPoint) {
        possibleNext = append(possibleNext, value.Nodes[1])
      } else if (value.Nodes[1] == currentPoint) {
        possibleNext = append(possibleNext, value.Nodes[0])
      }
    }
    nextPoint := possibleNext[rand.Intn(len(possibleNext))]
    path = append(path, nextPoint)
    currentPoint = nextPoint
  }
  return path
}

func main() {
  rand.Seed(time.Now().UnixNano())
  ourGraph := Graph{
    Name: "thisGraph",
    TotalNodes: 8,
    TotalEdges: 10,
    Nodes: []string{"A", "B", "C", "D", "E", "F", "G"},
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
    },
  }
  lengthOfPaths := 5
  numPaths := 5
  
  paths := [][]string{}
  for i := 0; i < numPaths; i++ {
    paths = append(paths, randomPaths(ourGraph, lengthOfPaths))
  }
  for _, value := range paths {
    fmt.Println("Path: ", value)
  }
}