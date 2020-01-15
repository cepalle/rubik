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

func sum(lst []uint8) uint8 {
	var tot = uint8(0)
	for _, value := range lst {
		tot += value
	}
	return tot
}

func getTargetFaceP2(target uint8) uint8 {
	var targetFace = [12]uint8{0, 1, 2, 3, 0, 1, 2, 3, 2, 1, 0, 3}
	return targetFace[target]
}

func toFace(actualFace, targetFace uint8) makemove.RubikMoves {
	switch (targetFace + 4 - actualFace) % 4 {
	case 1:
		return makemove.AllRubikMovesWithName[5].Move
	case 2:
		return makemove.AllRubikMovesWithName[4].Move
	case 3:
		return makemove.AllRubikMovesWithName[3].Move
	}
	return makemove.AllRubikMovesWithName[3].Move
}

func numberGoodCorner(rubik makemove.Rubik) uint8 {
	var tot = uint8(0)
	for i := 20; i < 24; i++ {
		if rubik.PosFP3[i] >= 20 {
			tot += 1
		}
	}
	return tot
}

func rotLst(lst [4]uint8, rot int) [4]uint8 {
	var res [4]uint8
	res[(0+(4-rot))%4] = lst[0]
	res[(1+(4-rot))%4] = lst[1]
	res[(2+(4-rot))%4] = lst[2]
	res[(3+(4-rot))%4] = lst[3]
	fmt.Println("From", lst, "To", res)
	return res
}

func numberGoodOrientedCorner(rubik makemove.Rubik) [4]uint8 {
	var saveSequence [4]uint8
	for j := uint8(0); j < 4; j++ {
		if rubik.PosFP3[13] == 13 && rubik.PosFP3[14] == 14 && rubik.PosFP3[20] == 20 {
			saveSequence[0] = j
		}
		if rubik.PosFP3[15] == 15 && rubik.PosFP3[16] == 16 && rubik.PosFP3[21] == 21 {
			saveSequence[1] = j
		}
		if rubik.PosFP3[17] == 17 && rubik.PosFP3[18] == 18 && rubik.PosFP3[23] == 23 {
			saveSequence[2] = j
		}
		if rubik.PosFP3[12] == 12 && rubik.PosFP3[19] == 19 && rubik.PosFP3[22] == 22 {
			saveSequence[3] = j
		}
		rubik = rubik.DoMove(makemove.AllRubikMovesWithName[3].Move)
	}
	return saveSequence
}

func finalMovesClockwise(face int) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	frontMove := makemove.AllRubikMovesWithName[6+3*(face)+1]
	rightMove := makemove.AllRubikMovesWithName[6+3*((face+3)%4)]
	leftMove := makemove.AllRubikMovesWithName[6+3*((face+1)%4)]
	downMove := makemove.AllRubikMovesWithName[3]
	sequence = append(sequence, frontMove.Move)
	sequence = append(sequence, downMove.Move)
	sequence = append(sequence, rightMove.Move)
	sequence = append(sequence, leftMove.Rev)
	sequence = append(sequence, frontMove.Move)
	sequence = append(sequence, rightMove.Rev)
	sequence = append(sequence, leftMove.Move)
	sequence = append(sequence, downMove.Move)
	sequence = append(sequence, frontMove.Move)
	return sequence
}

func finalMovesCounterClockwise(face int) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	frontMove := makemove.AllRubikMovesWithName[6+3*(face)+1]
	rightMove := makemove.AllRubikMovesWithName[6+3*((face+3)%4)]
	leftMove := makemove.AllRubikMovesWithName[6+3*((face+1)%4)]
	downMove := makemove.AllRubikMovesWithName[3]
	sequence = append(sequence, frontMove.Move)
	sequence = append(sequence, downMove.Rev)
	sequence = append(sequence, rightMove.Move)
	sequence = append(sequence, leftMove.Rev)
	sequence = append(sequence, frontMove.Move)
	sequence = append(sequence, rightMove.Rev)
	sequence = append(sequence, leftMove.Move)
	sequence = append(sequence, downMove.Rev)
	sequence = append(sequence, frontMove.Move)
	return sequence
}

