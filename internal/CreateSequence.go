package internal

import (
	"math/rand"
	"strings"
)

func getMoves(move string) RubikMoves {
	for _, rubikMove := range AllRubikMovesWithName {
		if move == rubikMove.name {
			return rubikMove.move
		}
	}
	return nil
}

func GenerateRandom(nbrMove int) []RubikMoves {
	var sequence []RubikMoves

	for i := 0; i < nbrMove; i++ {
		sequence = append(sequence, AllRubikMovesWithName[rand.Intn(NbRubikMoves)].move)
	}
	return sequence
}

func GenereateFromString(moves string) []RubikMoves {
	var listMoves []RubikMoves
	split := strings.Split(moves, " ")
	for _, move := range split {
		rubikMove := getMoves(move)
		if rubikMove == nil {
			//todo
			return nil
		} else {
			listMoves = append(listMoves, rubikMove)
		}
	}
	return nil
}
