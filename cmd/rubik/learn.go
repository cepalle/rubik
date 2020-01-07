package main

import (
	"encoding/gob"
	"fmt"
	"math/rand"
	"os"
	"github.com/cepalle/rubik/internal/solve"
	"github.com/goml/gobrain"
)

func makePatterns(bfsRes []solve.NodeExp) [][][]float64 {
	var patterns [][][]float64

	for _, e := range bfsRes {
		var input []float64

		for i := 0; i < 12; i++ {
			input = append(input, float64(e.Cube.PosP2[i]))
		}
		for i := 0; i < 12; i++ {
			input = append(input, float64(e.Cube.RotP2[i]))
		}
		for i := 0; i < 24; i++ {
			input = append(input, float64(e.Cube.PosFP3[i]))
		}

		line := [][]float64{input, {float64(e.Depth)}}
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

	fmt.Println(all)

	rand.Seed(0)

	ff := &gobrain.FeedForward{}

	ff.Init(48, 10, 1)

	patterns := makePatterns(all)
	fmt.Println(patterns)

	ff.Train(patterns, 1000, 0.6, 0.4, true)
	ff.Test(patterns)

}
