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

func stringToMove(move string) (makemove.RubikMoves, error) {
	for _, rubikMoves := range makemove.AllRubikMovesWithName {
		if move == rubikMoves.Name {
			return rubikMoves.Move, nil
		}
	}
	return makemove.RubikMoves{}, errors.New(fmt.Sprintf("Input error: <%s> is not a valid move", move))
}

func moveToString(move makemove.RubikMoves) (string, error) {
	for _, rubikMoves := range makemove.AllRubikMovesWithName {
		if move == rubikMoves.Move {
			return rubikMoves.Name, nil
		}
	}
	return "", errors.New(fmt.Sprintf("You shouldn't get there"))
}

func GenerateRandomSequence(nbrMove int) []makemove.RubikMoves {
	var Sequence []makemove.RubikMoves

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nbrMove; i++ {
		tmp := makemove.AllRubikMovesWithName[rand.Intn(makemove.NbRubikMoves)]
		fmt.Printf("%s ", tmp.Name)
		Sequence = append(Sequence, tmp.Move)
	}
	fmt.Printf("\n")
	return Sequence
}

func SequenceToString(moves []makemove.RubikMoves) string {
	var output string

	for i, move := range moves {
		newMove, err := moveToString(move)
		if err != nil {
			log.Fatal(err)
		} else if len(moves) == i+1 {
			output += newMove
		} else {
			output += newMove + " "
		}
	}
	return output
}

func StringToSequence(moves string) []makemove.RubikMoves {
	var listMoves []makemove.RubikMoves

	split := strings.Split(moves, " ")
	for _, move := range split {
		rubikMove, err := stringToMove(move)
		if err != nil {
			log.Fatal(err)
		} else {
			listMoves = append(listMoves, rubikMove)
		}
	}
	return listMoves
}

func ReverseMove(sequence []makemove.RubikMoves) []makemove.RubikMoves {
	var res []makemove.RubikMoves

	for i := 0; i < len(sequence); i++ {
		res = append(res, sequence[len(sequence)-i-1])
	}
	return res
}
