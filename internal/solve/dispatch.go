package solve

import (
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
	"os"
)

func DispatchSolve(moves []makemove.RubikMoves, algorithm int) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	rubik := makemove.InitRubik()

	rubik = rubik.DoMoves(moves)

	switch algorithm {
	case 1:
		sequence = MechanicalHuman(rubik, false)
	case 2:
		sequence = AStart(rubik, ScoringHuman)
	case 3:
		sequence = AStart(rubik, MakeBfsScore(2, ScoringHuman))
	case 4:
		sequence = Thistlethwaite(moves)
	case 5:
		sequence = BidiBfs(moves)
	case 6:
		sequence = MechanicalHuman(rubik, true)
	}

	finalSequence := CleanMoves(sequence)
	rubik = rubik.DoMoves(sequence)
	if rubik.IsResolve() {
		fmt.Println("Good job !")
	} else {
		os.Exit(1)
	}
	return finalSequence
}
