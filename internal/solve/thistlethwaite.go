package solve

import (
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
)

/*
	Algo: http://www.stefan-pochmann.info/spocc/other_stuff/tools/solver_thistlethwaite/solver_thistlethwaite_cpp.txt

	turns = move % 3
	face = move / 3
*/

var affectedCubies = [6][2][4]uint8{
	{{0, 1, 2, 3}, {0, 1, 2, 3}},   // U
	{{4, 7, 6, 5}, {4, 5, 6, 7}},   // D
	{{0, 9, 4, 8}, {0, 3, 5, 4}},   // F
	{{2, 10, 6, 11}, {2, 1, 7, 6}}, // B
	{{3, 11, 7, 9}, {3, 2, 6, 5}},  // L
	{{1, 8, 5, 10}, {1, 0, 4, 7}},  // R
}

type cube struct {
	PosF2 [12]uint8
	RotF2 [12]uint8

	PosF3 [8]uint8
	RotF3 [8]uint8
}

var goalCube = cube{
	[12]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	[12]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[8]uint8{0, 1, 2, 3, 4, 5, 6, 7},
	[8]uint8{0, 0, 0, 0, 0, 0, 0, 0},
}

func doMove(cur cube, move uint8) cube {
	var nbTurns = move%3 + 1
	var face = move / 3

	for t := uint8(0); t < nbTurns; t++ {
		var oldC = cur

		for i := uint8(0); i < 4; i++ {
			var target uint8 = affectedCubies[face][0][i]
			var killer uint8 = affectedCubies[face][0][(i+1)%4]
			var orientationDelta uint8 = 0
			// F or B
			if face == 2 || face == 3 {
				orientationDelta = 1
			}

			cur.PosF2[target] = oldC.PosF2[killer]
			cur.RotF2[target] = oldC.RotF2[killer] + orientationDelta
		}

		for i := uint8(0); i < 4; i++ {
			var target uint8 = affectedCubies[face][1][i]
			var killer uint8 = affectedCubies[face][1][(i+1)%4]
			var orientationDelta uint8 = 0
			// F or B or L or R
			if face > 1 {
				orientationDelta = 2 - (i % 2)
			}

			cur.PosF3[target] = oldC.PosF3[killer]
			cur.RotF3[target] = oldC.RotF3[killer] + orientationDelta
		}
	}

	for i := 0; i < 12; i++ {
		cur.RotF2[i] = (cur.RotF2[i] + 4) % 2
	}

	for i := 0; i < 8; i++ {
		cur.RotF3[i] = (cur.RotF3[i] + 6) % 3
	}

	return cur
}

func doMoves(c cube, moves []uint8) cube {
	res := c

	for _, e := range moves {
		res = doMove(res, e)
	}

	return res
}

func reversUint8Move(m uint8) uint8 {
	turns := m % 3
	face := m / 3

	return face * 3 + (2 - turns)
}

func appendMovesReversed(a []uint8, b []uint8) []uint8 {
	aCp := make([]uint8, len(a))
	copy(aCp, a)

	for i := len(b); i > 0; i-- {
		aCp = append(aCp, reversUint8Move(b[i-1]))
	}

	return aCp
}

func bidirectionalBfs(src cube, dst cube, id func(c cube) cube, dir []uint8) []uint8 {

	if id(src) == id(dst) {
		return []uint8{}
	}

	type node struct {
		cube  cube
		moves []uint8
	}

	hysSrc := make(map[cube][]uint8)
	hysSrc[id(src)] = []uint8{}

	var pileSrc []node
	pileSrc = append(pileSrc, node{src, []uint8{}})

	hysDst := make(map[cube][]uint8)
	hysDst[id(dst)] = []uint8{}

	var pileDst []node
	pileDst = append(pileDst, node{dst, []uint8{}})

	for len(pileSrc) > 0 || len(pileDst) > 0 {

		if len(pileSrc) > 0 {
			// SRC
			curSrc := pileSrc[0]
			pileSrc = pileSrc[1:]

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

				movesDst, found := hysDst[idCubeSrc]
				if found {
					return appendMovesReversed(nwMvs, movesDst)
				}

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

			for _, d := range dir {
				var nwCube = doMove(curDst.cube, d)

				idCubeDst := id(nwCube)
				_, found := hysDst[idCubeDst]
				if found {
					continue
				}

				mvsCp := make([]uint8, len(curDst.moves))
				copy(mvsCp, curDst.moves)
				nwMvs := append(mvsCp, d)

				movesSrc, found := hysSrc[idCubeDst]
				if found {
					return appendMovesReversed(movesSrc, nwMvs)
				}

				var nNode = node{
					nwCube,
					nwMvs,
				}

				hysDst[idCubeDst] = nwMvs
				pileDst = append(pileDst, nNode)
			}
		}

	}
	println("BFS failed...")
	return nil
}

func idG0(c cube) cube {
	for i := uint8(0); i < 8; i++ {
		c.RotF3[i] = 0
		c.PosF3[i] = 0
	}

	for i := uint8(0); i < 12; i++ {
		c.PosF2[i] = 0
	}
	return c
}

