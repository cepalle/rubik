package input

import (
	"errors"
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
	"log"
	"math/rand"
	"strings"
)

func getMoves(move string) (makemove.RubikMoves, error) {
	for _, rubikMoves := range makemove.AllRubikMoves {
		if move == rubikMoves.name {
			return rubikMoves.move, nil
		}
	}
	return makemove.RubikMoves{}, errors.New(fmt.Sprintf("Input error: <%s> is not a valid move", move))
}

func GenerateRandom(nbrMove int) []makemove.RubikMoves {
	var Sequence []makemove.RubikMoves

	for i := 0; i < nbrMove; i++ {
		Sequence[i] = makemove.AllRubikMoves[rand.Intn(makemove.NbRubikMoves)].move
	}
	return Sequence
}

func GenereateFromString(moves string) []makemove.RubikMoves {
	var listMoves []makemove.RubikMoves
	split := strings.Split(moves, " ")
	for _, move := range split {
		rubikMove, err := getMoves(move)
		if err != nil {
			log.Fatal(err)
		} else {
			listMoves = append(listMoves, rubikMove)
		}
	}
	return listMoves
}