func downEdges(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var seqTmp []makemove.RubikMoves
	fmt.Println(rubik)
	if debug {
		fmt.Print("\nSolving this Cube :")
	}
	var status [4]int
	for i := 8; i < 12; i++ {
		status[i%4] = int(rubik.PosP2[i] % 4)
	}
	var index, tot int
	for i, value := range status {
		if value == i {
			index = i
			tot++
		}
	}
	if tot == 0 {
		sequence = finalMovesClockwise(0)
		rubik = rubik.DoMoves(sequence)
		for i := 8; i < 12; i++ {
			status[i%4] = int(rubik.PosP2[i] % 4)
		}
		for i, value := range status {
			if value == i {
				index = i
				tot++
			}
		}
	}
	if tot != 4 {
		fmt.Println(status, index)
		face := (4 - index) % 4
		if status[face]-face > 0 {
			seqTmp = finalMovesClockwise(face)
			if debug {
				fmt.Printf("Still 3 edges wrongly place, turning clockwise : `%s`\n", input.SequenceToString(seqTmp))
			}
		} else {
			seqTmp = finalMovesCounterClockwise(face)
			if debug {
				fmt.Printf("Still 3 edges wrongly place, turning counter clockwise : `%s`\n", input.SequenceToString(seqTmp))
			}
		}
		sequence = append(sequence, seqTmp...)
	}
	if debug {
		if len(sequence) != 0 {
			fmt.Printf("Cube solved ! Small recap : `%s`\n", input.SequenceToString(sequence))
		} else {
			fmt.Println("There was nothing to do")
		}
	}
	return sequence
}

func downCornersOrientationMove(face uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	sideMove := makemove.AllRubikMovesWithName[6+3*((face+1)%4)]
	frontMove := makemove.AllRubikMovesWithName[6+3*(face)]
	backMove := makemove.AllRubikMovesWithName[6+3*((face+2)%4)+1].Move
	sequence = append(sequence, sideMove.Rev)
	sequence = append(sequence, frontMove.Move)
	sequence = append(sequence, sideMove.Rev)
	sequence = append(sequence, backMove)
	sequence = append(sequence, sideMove.Move)
	sequence = append(sequence, frontMove.Rev)
	sequence = append(sequence, sideMove.Rev)
	sequence = append(sequence, backMove)
	sequence = append(sequence, makemove.AllRubikMovesWithName[6+3*((face+1)%4)+1].Move)
	sequence = append(sequence, makemove.AllRubikMovesWithName[5].Move)
	return sequence
}

func downCornersOrientation(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var seqTmp []makemove.RubikMoves
	if debug {
		fmt.Printf("\nChanging corners orientation : ")
	}
	saveState := numberGoodOrientedCorner(rubik)
	if sum(saveState[:]) == 0 {
	} else if saveState[0] == saveState[2] || saveState[1] == saveState[3] {
		seqTmp = downCornersOrientationMove(0)
		rubik = rubik.DoMoves(seqTmp)
		sequence = append(sequence, seqTmp...)
		if debug {
			fmt.Printf("Corners are forming a diagonal : `%s`\n", input.SequenceToString(seqTmp))
		}
	}
	saveState = numberGoodOrientedCorner(rubik)
	var rot [4]uint8
	for _, value := range saveState {
		rot[value]++
	}
	face := getIndex(rot[:], 2)
	if face != -1 {
		if face != 0 {
			rubik = rubik.DoMove(makemove.AllRubikMovesWithName[3+face-1].Move)
			sequence = append(sequence, makemove.AllRubikMovesWithName[3+face-1].Move)
		}
		saveState = numberGoodOrientedCorner(rubik)
		face = getIndex(saveState[:], 0)
		if face == 0 && saveState[3] == 0 {
			face = 3
		}
		seqTmp = downCornersOrientationMove(uint8(4-face) % 4)
		rubik = rubik.DoMoves(seqTmp)
		sequence = append(sequence, seqTmp...)
		if debug {
			fmt.Printf("Corners are next to each other : `%s`\n", input.SequenceToString(seqTmp))
		}
	} else if saveState[0] != 0 {
		rubik = rubik.DoMove(makemove.AllRubikMovesWithName[3+saveState[0]-1].Move)
		sequence = append(sequence, makemove.AllRubikMovesWithName[3+saveState[0]-1].Move)
	}
	if debug {
		if len(sequence) != 0 {
			fmt.Printf("The bottom edges should be in a good orientation now !\nSmall recap : `%s`\n", input.SequenceToString(sequence))
		} else {
			fmt.Println("There was nothing to do")
		}
	}
	return sequence
}

