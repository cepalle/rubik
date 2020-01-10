package solve

import "github.com/cepalle/rubik/internal/makemove"


func predicateG0(*makemove.Rubik) bool {

}

func g0(r makemove.Rubik) []makemove.RubikMoves {
	return IddfsPredicate(&r, predicateG0)
}

func g1(r makemove.Rubik) []makemove.RubikMoves  {

	return nil
}

func g2(r makemove.Rubik) []makemove.RubikMoves  {

	return nil
}

func g3(r makemove.Rubik) []makemove.RubikMoves  {

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
