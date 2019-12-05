package internal

func Dls(r Rubik, depth uint32) []RubikMoves {
	var res []RubikMoves

	if depth == 0 && r.IsResolve() {
		return []RubikMoves{}
	}

	for _, m := range AllRubikMovesWithName {
		res = Dls(r.Move(m.move), depth-1)
		if res != nil {
			return append(res, m.move)
		}
	}
	return nil
}

func Iddfs(r Rubik) []RubikMoves {
	var res []RubikMoves

	for i := uint32(0); ; i++ {
		res = Dls(r, i)
		if res != nil {
			return res
		}
	}
}
