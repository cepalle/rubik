package internal

func dls(r Rubik, depth uint32) []RubikMoves {
	var res []RubikMoves

	if depth == 0 && r.IsResolve() {
		return []RubikMoves{}
	}

	for _, m := range AllRubikMovesWithName {
		res = dls(r.Move(m.move), depth-1)
		if res != nil {
			return append(res, m.move)
		}
	}
	return nil
}

func Iddfs(r Rubik) []RubikMoves {
	var res []RubikMoves

	for i := uint32(0); ; i++ {
		res = dls(r, i)
		if res != nil {
			return res
		}
	}
}
