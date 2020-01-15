package solve

import (
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
)

func DispatchSolve(moves []makemove.RubikMoves, help string) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	rubik := makemove.InitRubik()

	rubik = rubik.DoMoves(moves)
	fmt.Println(rubik)

	// sequence = Bfs(rubik)
	// sequence = IddfsItHamming(rubik)
	// sequence = AStart(rubik, MakeNNScoring(Nnfilename))
	// sequence = AStart(rubik, MakeNNDeepScoring(NnDeepFilename))
	// sequence = Bfs(rubik)
	sequence = Thistlethwaite(moves)

	// sequence = MechanicalHuman(rubik, true)
	// fmt.Println()
	/*
	if help == "n" {
		sequence = MechanicalHuman(rubik, false)
	} else {
		sequence = MechanicalHuman(rubik, true)
	}
	if sequence == nil {
		os.Exit(1)
	}
	if help != "n" {
		fmt.Println()
	}
	*/
	// fmt.Println()
	// fmt.Println()
	// sequence = IdaStar(rubik, ScoringHamming)
	// sequence = IdaStar(rubik, ScoringHamming)
	rubik = rubik.DoMoves(sequence)
	//sequence = []makemove.RubikMoves{}
	return sequence
}
