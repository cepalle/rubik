package main

import (
	"flag"
	"fmt"
	"github.com/cepalle/rubik/internal/input"
	"github.com/cepalle/rubik/internal/makemove"
	"github.com/cepalle/rubik/internal/solve"
	"log"
)

func main() {
	var moves string
	var nbrMove int
	var nbr bool
	var algorithm int
	var load int64
	var soluce []makemove.RubikMoves

	flag.StringVar(&moves, "m", "",
		"Sequence of move that has to be done to shuffle the cube")
	flag.IntVar(&nbrMove, "r", 0,
		"Number of random move to shuffle the cube")
	flag.Int64Var(&load, "l", -1,
		"Load a given seed for the random generator")
	flag.BoolVar(&nbr, "n", false,
		"To print the number of move done")
	flag.IntVar(&algorithm, "a", 4,
		`The algorithm you want to use to solve the rubik :
		1 -> Solving like a human, without steps.
		2 -> Using A*. The heuristic is the numbers of move needed by a human to solve the cube.
		3 -> Using A*. The heuristic is a BFS on depth 2 with a weight determined by the number of move needed by a human to solve the cube.
		4 -> thistlethwaite algorithm.
		5 -> Bidirectionnal BFS. Get the best number of move.
		6 -> Solving like a human, with steps to describe each moves`)
	flag.Parse()
	if nbrMove == 0 && moves == "" && load == -1 {
		log.Fatalf("Input error: Missing argument")
	}
	if algorithm < 1 || algorithm > 6 {
		log.Fatalf("Wrong choice of algorithm")
	}
	if nbrMove < 0 {
		log.Fatalf("Input error: Number of move to shuffle not valid\n")
	}
	if (nbrMove != 0 && len(moves) != 0) || (len(moves) != 0 && load != -1) {
		log.Fatalf("Invalid input, either chose a random shuffle or write your own or load the old shuffle")
	}
	if moves == "" || load != -1 {
		soluce = solve.DispatchSolve(input.GenerateRandomSequence(load, nbrMove), algorithm)
	} else {
		soluce = solve.DispatchSolve(input.StringToSequence(moves), algorithm)
	}
	fmt.Println(input.SequenceToString(soluce))
	if nbr {
		fmt.Println("nbr coups :", len(soluce))
	}
}
