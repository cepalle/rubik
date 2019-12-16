package main

import (
	"encoding/gob"
	"fmt"
	"github.com/cepalle/rubik/internal/learn"
	"log"
	"math/rand"
	"os"
	"github.com/cepalle/rubik/internal/solve"
	"github.com/goml/gobrain"
	"github.com/goml/gobrain/persist"
)

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

	ff.Init(48, 48, learn.Bfs_depth+1)

	patterns := learn.MakePatterns(learn.Equalize(all, learn.Bfs_depth), learn.Bfs_depth)
	// fmt.Println(patterns)

	ff.Train(patterns, 1000, 0.001, 0.001, true)
	err := persist.Save("./ff.network", ff)
	if err != nil {
		log.Println("impossible to save network on file: ", err.Error())
	}
	// ff.Test(patterns)
}
