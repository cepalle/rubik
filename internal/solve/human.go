package solve

import (
	"bufio"
	"fmt"
	"github.com/cepalle/rubik/internal/input"
	"github.com/cepalle/rubik/internal/makemove"
	"os"
)

func getIndex(lst []uint8, value uint8) int {
	for i, v := range lst {
		if v == value {
			return i
		}
	}
	return -1
}

func downEdges(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if debug {
		fmt.Println("Solving this cube :")
	}
	if debug {
		if len(sequence) != 0 {
			fmt.Println("Cube solved !")
		} else {
			fmt.Println("There was nothing to do")
		}
	}
	return sequence
}

func downCornersOrientation(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if debug {
		fmt.Println("Changing corners orientation :")
	}
	if debug {
		if len(sequence) != 0 {
			fmt.Println("The bottom edges should be in a good orientation now !")
		} else {
			fmt.Println("There was nothing to do")
		}
	}
	return sequence
}

func downCorners(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if debug {
		fmt.Println("Placing bottom corner at their respectives places, not regarding the orientation :")
	}
	if debug {
		if len(sequence) != 0 {
			fmt.Println("Corners placed !")
		} else {
			fmt.Println("There was nothing to do")
		}
	}
	return sequence
}

func downCross(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if debug {
		fmt.Println("Creating the bottom cross :")
	}
	if debug {
		if len(sequence) != 0 {
			fmt.Println("Bottom cross placed !")
		} else {
			fmt.Println("There was nothing to do")
		}
	}
	return sequence
}

func secondRow(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if debug {
		fmt.Println("Setting up the second row :")
	}
	if debug {
		if len(sequence) != 0 {
			fmt.Println("Second row done !")
		} else {
			fmt.Println("There was nothing to do")
		}
	}
	return sequence
}

func downToUpCorners(rubik makemove.Rubik, face, target uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if face == target {
		return sequence
	}
	switch diff := (face + 4 - target) % 4; diff {
	case 1:
		sequence = append(sequence, makemove.AllRubikMovesWithName[3+2].Move)
	case 2:
		sequence = append(sequence, makemove.AllRubikMovesWithName[3+1].Move)
	case 3:
		sequence = append(sequence, makemove.AllRubikMovesWithName[3+0].Move)
	}
	return sequence
}

func upToUpCorners(rubik makemove.Rubik, face, target uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if face == target {
		return sequence
	}
	move := makemove.AllRubikMovesWithName[6+(3*face)]
	sequence = append(sequence, move.Move)
	switch diff := (face + 4 - target) % 4; diff {
	case 1:
		sequence = append(sequence, makemove.AllRubikMovesWithName[3+2].Move)
	case 2:
		sequence = append(sequence, makemove.AllRubikMovesWithName[3+1].Move)
	case 3:
		sequence = append(sequence, makemove.AllRubikMovesWithName[3+0].Move)
	}
	sequence = append(sequence, move.Rev)
	return sequence
}

func upCornersOrientation(rubik makemove.Rubik, face uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	move := makemove.AllRubikMovesWithName[6+(3*((face+4-1)%4))]
	sequence = append(sequence, move.Rev)
	sequence = append(sequence, makemove.AllRubikMovesWithName[3].Rev)
	sequence = append(sequence, move.Move)
	sequence = append(sequence, makemove.AllRubikMovesWithName[3].Move)
	return sequence
}

func upCorners(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	reader := bufio.NewReader(os.Stdin)
	if debug {
		fmt.Println("Placing the top corners :")
	}
	for i := uint8(0); i < 4; i++ {
		var seqTmp []makemove.RubikMoves
		index := uint8(getIndex(rubik.PosP3[:], i))
		face := uint8(index % 4)
		fmt.Printf("corner : %d\tindex: %d\tface : %d\n", i, index, face)
		switch floor := uint8(index / 4); floor {
		case 0:
			seqTmp = upToUpCorners(rubik, face, i)
		case 1:
			seqTmp = downToUpCorners(rubik, face, i)
		}
		rubik = rubik.DoMoves(seqTmp)
		sequence = append(sequence, seqTmp...)
		index = uint8(getIndex(rubik.PosP3[:], i))
		face = uint8(index % 4)
		for rubik.RotP3[i] != 0 || rubik.PosP3[i] != i {
			seqTmp = upCornersOrientation(rubik, face)
			fmt.Println(input.SequenceToString(seqTmp))
			rubik = rubik.DoMoves(seqTmp)
			reader.ReadString('\n')
			sequence = append(sequence, seqTmp...)
		}
	}
	if debug {
		if len(sequence) != 0 {
			fmt.Println("Corners done !")
		} else {
			fmt.Println("There was nothing to do")
		}
	}
	return sequence
}

