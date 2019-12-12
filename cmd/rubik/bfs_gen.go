package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"github.com/cepalle/rubik/internal/solve"
)

func main() {
	var all []solve.NodeExp

	all = solve.Bfs_explorer(1)

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
