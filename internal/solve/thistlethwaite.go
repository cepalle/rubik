package solve

import (
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
			var orientationDelta uint8
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
			var orientationDelta uint8
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

	return face*3 + (2 - turns)
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
	// println("BFS failed...")
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
	//-- Phase 2: Corner orientations, E slice edges. g1 -> g23
	for i := uint8(0); i < 12; i++ {
		c.PosF2[i] = c.PosF2[i] / 8
	}
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

func idG23(c cube) cube {
	//--- Phase 3: All. g2 -> g4
	return c
}

func g23(c cube) []uint8 {
	var dirG2 = []uint8{
		16,
		13,
		10,
		7,
		5, 4, 3,
		2, 1, 0,
	}

	return bidirectionalBfs(c, goalCube, idG23, dirG2)
}

func thistlethwaiteUint8(initMoves []uint8) []uint8 {

	var c cube = goalCube

	for _, m := range initMoves {
		c = doMove(c, m)
	}

	// fmt.Printf("%+v\n", c)

	// println("G0 -> G1 Start")
	moveG0 := g0(c)
	c = doMoves(c, moveG0)
	// fmt.Printf("%+v\n", c)

	// println("G1 -> G2 Start")
	moveG1 := g1(c)
	c = doMoves(c, moveG1)
	// fmt.Printf("%+v\n", c)

	// println("G2 -> G4 Start")
	moveG23 := g23(c)
	doMoves(c, moveG23)
	// fmt.Printf("%+v\n", c)

	// println("END")

	return append(moveG0, append(moveG1, moveG23...)...)
}

var uint8ToRbikMoves = [makemove.NbRubikMoves]makemove.RubikMoves{
	{Face: makemove.U, Turn: makemove.Clockwise, NbTurn: 1},
	{Face: makemove.U, Turn: makemove.Clockwise, NbTurn: 2},
	{Face: makemove.U, Turn: makemove.CounterClockwise, NbTurn: 1},
	{Face: makemove.D, Turn: makemove.Clockwise, NbTurn: 1},
	{Face: makemove.D, Turn: makemove.Clockwise, NbTurn: 2},
	{Face: makemove.D, Turn: makemove.CounterClockwise, NbTurn: 1},
	{Face: makemove.F, Turn: makemove.Clockwise, NbTurn: 1},
	{Face: makemove.F, Turn: makemove.Clockwise, NbTurn: 2},
	{Face: makemove.F, Turn: makemove.CounterClockwise, NbTurn: 1},
	{Face: makemove.B, Turn: makemove.Clockwise, NbTurn: 1},
	{Face: makemove.B, Turn: makemove.Clockwise, NbTurn: 2},
	{Face: makemove.B, Turn: makemove.CounterClockwise, NbTurn: 1},
	{Face: makemove.L, Turn: makemove.Clockwise, NbTurn: 1},
	{Face: makemove.L, Turn: makemove.Clockwise, NbTurn: 2},
	{Face: makemove.L, Turn: makemove.CounterClockwise, NbTurn: 1},
	{Face: makemove.R, Turn: makemove.Clockwise, NbTurn: 1},
	{Face: makemove.R, Turn: makemove.Clockwise, NbTurn: 2},
	{Face: makemove.R, Turn: makemove.CounterClockwise, NbTurn: 1},
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

func Thistlethwaite(initMoves []makemove.RubikMoves) []makemove.RubikMoves {
	var movesUnit8 []uint8

	for i := 0; i < len(initMoves); i++ {
		movesUnit8 = append(movesUnit8, rubikMovesToUint8(initMoves[i]))
	}

	solUint8 := thistlethwaiteUint8(movesUnit8)
	var sol []makemove.RubikMoves

	for i := 0; i < len(solUint8); i++ {
		sol = append(sol, uint8ToRbikMoves[solUint8[i]])
	}

	return sol
}

func BidiBfs(initMoves []makemove.RubikMoves) []makemove.RubikMoves {
	var movesUnit8 []uint8

	for i := 0; i < len(initMoves); i++ {
		movesUnit8 = append(movesUnit8, rubikMovesToUint8(initMoves[i]))
	}

	var c cube = goalCube

	for _, m := range movesUnit8 {
		c = doMove(c, m)
	}

	var dirG0 = []uint8{
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 10, 11,
		12, 13, 14,
		15, 16, 17,
	}

	solUint8 := bidirectionalBfs(c, goalCube, idG23, dirG0)
	var sol []makemove.RubikMoves

	for i := 0; i < len(solUint8); i++ {
		sol = append(sol, uint8ToRbikMoves[solUint8[i]])
	}

	return sol
}