func upToUpCross(rubik makemove.Rubik, target, index, face uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if face == target {
		return sequence
	} else {
		move := makemove.AllRubikMovesWithName[6+(3*face)+1]
		diff := (face + 4 - target) % 4
		new_face := (diff + face) % 4
		fmt.Println(face, diff, new_face)
		sequence = append(sequence, move.Move)
		sequence = append(sequence, makemove.AllRubikMovesWithName[3+diff-1].Move)
		sequence = append(sequence, makemove.AllRubikMovesWithName[6+(3*new_face)+1].Move)
		sequence = append(sequence, move.Rev)
	}
	return sequence
}

func middleToUpCross(rubik makemove.Rubik, target, index, face uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	diff := (face + 4 - target) % 4
	if diff == 3 {
		move := makemove.AllRubikMovesWithName[6+(3*target)+2]
		sequence = append(sequence, move.Move)
	} else if diff == 0 {
		move := makemove.AllRubikMovesWithName[6+(3*target)]
		sequence = append(sequence, move.Move)
	} else {
		move := makemove.AllRubikMovesWithName[6+(3*face)]
		sequence = append(sequence, move.Rev)
		sequence = append(sequence, makemove.AllRubikMovesWithName[3+diff-1].Move)
		var new_face uint8
		switch diff {
		case 1:
			new_face = (face + 4 - 1) % 4
		case 2:
			new_face = (face + 4 - 2) % 4
		case 3:
			new_face = (face + 4 + 1) % 4
		}
		sequence = append(sequence, makemove.AllRubikMovesWithName[6+(3*new_face)+1].Move)
		sequence = append(sequence, move.Move)
	}
	return sequence
}

func downToUpCross(rubik makemove.Rubik, target, index, face uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if face != target {
		switch diff := (face + 4 - target) % 4; diff {
		case 1:
			sequence = append(sequence, makemove.AllRubikMovesWithName[3+0].Move)
		case 2:
			sequence = append(sequence, makemove.AllRubikMovesWithName[3+1].Move)
		case 3:
			sequence = append(sequence, makemove.AllRubikMovesWithName[3+2].Move)
		}
	}
	sequence = append(sequence, makemove.AllRubikMovesWithName[6+(3*target+1)].Move)
	return sequence
}

func switchUpOrientation(rubik makemove.Rubik, target uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	sequence = append(sequence, makemove.AllRubikMovesWithName[6+(3*target)+2].Move)
	sequence = append(sequence, makemove.AllRubikMovesWithName[0].Move)
	sequence = append(sequence, makemove.AllRubikMovesWithName[6+(3*((target+1)%4))+2].Move)
	sequence = append(sequence, makemove.AllRubikMovesWithName[0].Rev)
	return sequence
}

func upCross(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	fmt.Println(rubik)
	if debug {
		fmt.Println("Placing the top cross :")
	}
	for i := uint8(0); i < 4; i++ {
		var seqTmp []makemove.RubikMoves
		index := uint8(getIndex(rubik.PosP2[:], i))
		face := uint8(index % 4)
		switch floor := index / 4; floor {
		case 0:
			seqTmp = upToUpCross(rubik, i, index, face)
		case 1:
			seqTmp = middleToUpCross(rubik, i, index, face)
		case 2:
			seqTmp = downToUpCross(rubik, i, index, face)
		}
		rubik = rubik.DoMoves(seqTmp)
		fmt.Println(input.SequenceToString(seqTmp))
		sequence = append(sequence, seqTmp...)
		if rubik.RotP2[i] == 1 {
			seqTmp = switchUpOrientation(rubik, i)
			fmt.Println(input.SequenceToString(seqTmp))
			sequence = append(sequence, seqTmp...)
			rubik = rubik.DoMoves(seqTmp)
		}
	}
	fmt.Println(rubik)
	if debug {
		if len(sequence) != 0 {
			fmt.Println("Top cross done !")
		} else {
			fmt.Println("There was nothing to do")
		}
	}
	return sequence
}

func MechanicalHuman(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var finalSequence []makemove.RubikMoves
	var tmpSequence []makemove.RubikMoves

	tmpSequence = upCross(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)

	tmpSequence = upCorners(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)

	//	tmpSequence = secondRow(rubik, debug)
	//	finalSequence = append(finalSequence, tmpSequence...)
	//	rubik = rubik.DoMoves(tmpSequence)
	//
	//	tmpSequence = downCross(rubik, debug)
	//	finalSequence = append(finalSequence, tmpSequence...)
	//	rubik = rubik.DoMoves(tmpSequence)
	//
	//	tmpSequence = downCorners(rubik, debug)
	//	finalSequence = append(finalSequence, tmpSequence...)
	//	rubik = rubik.DoMoves(tmpSequence)
	//
	//	tmpSequence = downCornersOrientation(rubik, debug)
	//	finalSequence = append(finalSequence, tmpSequence...)
	//	rubik = rubik.DoMoves(tmpSequence)
	//
	//	tmpSequence = downEdges(rubik, debug)
	//	finalSequence = append(finalSequence, tmpSequence...)
	//	rubik = rubik.DoMoves(tmpSequence)

	return finalSequence
}
