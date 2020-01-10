package solve

import (
	"fmt"
	"github.com/cepalle/rubik/internal/input"
	"github.com/cepalle/rubik/internal/makemove"
)

func dlsPredicate(r *makemove.Rubik, depth uint32, predicate func(*makemove.Rubik) bool) []makemove.RubikMoves {
	var res []makemove.RubikMoves

	if depth == 0 {
		if predicate(r) {
			return []makemove.RubikMoves{}
		}
		return nil
	}

	for _, m := range makemove.AllRubikMovesWithName {
		res = dlsPredicate(r.DoMovePtr(m.Move), depth-1, predicate)
		if res != nil {
			return append(res, m.Move)
		}
		r.DoMovePtr(m.Rev)
	}
	return nil
}

// /!\ update the cube with the state that has match the predicate
func IddfsPredicate(r *makemove.Rubik, predicate func(*makemove.Rubik) bool) []makemove.RubikMoves {
	var res []makemove.RubikMoves

	for i := uint32(0); ; i++ {
		res = dlsPredicate(r, i, predicate)
		if res != nil {
			return res
		}
	}
}

// /!\ update the cube with the state that has match the predicate
func IddfsPredicateImax(r *makemove.Rubik, predicate func(*makemove.Rubik) bool, imax uint32) []makemove.RubikMoves {
	var res []makemove.RubikMoves

	for i := uint32(0); i < imax; i++ {
		res = dlsPredicate(r, i, predicate)
		if res != nil {
			return res
		}
	}
	return nil
}

var rubikGoalIt = makemove.Rubik{
	PosP2:  [12]uint8{0, 2, 1, 3, 12, 13, 14, 15, 28, 30, 29, 31},
	RotP2:  [12]uint8{4, 6, 5, 7, 16, 17, 18, 19, 32, 33, 34, 35},
	PosFP3: [24]uint8{8, 9, 10, 11, 20, 21, 22, 23, 24, 25, 26, 27, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47},
}

func IddfsIt(r makemove.Rubik) []makemove.RubikMoves {
	var res []makemove.RubikMoves

	for nb_it := uint8(0); nb_it < 48; nb_it++ {
		fmt.Println(nb_it)
		move_it := IddfsPredicate(&r, func(cube *makemove.Rubik) bool {
			var i uint8
			for i = 0; i < 12; i++ {
				if rubikGoalIt.RotP2[i] <= nb_it && cube.RotP2[i] != 0 {
					return false
				}
			}
			for i = 0; i < 12; i++ {
				if rubikGoalIt.RotP2[i] <= nb_it && cube.PosP2[i] != i {
					return false
				}
			}
			for i = 0; i < 24; i++ {
				if rubikGoalIt.PosFP3[i] <= nb_it && cube.PosFP3[i] != i {
					return false
				}
			}
			return true
		})
		for _, e := range move_it {
			res = append(res, e)
		}
	}
	return input.ReverseMove(res)
}

func IddfsItHamming(r makemove.Rubik) []makemove.RubikMoves {
	var res []makemove.RubikMoves

	for nb_it := uint8(0); nb_it < 48; nb_it++ {
		fmt.Println(nb_it)
		move_it := IddfsPredicate(&r, func(cube *makemove.Rubik) bool {
			return float64(nb_it) < (48 - ScoringHamming(cube))
		})
		for _, e := range move_it {
			res = append(res, e)
		}
	}
	return input.ReverseMove(res)
}
