package solve

import (
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
)

/*
	Algo: http://www.stefan-pochmann.info/spocc/other_stuff/tools/solver_thistlethwaite/solver_thistlethwaite_cpp.txt

	turns = move % 3;
	face = move / 3;
*/

var affectedCubies = [6][8]uint8{
	{0, 1, 2, 3, 0, 1, 2, 3},   // U
	{4, 7, 6, 5, 4, 5, 6, 7},   // D
	{0, 9, 4, 8, 0, 3, 5, 4},   // F
	{2, 10, 6, 11, 2, 1, 7, 6}, // B
	{3, 11, 7, 9, 3, 2, 6, 5},  // L
	{1, 8, 5, 10, 1, 0, 4, 7},  // R
}

type cube struct {
	PosP2 [12]uint8
	RotP2 [12]uint8

	PosP3 [8]uint8
	RotP3 [8]uint8
}

var goalCube = cube{
	[12]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	[12]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[8]uint8{0, 1, 2, 3, 4, 5, 6, 7},
	[8]uint8{0, 0, 0, 0, 0, 0, 0, 0},
}

func doMove(c cube, move uint8) cube {
	// TODO
}

func doMoves(c cube, moves []uint8) cube {
	res := c

	for _, e := range moves {
		res = doMove(res, e)
	}

	return res
}

func appendMovesReversed(a []uint8, b []uint8) []uint8 {
	aCp := make([]uint8, len(a))
	copy(aCp, a)

	for i := len(b); i > 0; i-- {
		aCp = append(aCp, b[i-1])
	}

	return aCp
}

func bidirectionalBfs(src cube, dst cube, id func(c cube) cube, dir []uint8) []uint8 {

	type node struct {
		cube  cube
		moves []uint8
	}

	hysSrc := make(map[cube][]uint8)
	var pileSrc []node

	pileSrc = append(pileSrc, node{src, []uint8{}})
	hysSrc[id(src)] = []uint8{}

	hysDst := make(map[cube][]uint8)
	var pileDst []node

	pileDst = append(pileDst, node{dst, []uint8{}})
	hysDst[id(dst)] = []uint8{}

	for len(pileSrc) > 0 || len(pileDst) > 0 {

		if len(pileSrc) > 0 {
			// SRC
			curSrc := pileSrc[0]
			pileSrc = pileSrc[1:]

			movesDst, found := hysDst[id(curSrc.cube)]
			if found {
				fmt.Println(curSrc)
				return appendMovesReversed(curSrc.moves, movesDst)
			}

			for _, d := range dir {
				var nwCube = doMove(curSrc.cube, d)

				idCubeSrc := id(nwCube)
				_, found := hysSrc[idCubeSrc]
				if found {
					continue
				}

				mvsCp := make([]uint8, len(curSrc.moves))
				copy(mvsCp, curSrc.moves)
				nwMvs := append(mvsCp, d)
				var nNode = node{
					nwCube,
					nwMvs,
				}

				hysSrc[idCubeSrc] = nwMvs
				pileSrc = append(pileSrc, nNode)
			}
		}

		if len(pileDst) > 0 {
			// DST
			curDst := pileDst[0]
			pileDst = pileDst[1:]

			movesSrc, found := hysSrc[id(curDst.cube)]
			if found {
				fmt.Println(curDst)
				return appendMovesReversed(movesSrc, curDst.moves)
			}

			for _, d := range dir {
				var nwCube = doMove(curDst.cube, d)

				idCubeDst := id(nwCube)
				_, found := hysSrc[idCubeDst]
				if found {
					continue
				}

				mvsCp := make([]uint8, len(curDst.moves))
				copy(mvsCp, curDst.moves)
				nwMvs := append(mvsCp, d)
				var nNode = node{
					nwCube,
					nwMvs,
				}

				hysDst[idCubeDst] = nwMvs
				pileDst = append(pileDst, nNode)
			}
		}

	}
	return nil
}

func idG0(c cube) cube {
	// TODO
}

func g0(c cube) []uint8 {
	var dirG0 = []uint8{
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 10, 11,
		12, 13, 14,
		15, 16, 17,
	}

	return bidirectionalBfs(c, goalCube, idG0, dirG0)
}

func idG1(c cube) cube {
	// TODO
}

func g1(c cube) []uint8 {
	var dirG1 = []uint8{
		17, 16, 15,
		14, 13, 12,
		10,
		7,
		5, 4, 3,
		2, 1, 0,
	}

	return bidirectionalBfs(c, goalCube, idG1, dirG1)
}

func idG2(c cube) cube {
	// TODO
}

func g2(c cube) []uint8 {
	var dirG2 = []uint8{
		16,
		13,
		10,
		7,
		5, 4, 3,
		2, 1, 0,
	}

	return bidirectionalBfs(c, goalCube, idG2, dirG2)
}

func idG3(c cube) cube {
}

func g3(c cube) []uint8 {
	var dirG3 = []uint8{
		16,
		13,
		10,
		7,
		4,
		1,
	}

	return bidirectionalBfs(c, goalCube, idG3, dirG3)
}

func thistlethwaite_uint8(init_moves []uint8) []uint8 {

	var c cube
	// make cube with init_moves
	moveG0 := g0(c)
	c = doMoves(c, moveG0)
	moveG1 := g1(c)
	c = doMoves(c, moveG1)
	moveG2 := g2(c)
	c = doMoves(c, moveG2)
	moveG3 := g3(c)

	return append(moveG0, append(moveG1, append(moveG2, moveG3...)...)...)
}

func Thistlethwaite(init_moves []makemove.RubikMoves) []makemove.RubikMoves {
	// convert moves to a other intern function ?
	moves := thistlethwaite_uint8( /* TODO */)
	// convert moves to makemove
	return /* TODO */
}
