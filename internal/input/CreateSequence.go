package input

import (
	"errors"
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
	"log"
	"math/rand"
	"strings"
	"time"
)

func getMoves(move string) (makemove.RubikMoves, error) {
	for _, rubikMoves := range makemove.AllRubikMoves {
		if move == rubikMoves.Name {
			return rubikMoves.Move, nil
		}
	}
	return makemove.RubikMoves{}, errors.New(fmt.Sprintf("Input error: <%s> is not a valid move", move))
}

func GenerateRandom(nbrMove int) []makemove.RubikMoves {
	var Sequence []makemove.RubikMoves

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nbrMove; i++ {
		Sequence = append(Sequence, makemove.AllRubikMoves[rand.Intn(makemove.NbRubikMoves)].Move)
	}
	fmt.Println(Sequence)
	return Sequence
}

func GenerateFromString(moves string) []makemove.RubikMoves {
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
	fmt.Println(listMoves)
	return listMoves
}
