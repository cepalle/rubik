package main

import (
	"encoding/gob"
	"fmt"
	"math/rand"
	"os"
	"github.com/cepalle/rubik/internal/solve"
	"github.com/goml/gobrain"
)

func makePatterns(bfsRes []solve.NodeExp, bfs_depth uint32) [][][]float64 {
	var patterns [][][]float64

	for _, e := range bfsRes {
		var input []float64
		var output []float64

		for i := 0; i < 12; i++ {
			input = append(input, float64(e.Cube.PosP2[i]))
		}
		for i := 0; i < 12; i++ {
			input = append(input, float64(e.Cube.RotP2[i]))
		}
		for i := 0; i < 24; i++ {
			input = append(input, float64(e.Cube.PosFP3[i]))
		}
		for i := uint32(0); i <= bfs_depth; i++ {
			output = append(output, 0)
		}
		output[e.Depth] = 1

		line := [][]float64{input, output}
		patterns = append(patterns, line)
	}
	return patterns
}

func main() {
	var all []solve.NodeExp
	dataFile, err1 := os.Open("node_exp.gob")

	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err1 = dataDecoder.Decode(&all)

	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}

	err1 = dataFile.Close()
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}

	// fmt.Println(all)

	rand.Seed(0)

	ff := &gobrain.FeedForward{}

	ff.Init(48, 24, 4)

	patterns := makePatterns(all, 3)
	// fmt.Println(patterns)

	ff.Train(patterns, 100, 0.6, 0.4, true)
	ff.Test(patterns)
}