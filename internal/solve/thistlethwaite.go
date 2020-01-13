package solve

import "github.com/cepalle/rubik/internal/makemove"

func predicateG0(r *makemove.Rubik) bool {
	for i := uint8(0); i < uint8(len(r.PosP2)); i++ {
		var cp = *r
		var d = IddfsPredicate(&cp, func(rubik *makemove.Rubik) bool {
			return rubik.PosP2[i] == i
		})
		for _, rm := range d {
			if (rm.Face == makemove.U || rm.Face == makemove.D) && rm.NbTurn != 2 {
				return false
			}
		}
	}
	return true
}

func g0(r makemove.Rubik) []makemove.RubikMoves {
	return IddfsPredicate(&r, predicateG0)
}

func g1(r makemove.Rubik) []makemove.RubikMoves {

	return nil
}

func g2(r makemove.Rubik) []makemove.RubikMoves {

	return nil
}

func g3(r makemove.Rubik) []makemove.RubikMoves {

	return nil
}

func Thistlethwaite(r makemove.Rubik) []makemove.RubikMoves {
	moveG0 := g0(r)
	r = r.DoMoves(moveG0)
	moveG1 := g1(r)
	r = r.DoMoves(moveG0)
	moveG2 := g2(r)
	r = r.DoMoves(moveG0)
	moveG3 := g3(r)

	return append(moveG0, append(moveG1, append(moveG2, moveG3...)...)...)
}
