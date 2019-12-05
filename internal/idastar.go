package internal

func AStart(r Rubik, scoring func(Rubik) uint32) []RubikMoves {
	return AStartWithScoreMax(r, scoring, ^uint32(0))
}

func AStartWithScoreMax(r Rubik, scoring func(Rubik) uint32, scoreMax uint32) []RubikMoves {

}

func IdaStar(r Rubik, scoring func(Rubik) uint32) []RubikMoves {
	var res []RubikMoves

	for i := uint32(0); ; i += 10 {
		res = AStartWithScoreMax(r, scoring, i)
		if res != nil {
			return res
		}
	}

}
