package solve

import (
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

func rightInsertion(face uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var vertMove makemove.RubikMovesWithName
	var frontMove makemove.RubikMovesWithName
	var downMove makemove.RubikMovesWithName
	downMove = makemove.AllRubikMovesWithName[3]
	switch face {
	case 0:
		vertMove = makemove.AllRubikMovesWithName[15]
		frontMove = makemove.AllRubikMovesWithName[6]
	case 1:
		vertMove = makemove.AllRubikMovesWithName[6]
		frontMove = makemove.AllRubikMovesWithName[9]
	case 2:
		vertMove = makemove.AllRubikMovesWithName[9]
		frontMove = makemove.AllRubikMovesWithName[12]
	case 3:
		vertMove = makemove.AllRubikMovesWithName[12]
		frontMove = makemove.AllRubikMovesWithName[15]
	}
	sequence = append(sequence, downMove.Rev)
	sequence = append(sequence, vertMove.Rev)
	sequence = append(sequence, downMove.Move)
	sequence = append(sequence, vertMove.Move)
	sequence = append(sequence, downMove.Move)
	sequence = append(sequence, frontMove.Move)
	sequence = append(sequence, downMove.Rev)
	sequence = append(sequence, frontMove.Rev)
	return sequence
}

func leftInsertion(face uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var vertMove makemove.RubikMovesWithName
	var frontMove makemove.RubikMovesWithName
	var downMove makemove.RubikMovesWithName
	downMove = makemove.AllRubikMovesWithName[3]
	switch face {
	case 0:
		vertMove = makemove.AllRubikMovesWithName[9]
		frontMove = makemove.AllRubikMovesWithName[6]
	case 1:
		vertMove = makemove.AllRubikMovesWithName[12]
		frontMove = makemove.AllRubikMovesWithName[9]
	case 2:
		vertMove = makemove.AllRubikMovesWithName[15]
		frontMove = makemove.AllRubikMovesWithName[12]
	case 3:
		vertMove = makemove.AllRubikMovesWithName[6]
		frontMove = makemove.AllRubikMovesWithName[15]
	}
	sequence = append(sequence, downMove.Move)
	sequence = append(sequence, vertMove.Move)
	sequence = append(sequence, downMove.Rev)
	sequence = append(sequence, vertMove.Rev)
	sequence = append(sequence, downMove.Rev)
	sequence = append(sequence, frontMove.Rev)
	sequence = append(sequence, downMove.Move)
	sequence = append(sequence, frontMove.Move)
	return sequence
}

func secondRowExtraction(face uint8) []makemove.RubikMoves {
	return leftInsertion(face)
}

func secondRowEdgeWrongOrientation(face uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	sequence = leftInsertion(face)
	sequence = append(sequence, makemove.AllRubikMovesWithName[4].Move)
	sequence = append(sequence, leftInsertion(face)...)
	return sequence
}

func insertSecondRow(rubik makemove.Rubik, target, index, rot uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	face := index % 4
	targetFace := target % 4
	rotMove := (targetFace + 4 - face) % 4
	newRot := (rot + (rotMove+1)%2) % 2
	switch rotMove = (rotMove + newRot) % 4; rotMove {
	case 1:
		sequence = append(sequence, makemove.AllRubikMovesWithName[5].Move)
	case 2:
		sequence = append(sequence, makemove.AllRubikMovesWithName[4].Move)
	case 3:
		sequence = append(sequence, makemove.AllRubikMovesWithName[3].Move)
	}
	rubik = rubik.DoMoves(sequence)
	newIndex := uint8(getIndex(rubik.PosP2[:], target))
	if target%2 == newIndex%2 {
		sequence = append(sequence, leftInsertion(newIndex%4)...)
	} else {
		sequence = append(sequence, rightInsertion(newIndex%4)...)
	}
	return sequence
}

func secondRow(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var seqTmp []makemove.RubikMoves
	if debug {
		fmt.Println("\nSetting up the second row :")
	}
	for target := uint8(4); target < 8; target++ {
		if debug {
			fmt.Printf("Edge number %d : ", target)
		}
		index := uint8(getIndex(rubik.PosP2[:], target))
		if index == target && rubik.RotP2[target] == 1 {
			seqTmp = secondRowEdgeWrongOrientation(target % 4)
			if debug {
				fmt.Printf("wrong orientation, `%s`\n", input.SequenceToString(seqTmp))
			}
			sequence = append(sequence, seqTmp...)
			rubik = rubik.DoMoves(seqTmp)
			continue
		} else if index == target {
			if debug {
				fmt.Println("already placed")
			}
			continue
		}
		if index/4 == 1 {
			seqTmp = secondRowExtraction(index % 4)
			if debug {
				fmt.Printf("need to remove from second layer `%s`.\nThen place it right ", input.SequenceToString(seqTmp))
			}
			sequence = append(sequence, seqTmp...)
			rubik = rubik.DoMoves(seqTmp)
			index = uint8(getIndex(rubik.PosP2[:], target))
		}
		seqTmp = insertSecondRow(rubik, target, index, rubik.RotP2[target])
		if debug {
			if len(seqTmp) != 0 {
				fmt.Print("`", input.SequenceToString(seqTmp), "`\n")
			} else {
				fmt.Print("The corner is already at the right face\n")
			}
		}
		sequence = append(sequence, seqTmp...)
		rubik = rubik.DoMoves(seqTmp)
	}
	if debug {
		if len(sequence) != 0 {
			fmt.Println("Second row done ! Small recap :", input.SequenceToString(sequence))
		} else {
			fmt.Println("There was nothing to do")
		}
	}
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
	var faces = [24]uint8{3, 0, 2, 1, 3, 2, 2, 1, 1, 0, 0, 3, 3, 2, 2, 1, 1, 0, 0, 3, 2, 1, 3, 0}
	var targetFace = [4]uint8{3, 0, 2, 1}
	var rots = [4][4]uint8{{3, 2, 1, 0}, {0, 3, 2, 1}, {1, 0, 3, 2}, {2, 1, 0, 3}}
	var repeat uint8
	if debug {
		fmt.Println("\nPlacing the top corners :")
	}
	for i := uint8(0); i < 4; i++ {
		if debug {
			fmt.Printf("Up corner %d ", i)
		}
		var seqTmp []makemove.RubikMoves
		index := uint8(getIndex(rubik.PosFP3[:], i))
		if index == i {
			if debug {
				fmt.Println(" : this corner is already placed")
			}
			continue
		}
		face := faces[index]
		targetMove := rots[face][targetFace[i]]
		if targetMove == 3 {
		} else if index < 12 {
			seqTmp = append(seqTmp, makemove.AllRubikMovesWithName[6+(3*face)].Rev)
			seqTmp = append(seqTmp, makemove.AllRubikMovesWithName[3+targetMove].Move)
			seqTmp = append(seqTmp, makemove.AllRubikMovesWithName[6+(3*face)].Move)
			if targetMove == 0 {
				seqTmp = append(seqTmp, makemove.AllRubikMovesWithName[3].Move)
			}
		} else {
			seqTmp = append(seqTmp, makemove.AllRubikMovesWithName[3+targetMove].Move)
		}
		rubik = rubik.DoMoves(seqTmp)
		if debug {
			if len(seqTmp) != 0 {
				fmt.Print("`", input.SequenceToString(seqTmp), "`\n")
			} else {
				fmt.Print("The corner is already at the right face\n")
			}
		}
		sequence = append(sequence, seqTmp...)
		index = uint8(getIndex(rubik.PosFP3[:], i))
		face = faces[index]
		repeat = 0
		for rubik.PosFP3[i] != i {
			repeat += 1
			seqTmp = upCornersOrientation(rubik, i)
			if debug && repeat == 1 {
				fmt.Print("Repeat `", input.SequenceToString(seqTmp))
			}
			rubik = rubik.DoMoves(seqTmp)
			sequence = append(sequence, seqTmp...)
		}
		if debug && repeat != 0 {
			fmt.Printf("` %d times\n", repeat)
		}
	}
	if debug {
		if len(sequence) != 0 {
			fmt.Print("Up corners done ! Small recap : `", input.SequenceToString(sequence), "`\n")
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
		var new_face uint8
		switch diff {
		case 1:
			new_face = (face + 4 - 1) % 4
		case 2:
			new_face = (face + 4 - 2) % 4
		case 3:
			new_face = (face + 4 + 1) % 4
		}
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
	if debug {
		fmt.Println("\nPlacing the up cross :")
	}
	for i := uint8(0); i < 4; i++ {
		if debug {
			fmt.Printf("Up edge number %d `", i)
		}
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
		if debug {
			if len(seqTmp) != 0 {
				fmt.Print(input.SequenceToString(seqTmp))
			} else {
				fmt.Print("Nothing to do for this edge")
			}
			fmt.Printf("` and done !\n")
		}
		sequence = append(sequence, seqTmp...)
		if rubik.RotP2[i] == 1 {
			if debug {
				fmt.Printf("Changing orientation : `")
			}
			seqTmp = switchUpOrientation(rubik, i)
			if debug {
				fmt.Print(input.SequenceToString(seqTmp), "`\n")
			}
			sequence = append(sequence, seqTmp...)
			rubik = rubik.DoMoves(seqTmp)
		}
	}
	if debug {
		if len(sequence) != 0 {
			fmt.Print("Up cross done ! Small recap : `", input.SequenceToString(sequence), "`\n")
		} else {
			fmt.Println("There was nothing to do")
		}
	}
	return sequence
}

func checkTopCross(rubik makemove.Rubik) bool {
	for i, pos := range rubik.PosP2 {
		if i > 3 {
			break
		}
		if uint8(i) != pos || rubik.RotP2[i] != 0 {
			fmt.Fprintf(os.Stderr, "Up cross failed\n")
			return false
		}
	}
	return true
}

func checkTopCorners(rubik makemove.Rubik) bool {
	if checkTopCross(rubik) == false {
		return false
	}
	for i, pos := range rubik.PosFP3 {
		if i > 11 {
			break
		}
		if uint8(i) != pos {
			fmt.Println(i, pos)
			fmt.Fprintf(os.Stderr, "Up corners failed\n")
			return false
		}
	}
	return true
}

func checkSecondRow(rubik makemove.Rubik) bool {
	if checkTopCorners(rubik) == false {
		return false
	}
	for i, pos := range rubik.PosP2 {
		if i > 7 {
			break
		}
		if uint8(i) != pos || rubik.RotP2[i] != 0 {
			fmt.Fprintf(os.Stderr, "Second row failed\n")
			return false
		}
	}
	return true
}

func checkDownCross(rubik makemove.Rubik) bool {
	if checkSecondRow(rubik) == false {
		return false
	}
	for i, pos := range rubik.PosP2 {
		if uint8(i) != pos || rubik.RotP2[i] != 0 {
			fmt.Fprintf(os.Stderr, "Top cross failed\n")
			return false
		}
	}
	return true
}

func MechanicalHuman(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var finalSequence []makemove.RubikMoves
	var tmpSequence []makemove.RubikMoves
	var correct bool

	tmpSequence = upCross(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)

	tmpSequence = upCorners(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)

	tmpSequence = secondRow(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)
	correct = checkSecondRow(rubik)
	if !correct {
		return nil
	} else if debug {
		fmt.Println("Check on second row done.")
	}

	tmpSequence = downCross(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)
	correct = checkDownCross(rubik)
	if !correct {
		return nil
	} else {
		fmt.Println("Check on down cross done.")
	}

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
