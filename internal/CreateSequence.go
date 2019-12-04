package internal

import (
	"math/rand"
	"strings"
)

func getMoves(move string) RubikMoves {
	for name, rubikMove := range AllRubikMoves {
		if move == name {
			return rubikMove
		}
	}
	return nil
}

func GenerateRandom(nbrMove int) []RubikMoves {
	var Sequence [nbrMove]RubikMoves

	for i := 0; i < nbrMove; i++ {
		Sequence[i] = AllRubikMoves[rand.Intn(NbRubikMoves)]
	}
	return Sequence
}

func GenereateFromString(moves string) []RubikMoves {
	var listMoves []RubikMoves
	split := strings.Split(moves, " ")
	for _, move := range split {
		rubikMove = getMoves(move)
		if rubikMove == nil {
			//todo
		} else {
			listMoves = append(listMoves, rubikMove)
		}
	}
	return listMoves
}
