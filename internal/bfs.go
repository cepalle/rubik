package internal

type Node struct {
	cube  Rubik
	moves []RubikMoves
}

func contains(s []Rubik, e Rubik) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Bfs(r Rubik) []RubikMoves {
	var hys []Rubik
	var pile []Node

	pile = append(pile, Node{r, []RubikMoves{}})
	hys = append(hys, r)

	for ; len(pile) > 0; {
		cur := pile[0]
		pile = pile[1:]
		if cur.cube.IsResolve() {
			return cur.moves
		}

		for i := uint8(0); i < NbRubikMoves; i++ {
			var nCube = cur.cube.Move(AllRubikMoves[i])

			if contains(hys, nCube) {
				continue
			}
			hys = append(hys, nCube)

			var mvsCp = cur.moves
			var nNode = Node{
				nCube,
				append(mvsCp, AllRubikMoves[i]),
			}
			pile = append(pile, nNode)
		}
	}
	return []RubikMoves{}
}
