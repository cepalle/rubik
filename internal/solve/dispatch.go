package solve

import (
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
)

func DispatchSolve(moves []makemove.RubikMoves) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	rubik := makemove.InitRubik()

	rubik = rubik.DoMoves(moves)

	sequence = MechanicalHuman(rubik, true)
	rubik = rubik.DoMoves(sequence)
	fmt.Println(rubik)
	//sequence = []makemove.RubikMoves{}
	return sequence
}
