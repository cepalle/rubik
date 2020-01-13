package solve

import "github.com/cepalle/rubik/internal/makemove"


// algo find: http://www.stefan-pochmann.info/spocc/other_stuff/tools/solver_thistlethwaite/solver_thistlethwaite_cpp.txt

type cube struct {
	PosP2 	[12]uint8
	RotP2 	[12]uint8

	PosP3	[8]uint8
	RotP3	[8]uint8
}

func doMove(c cube, move uint8) cube {
	// TODO
}

func doMoves(c cube, moves []uint8) cube {
	res := c

	for _, e := range moves {
		res = doMove(e, res)
	}

	return res
}

func idG0(c cube) cube {
	// TODO
}

func idG1(c cube) cube {
	// TODO
}

func idG2(c cube) cube {
	// TODO
}

func idG3(c cube) cube {
	// TODO
}

func idG4(c cube) cube {
	return c
}

func bidirectionalBfs(cur cube, goal cube, id func(c cube) cube, dir []uint8) []uint8 {

}

func g0(c cube) []uint8 {
	return nil
}

func g1(c cube) []uint8 {
	return nil
}

func g2(c cube) []uint8 {
	return nil
}

func g3(c cube) []uint8 {
	return nil
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
	moves := thistlethwaite_uint8(/* TODO */)
	// convert moves to makemove
	return /* TODO */
}
