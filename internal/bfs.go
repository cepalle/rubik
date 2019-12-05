package internal

type node struct {
	cube  Rubik
	moves []RubikMoves
}

func Bfs(r Rubik) []RubikMoves {
	hys := make(map[Rubik]bool)
	var pile []node

	pile = append(pile, node{r, []RubikMoves{}})
	hys[r] = true

	for len(pile) > 0 {
		cur := pile[0]
		pile = pile[1:]
		if cur.cube.IsResolve() {
			return cur.moves
		}

		for _, m := range AllRubikMovesWithName {
			var nCube = cur.cube.Move(m.move)

			_, found := hys[nCube]
			if found {
				continue
			}
			hys[nCube] = true

			var mvsCp = cur.moves
			var nNode = node{
				nCube,
				append(mvsCp, m.move),
			}
			pile = append(pile, nNode)
		}
	}
	return nil
}
