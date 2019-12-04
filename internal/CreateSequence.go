package internal

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
)

func getMoves(move string) (RubikMoves, error) {
	for _, rubikMoves := range AllRubikMoves {
		if move == rubikMoves.name {
			return rubikMoves.move, nil
		}
	}
	return RubikMoves{}, errors.New(fmt.Sprintf("Input error: <%s> is not a valid move", move))
}

func GenerateRandom(nbrMove int) []RubikMoves {
	var Sequence []RubikMoves

	for i := 0; i < nbrMove; i++ {
		Sequence[i] = AllRubikMoves[rand.Intn(NbRubikMoves)].move
	}
	return Sequence
}

func GenereateFromString(moves string) []RubikMoves {
	var listMoves []RubikMoves
	split := strings.Split(moves, " ")
	for _, move := range split {
		rubikMove, err := getMoves(move)
		if err != nil {
			log.Fatal(err)
			//todo
		} else {
			listMoves = append(listMoves, rubikMove)
		}
	}
	return listMoves
}
