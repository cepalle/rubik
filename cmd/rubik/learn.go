package main

import (
	"encoding/gob"
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
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

		input = makemove.Rubik_to_nn_input(&e.Cube)
		for i := uint32(0); i <= bfs_depth; i++ {
			output = append(output, 0)
		}
		output[e.Depth] = 1

		line := [][]float64{input, output}
		patterns = append(patterns, line)
	}
	return patterns
}

func equalize(bfsRes []solve.NodeExp, bfs_depth uint32) []solve.NodeExp {
	var nbByDepth []uint32

	for i := uint32(0); i <= bfs_depth; i++ {
		nbByDepth = append(nbByDepth, 0)
	}
	for _, e := range bfsRes {
		nbByDepth[e.Depth]++
	}
	var mx uint32
	mx = 0

	for _, e := range nbByDepth {
		if e > mx {
			mx = e
		}
	}

	for i := uint32(0); i < uint32(len(nbByDepth)); i++ {
		for ; nbByDepth[i] < mx; {
			for j := int(0); j < len(bfsRes) && nbByDepth[i] < mx; j++ {
				if bfsRes[j].Depth == i {
					nbByDepth[i]++
					bfsRes = append(bfsRes, bfsRes[j])
				}
			}
		}

	}
	return bfsRes
}

func main() {
	const bfs_depth = 4

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

	ff.Init(48, 48, bfs_depth+1)

	patterns := makePatterns(equalize(all, bfs_depth), bfs_depth)
	fmt.Println(patterns)

	ff.Train(patterns, 1000, 0.001, 0.001, true)
	ff.Test(patterns)
}
