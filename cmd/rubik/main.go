package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var moves string
	var nbrMove int
	// var soluce []RubikMoves

	flag.StringVar(&moves, "m", "",
		"Moves that has to be done to shuffle the cube")
	flag.IntVar(&nbrMove, "r", 0,
		"Number of random move to shuffle the cube")
	flag.Parse()
	if nbrMove < 0 {
		fmt.Fprintf(os.Stderr, "Number of move to shuffle not valid\n")
		os.Exit(1)
	}
	if nbrMove != 0 && len(moves) != 0 {
		fmt.Fprintf(os.Stderr, "Invalid input, either chose a random shuffle or write your own, random shuffle ignored\n")
	}
	if moves == "" {
		//soluce =solve(GenerateRandom(nbrMove))
	} else {
		//soluce =solve(GenerateFromString(nbrMove))
	}
}
