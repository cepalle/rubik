package solve

import (
	"github.com/cepalle/rubik/internal/makemove"
)

func dls(r makemove.Rubik, depth uint32) []makemove.RubikMoves {
	var res []makemove.RubikMoves

	if depth == 0 && r.IsResolve() {
		return []makemove.RubikMoves{}
	}

	for _, m := range makemove.AllRubikMovesWithName {
		res = dls(r.Move(m.move), depth-1)
		if res != nil {
			return append(res, m.move)
		}
	}
	return nil
}

func Iddfs(r makemove.Rubik) []makemove.RubikMoves {
	var res []makemove.RubikMoves

	for i := uint32(0); ; i++ {
		res = dls(r, i)
		if res != nil {
			return res
		}
	}
}