func g0(c cube) []uint8 {
	//--- Phase 1: Edge orientations. g0 -> g1
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
	//-- Phase 2: Corner orientations, E slice edges. g1 -> g2
	for i := uint8(0); i < 8; i++ {
		// c.PosF3[i] = 0
	}

	// var r2 uint8 = c.RotF2[11]
	for i := uint8(0); i < 12; i++ {
		c.PosF2[i] = c.PosF2[i] / 8
		c.RotF2[i] = 0
	}
	// c.RotF2[11] = r2
	return c
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

func bool_to_uint8(a bool) uint8 {
	if a {
		return 1
	}
	return 0
}

func idG2(c cube) cube {
	//--- Phase 3: Edge slices M and S, corner tetrads, overall parity. g2 -> g3

	var r2 uint8 = 0
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 8; j++ {
			r2 = r2 ^ bool_to_uint8(c.PosF3[i] > c.PosF3[j])
		}
	}

	for i := uint8(0); i < 8; i++ {
		c.RotF3[i] = 0

		c.PosF3[i] = c.PosF3[i] & 5
	}

	for i := uint8(0); i < 12; i++ {
		c.RotF2[i] = 0

		c.PosF2[i] = 2
		if c.PosF2[i] < 8 {
			c.PosF2[i] = c.PosF2[i] % 2
		}
	}
	c.RotF2[0] = r2

	return c
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
	//--- Phase 4: The rest. g3 -> g4
	return c
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
	/*
	var dirG0 = []uint8{
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 10, 11,
		12, 13, 14,
		15, 16, 17,
	}
	*/

	return bidirectionalBfs(c, goalCube, idG3, dirG3)
}

func thistlethwaiteUint8(init_moves []uint8) []uint8 {

	var c cube = goalCube

	for _, m := range init_moves {
		c = doMove(c, m)
	}

	fmt.Printf("%+v\n", c)

	println("G0 Start")
	moveG0 := g0(c)
	c = doMoves(c, moveG0)
	fmt.Printf("%+v\n", c)

	println("G1 Start")
	moveG1 := g1(c)
	c = doMoves(c, moveG1)
	fmt.Printf("%+v\n", c)

	var moveG2 []uint8
	// println("G2 Start")
	// moveG2 := g2(c)
	// c = doMoves(c, moveG2)
	// fmt.Printf("%+v\n", c)

	println("G3 Start")
	moveG3 := g3(c)
	c = doMoves(c, moveG3)
	fmt.Printf("%+v\n", c)

	println("END")

	return append(moveG0, append(moveG1, append(moveG2, moveG3...)...)...)
}

var uint8ToRbikMoves = [makemove.NbRubikMoves]makemove.RubikMoves{
	makemove.RubikMoves{
		makemove.U, makemove.Clockwise, 1,
	},
	makemove.RubikMoves{
		makemove.U, makemove.Clockwise, 2,
	},
	makemove.RubikMoves{
		makemove.U, makemove.CounterClockwise, 1,
	},
	makemove.RubikMoves{
		makemove.D, makemove.Clockwise, 1,
	},
	makemove.RubikMoves{
		makemove.D, makemove.Clockwise, 2,
	},
	makemove.RubikMoves{
		makemove.D, makemove.CounterClockwise, 1,
	},
	makemove.RubikMoves{
		makemove.F, makemove.Clockwise, 1,
	},
	makemove.RubikMoves{
		makemove.F, makemove.Clockwise, 2,
	},
	makemove.RubikMoves{
		makemove.F, makemove.CounterClockwise, 1,
	},
	makemove.RubikMoves{
		makemove.B, makemove.Clockwise, 1,
	},
	makemove.RubikMoves{
		makemove.B, makemove.Clockwise, 2,
	},
	makemove.RubikMoves{
		makemove.B, makemove.CounterClockwise, 1,
	},
	makemove.RubikMoves{
		makemove.L, makemove.Clockwise, 1,
	},
	makemove.RubikMoves{
		makemove.L, makemove.Clockwise, 2,
	},
	makemove.RubikMoves{
		makemove.L, makemove.CounterClockwise, 1,
	},
	makemove.RubikMoves{
		makemove.R, makemove.Clockwise, 1,
	},
	makemove.RubikMoves{
		makemove.R, makemove.Clockwise, 2,
	},
	makemove.RubikMoves{
		makemove.R, makemove.CounterClockwise, 1,
	},
}

func rubikMovesToUint8(m makemove.RubikMoves) uint8 {
	for i := uint8(0); ; i++ {
		if m.NbTurn == uint8ToRbikMoves[i].NbTurn &&
			m.Face == uint8ToRbikMoves[i].Face &&
			m.Turn == uint8ToRbikMoves[i].Turn {
			return i
		}
	}
}

func Thistlethwaite(init_moves []makemove.RubikMoves) []makemove.RubikMoves {
	var movesUnit8 []uint8

	for i := 0; i < len(init_moves); i++ {
		movesUnit8 = append(movesUnit8, rubikMovesToUint8(init_moves[i]))
	}

	solUint8 := thistlethwaiteUint8(movesUnit8)
	var sol []makemove.RubikMoves

	for i := 0; i < len(solUint8); i++ {
		sol = append(sol, uint8ToRbikMoves[solUint8[i]])
	}

	return sol
}
