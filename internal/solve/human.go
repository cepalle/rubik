package solve

import (
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
)

func bottomEdges(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
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

func bottomCornersOrientation(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
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

func bottomCorners(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
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

func bottomCross(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
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

func topCorners(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if debug {
		fmt.Println("Placing the top corners :")
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

func topCross(rubik makemove.Rubik, debug bool) []makemove.RubikMoves {
	var sequence []makemove.RubikMoves
	if debug {
		fmt.Println("Placing the top cross :")
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

	tmpSequence = topCross(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)

	tmpSequence = topCorners(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)

	tmpSequence = secondRow(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)

	tmpSequence = bottomCross(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)

	tmpSequence = bottomCorners(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)

	tmpSequence = bottomCornersOrientation(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)

	tmpSequence = bottomEdges(rubik, debug)
	finalSequence = append(finalSequence, tmpSequence...)
	rubik = rubik.DoMoves(tmpSequence)

	return finalSequence
}
