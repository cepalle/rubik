package main

import (
	"encoding/gob"
	"fmt"
	"github.com/cepalle/rubik/internal/solve"
	"os"
)

func main() {

	var all []solve.NodeExp

	all = solve.Bfs_explorer(solve.Bfs_depth)

	dataFile, err := os.Create("node_exp.gob")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dataEncoder := gob.NewEncoder(dataFile)
	err = dataEncoder.Encode(all)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = dataFile.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
