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
	var res string
	var algorithm int
	var load string
	var soluce []makemove.RubikMoves

	flag.StringVar(&moves, "m", "",
		"Moves that has to be done to shuffle the cube")
	flag.IntVar(&nbrMove, "r", 0,
		"Number of random move to shuffle the cube")
	flag.StringVar(&load, "l", "n",
		"Load the last moves so you can try with new algorithm (y or n)")
	flag.StringVar(&res, "re", "n",
		"To print help on human solver")
	flag.IntVar(&algorithm, "a", 4,
		`The algorithm you want to use to solve the rubik :
		1 -> Solving like a human, without steps.
		2 -> Using A*. The heuristic is the numbers of move needed by a human to solve the cube.
		3 -> Using A*. The heuristic is a BFS on depth 2 with a weight determined by the number of move needed by a human to solve the cube.
		4 -> thistlethwaite algorithm.
		5 -> Bidirectionnal BFS. Get the best number of move.
		6 -> Solving like a human, with steps to describe each moves`)
	flag.Parse()
	if nbrMove == 0 && moves == "" && load == "n" {
		log.Fatalf("Input error: Missing argument")
	}
	if nbrMove < 0 {
		log.Fatalf("Input error: Number of move to shuffle not valid\n")
	}
	if (nbrMove != 0 && len(moves) != 0) || (nbrMove != 0 && load == "y") || (len(moves) != 0 && load == "y") {
		log.Fatalf("Invalid input, either chose a random shuffle or write your own or load the old shuffle")
	}
	if load == "y" {
		soluce = solve.DispatchSolve(input.LoadSequence(), algorithm)
	} else if moves == "" {
		soluce = solve.DispatchSolve(input.GenerateRandomSequence(nbrMove), algorithm)
	} else {
		soluce = solve.DispatchSolve(input.StringToSequence(moves), algorithm)
	}
	fmt.Println(input.SequenceToString(soluce))
	if res != "n" {
		fmt.Println("nbr coups :", len(soluce))
	}
}
