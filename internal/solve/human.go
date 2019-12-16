package solve

import (
	"fmt"
	"github.com/cepalle/rubik/internal/input"
	"github.com/cepalle/rubik/internal/makemove"
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
		fmt.Println("Solving this Cube :")
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
	var targetFaces = [8]uint8{4, 5, 5, 2, 2, 3, 3, 4}
	if debug {
		fmt.Println("Setting up the second row :")
	}
	for target := uint8(4); target < 8; target++ {
		index := uint8(getIndex(rubik.PosP2[:], target))
		targetFace := targetFaces[target%4+rubik.RotP2[target]]
		if index == target {
			continue
		}
		fmt.Printf("target : %d, index :%d, face : %d targetFace : %d\n", target, index, index%4+2, targetFace)

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

func upCornersOrientation(rubik makemove.Rubik, corner uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var move makemove.RubikMovesWithName
	switch corner {
	case 0:
		move = makemove.AllRubikMovesWithName[15]
	case 1:
		move = makemove.AllRubikMovesWithName[6]
	case 2:
		move = makemove.AllRubikMovesWithName[12]
	case 3:
		move = makemove.AllRubikMovesWithName[9]
	}
	sequence = append(sequence, move.Rev)
	sequence = append(sequence, makemove.AllRubikMovesWithName[3].Rev)
	sequence = append(sequence, move.Move)
	sequence = append(sequence, makemove.AllRubikMovesWithName[3].Move)
	return sequence
}

func upCorners(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var faces = [24]uint8{5, 5, 5, 5, 3, 3, 2, 2, 1, 1, 0, 0, 3, 3, 2, 2, 1, 1, 0, 0, 4, 4, 4, 4}
	var corners = [24]uint8{0, 1, 2, 3, 0, 2, 2, 3, 3, 1, 1, 0, 0, 2, 2, 3, 3, 1, 1, 0, 2, 3, 0, 1}
	var rots = [4][4]uint8{{0, 0, 2, 1}, {2, 0, 1, 0}, {0, 1, 0, 2}, {1, 2, 0, 0}}
	if debug {
		fmt.Println("Placing the top corners :")
	}
	for i := uint8(0); i < 4; i++ {
		var seqTmp []makemove.RubikMoves
		index := uint8(getIndex(rubik.PosFP3[:], i))
		if index == i {
			continue
		}
		face := faces[index]
		corner := corners[index]
		if corner != i && index < 12 {
			seqTmp = append(seqTmp, makemove.AllRubikMovesWithName[6+(3*face)].Move)
			seqTmp = append(seqTmp, makemove.AllRubikMovesWithName[3+rots[corner][i]].Move)
			seqTmp = append(seqTmp, makemove.AllRubikMovesWithName[6+(3*face)].Rev)
		} else if corner != i && index > 11 {
			seqTmp = append(seqTmp, makemove.AllRubikMovesWithName[3+rots[corner][i]].Move)
		}
		rubik = rubik.DoMoves(seqTmp)
		sequence = append(sequence, seqTmp...)
		index = uint8(getIndex(rubik.PosFP3[:], i))
		face = faces[index]
		for rubik.PosFP3[i] != i {
			seqTmp = upCornersOrientation(rubik, i)
			fmt.Println(input.SequenceToString(seqTmp))
			rubik = rubik.DoMoves(seqTmp)
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
			sequence = append(sequence, makemove.AllRubikMovesWithName[3+2].Move)
		case 2:
			sequence = append(sequence, makemove.AllRubikMovesWithName[3+1].Move)
		case 3:
			sequence = append(sequence, makemove.AllRubikMovesWithName[3+0].Move)
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
		if len(seqTmp) != 0 {
			fmt.Println(input.SequenceToString(seqTmp))
		}
		sequence = append(sequence, seqTmp...)
		if rubik.RotP2[i] == 1 {
			seqTmp = switchUpOrientation(rubik, i)
			if len(seqTmp) == 0 {
				fmt.Println(input.SequenceToString(seqTmp))
			}
			sequence = append(sequence, seqTmp...)
			rubik = rubik.DoMoves(seqTmp)
		}
	}
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

	fmt.Println(rubik)
	tmpSequence = secondRow(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)

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
