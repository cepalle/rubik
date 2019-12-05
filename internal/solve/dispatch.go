package solve

import (
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
)

func DispatchSolve(moves []makemove.RubikMoves) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	rubik := makemove.InitRubik()

	rubik = rubik.DoMoves(moves)
	fmt.Println(rubik)

	// sequence = IdaStar(rubik, ScoringHamming)
	// sequence = IdaStar(rubik, ScoringHamming)
	sequence = Iddfs(rubik)
	rubik = rubik.DoMoves(sequence)
	fmt.Println(rubik)
	//sequence = []makemove.RubikMoves{}
	return sequence
}
