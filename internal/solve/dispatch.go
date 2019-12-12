package solve

import (
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
)

func DispatchSolve(moves []makemove.RubikMoves) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	rubik := makemove.InitRubik()

	rubik = rubik.DoMoves(moves)

	sequence = Iddfs(rubik)
	//	sequence = MechanicalHuman(rubik, true)
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// sequence = IdaStar(rubik, ScoringHamming)
	// sequence = IdaStar(rubik, ScoringHamming)
	rubik = rubik.DoMoves(sequence)
	fmt.Println(rubik)
	//sequence = []makemove.RubikMoves{}
	return sequence
}
