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

	// sequence = Iddfs_it_hamming(rubik)
	// sequence = AStart(rubik, MakeNNScoring(learn.Nnfilename))
	sequence = AStart(rubik, MakeNNDeepScoring(NnDeepFilename))
	// sequence = Bfs(rubik)

	// sequence = MechanicalHuman(rubik, true)
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