func downCornerMoves(face uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	sideMove := makemove.AllRubikMovesWithName[6+3*((face+1)%4)]
	downMove := makemove.AllRubikMovesWithName[3]
	sequence = append(sequence, sideMove.Move)
	sequence = append(sequence, downMove.Move)
	sequence = append(sequence, sideMove.Rev)
	sequence = append(sequence, downMove.Move)
	sequence = append(sequence, sideMove.Move)
	sequence = append(sequence, makemove.AllRubikMovesWithName[4].Move)
	sequence = append(sequence, sideMove.Rev)
	return sequence
}

func zeroCornerUp(rubik makemove.Rubik) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var index = [4]uint8{18, 16, 14, 12}
	var actualFace uint8
	for face, i := range index {
		if rubik.PosFP3[i] > 19 {
			actualFace = uint8((face + 1) % 4)
			break
		}
	}
	sequence = downCornerMoves(actualFace)
	return sequence
}

func oneCornerUp(rubik makemove.Rubik) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var index = [4]uint8{22, 23, 21, 20}
	var actualFace uint8
	for face, i := range index {
		if rubik.PosFP3[i] > 19 {
			actualFace = uint8(face)
			break
		}
	}
	sequence = downCornerMoves(actualFace)
	return sequence
}

func twoCornerUp(rubik makemove.Rubik) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var index = [4]uint8{19, 17, 15, 13}
	var actualFace uint8
	for face, i := range index {
		if rubik.PosFP3[i] > 19 {
			actualFace = uint8(face)
			break
		}
	}
	sequence = downCornerMoves(actualFace)
	return sequence
}

func downCorners(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var seqTmp []makemove.RubikMoves
	if debug {
		fmt.Println("\nPlacing bottom corner at their respectives places, not regarding the orientation :")
	}
	a := numberGoodCorner(rubik)
	for a != 4 {
		switch a {
		case 0:
			seqTmp = zeroCornerUp(rubik)
		case 1:
			seqTmp = oneCornerUp(rubik)
		case 2:
			seqTmp = twoCornerUp(rubik)
		}
		rubik = rubik.DoMoves(seqTmp)
		sequence = append(sequence, seqTmp...)
		a = numberGoodCorner(rubik)
	}
	if debug {
		if len(sequence) != 0 {
			fmt.Printf("Corners placed ! Small recap : `%s`\n", input.SequenceToString(sequence))
		} else {
			fmt.Println("There was nothing to do")
		}
	}
	return sequence
}

func downCrossByQuarter(odd uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var vertMove makemove.RubikMovesWithName
	var frontMove makemove.RubikMovesWithName
	var downMove makemove.RubikMovesWithName
	downMove = makemove.AllRubikMovesWithName[3]
	frontMove = makemove.AllRubikMovesWithName[6+3*((odd)%4)]
	vertMove = makemove.AllRubikMovesWithName[6+3*((odd+1)%4)]
	sequence = append(sequence, frontMove.Move)
	sequence = append(sequence, downMove.Move)
	sequence = append(sequence, vertMove.Move)
	sequence = append(sequence, downMove.Rev)
	sequence = append(sequence, vertMove.Rev)
	sequence = append(sequence, frontMove.Rev)
	return sequence
}

func downCrossByLine(odd uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	var vertMove makemove.RubikMovesWithName
	var frontMove makemove.RubikMovesWithName
	var downMove makemove.RubikMovesWithName
	downMove = makemove.AllRubikMovesWithName[3]
	switch odd {
	case 0:
		vertMove = makemove.AllRubikMovesWithName[15]
		frontMove = makemove.AllRubikMovesWithName[12]
	case 1:
		vertMove = makemove.AllRubikMovesWithName[6]
		frontMove = makemove.AllRubikMovesWithName[15]
	}
	sequence = append(sequence, frontMove.Move)
	sequence = append(sequence, vertMove.Move)
	sequence = append(sequence, downMove.Move)
	sequence = append(sequence, vertMove.Rev)
	sequence = append(sequence, downMove.Rev)
	sequence = append(sequence, frontMove.Rev)
	return sequence
}

