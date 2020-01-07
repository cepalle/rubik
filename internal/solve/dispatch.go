package solve

import (
	"github.com/cepalle/rubik/internal/makemove"
	"os"
)

func DispatchSolve(moves []makemove.RubikMoves) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	rubik := makemove.InitRubik()

	rubik = rubik.DoMoves(moves)

	sequence = MechanicalHuman(rubik, true)
	if sequence == nil {
		os.Exit(1)
	}
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// sequence = IdaStar(rubik, ScoringHamming)
	// sequence = IdaStar(rubik, ScoringHamming)
	rubik = rubik.DoMoves(sequence)
	//sequence = []makemove.RubikMoves{}
	return sequence
}