func downCross(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if debug {
		fmt.Print("\nCreating the bottom cross :")
	}
	if debug {
		var res = [4]uint8{2, 2, 2, 2}
		for i, pos := range rubik.PosP2 {
			if i < 8 {
				continue
			}
			res[i%4] = (pos + uint8(i) + rubik.RotP2[pos]) % 2
		}
		evenLine := res[0] + res[2]
		oddLine := res[1] + res[3]
		if oddLine+evenLine >= 3 {
			seqTmp := downCrossByQuarter(0)
			if debug {
				fmt.Printf(" found no pattern. `%s`\n", input.SequenceToString(seqTmp))
			}
			sequence = append(sequence, seqTmp...)
			rubik = rubik.DoMoves(seqTmp)
			for i, pos := range rubik.PosP2 {
				if i < 8 {
					continue
				}
				res[i%4] = (pos + uint8(i) + rubik.RotP2[pos]) % 2
			}
			evenLine = res[0] + res[2]
			oddLine = res[1] + res[3]
		}
		if evenLine+oddLine == 0 {
			if debug {
				fmt.Printf(" cross already done\n")
			}
		} else if evenLine == 0 {
			seqTmp := downCrossByLine(1)
			if debug {
				fmt.Printf(" found a straight line. `%s`\n", input.SequenceToString(seqTmp))
			}
			sequence = append(sequence, seqTmp...)
		} else if oddLine == 0 {
			seqTmp := downCrossByLine(0)
			if debug {
				fmt.Printf(" found a straight line. `%s`\n", input.SequenceToString(seqTmp))
			}
			sequence = append(sequence, seqTmp...)
		} else {
			var seqTmp []makemove.RubikMoves
			index := uint8(getIndex(res[:], 0))
			if index == 0 && res[3] == 0 {
				seqTmp = downCrossByQuarter(0)
			} else if index == 0 {
				seqTmp = downCrossByQuarter(3)
			} else if index == 1 {
				seqTmp = downCrossByQuarter(2)
			} else if index == 2 {
				seqTmp = downCrossByQuarter(1)
			}
			if debug {
				fmt.Printf(" found comma. `%s`\n", input.SequenceToString(seqTmp))
			}
			sequence = append(sequence, seqTmp...)
		}

		if debug {
			if len(sequence) != 0 {
				fmt.Printf("Bottom cross placed ! Small recap : `%s`\n", input.SequenceToString(sequence))
			} else {
				fmt.Println("There was nothing to do")
			}
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
	actualFace := getTargetFaceP2(index)
	targetFace := getTargetFaceP2(target)
	if actualFace != targetFace {
		sequence = append(sequence, toFace(actualFace, targetFace))
	}
	rubik = rubik.DoMoves(sequence)
	if rubik.RotP2[target] == 1 {
		sequence = append(sequence, leftInsertion(getTargetFaceP2(targetFace))...)
	} else {
		sequence = append(sequence, makemove.AllRubikMovesWithName[5].Move)
		sequence = append(sequence, rightInsertion(getTargetFaceP2((targetFace+1)%4))...)
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
			fmt.Printf("Second row done ! Small recap : `%s`\n", input.SequenceToString(sequence))
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

func upToUpCross(actualFace, targetFace uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if actualFace == targetFace {
		return sequence
	} else {
		sequence = append(sequence, makemove.AllRubikMovesWithName[6+(3*actualFace)+1].Move)
		sequence = append(sequence, toFace(actualFace, targetFace))
		sequence = append(sequence, makemove.AllRubikMovesWithName[6+(3*targetFace)+1].Move)
	}
	return sequence
}

func middleToUpCross(actualFace, targetFace uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	diff := (actualFace + 4 - targetFace) % 4
	if diff == 3 {
		move := makemove.AllRubikMovesWithName[6+(3*targetFace)+2]
		sequence = append(sequence, move.Move)
	} else if diff == 0 {
		move := makemove.AllRubikMovesWithName[6+(3*targetFace)]
		sequence = append(sequence, move.Move)
	} else {
		move := makemove.AllRubikMovesWithName[6+(3*actualFace)]
		sequence = append(sequence, move.Rev)
		sequence = append(sequence, makemove.AllRubikMovesWithName[3+diff-1].Move)
		sequence = append(sequence, makemove.AllRubikMovesWithName[6+(3*targetFace)+1].Move)
		sequence = append(sequence, move.Move)
	}
	return sequence
}

func downToUpCross(actualFace, targetFace uint8) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if actualFace != targetFace {
		sequence = append(sequence, toFace(actualFace, targetFace))
	}
	sequence = append(sequence, makemove.AllRubikMovesWithName[6+(3*targetFace+1)].Move)
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
		actualFace := getTargetFaceP2(index)
		switch floor := index / 4; floor {
		case 0:
			seqTmp = upToUpCross(actualFace, i)
		case 1:
			seqTmp = middleToUpCross(actualFace, i)
		case 2:
			seqTmp = downToUpCross(actualFace, i)
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
	var sum = uint8(0)
	if checkSecondRow(rubik) == false {
		return false
	}
	for i, pos := range rubik.PosP2 {
		if i < 8 {
			continue
		}
		sum += pos + uint8(i) + rubik.RotP2[pos]
	}
	if sum%2 == 1 {
		fmt.Fprintf(os.Stderr, "Down cross failed\n")
		return false
	}
	return true
}

func checkDownCorners(rubik makemove.Rubik) bool {
	if checkDownCross(rubik) == false {
		return false
	}
	if numberGoodCorner(rubik) != 4 {
		fmt.Fprintf(os.Stderr, "Down corners failed\n")
		return false
	}
	return true
}

func checkDownFace(rubik makemove.Rubik) bool {
	if checkDownCorners(rubik) == false {
		return false
	}
	for i := uint8(12); i < 24; i++ {
		if rubik.PosFP3[i] != i {
			fmt.Fprintf(os.Stderr, "Down face failed\n")
			return false
		}
	}
	return true
}

func checkRubik(rubik makemove.Rubik) bool {
	for i := uint8(0); i < 24; i++ {
		if rubik.PosFP3[i] != i {
			fmt.Fprintf(os.Stderr, "Down orientation failed\n")
			return false
		}
	}
	for i := uint8(0); i < 12; i++ {
		if rubik.PosP2[i] != i || rubik.RotP2[i] == 1 {
			fmt.Fprintf(os.Stderr, "Down orientation failed\n")
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
	//correct = checkTopCross(rubik)
	//if !correct {
	//	return nil
	//} else if debug {
	//	fmt.Println("Check on up cross done.")
	//}

	tmpSequence = upCorners(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)
	//correct = checkTopCorners(rubik)
	//if !correct {
	//	return nil
	//} else if debug {
	//	fmt.Println("Check on up corners done.")
	//}

	tmpSequence = secondRow(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)
	//correct = checkSecondRow(rubik)
	//if !correct {
	//	return nil
	//} else if debug {
	//	fmt.Println("Check on second row done.")
	//}

	tmpSequence = downCross(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)
	//correct = checkDownCross(rubik)
	//if !correct {
	//	return nil
	//} else if debug {
	//	fmt.Println("Check on down cross done.")
	//}

	tmpSequence = downCorners(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)
	//correct = checkDownCorners(rubik)
	//if !correct {
	//	return nil
	//} else if debug {
	//	fmt.Println("Check on down cross done.")
	//}

	tmpSequence = downCornersOrientation(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)
	//correct = checkDownFace(rubik)
	//if !correct {
	//	fmt.Println(rubik)
	//	return nil
	//} else if debug {
	//	fmt.Println("Check on down face done.")
	//}

	tmpSequence = downEdges(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)
	correct = checkRubik(rubik)
	if !correct {
		fmt.Println(rubik)
		return nil
	} else if debug {
		fmt.Println("Check on down face done.")
	}

	return finalSequence
}
