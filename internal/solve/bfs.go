package solve

import (
	"github.com/cepalle/rubik/internal/makemove"
)

type node struct {
	cube  makemove.Rubik
	moves []makemove.RubikMoves
}

func Bfs(r makemove.Rubik) []makemove.RubikMoves {
	hys := make(map[makemove.Rubik]bool)
	var pile []node

	pile = append(pile, node{r, []makemove.RubikMoves{}})
	hys[r] = true

	for len(pile) > 0 {
		cur := pile[0]
		pile = pile[1:]
		if cur.cube.IsResolve() {
			return cur.moves
		}

		for _, m := range makemove.AllRubikMovesWithName {
			var nCube = cur.cube.DoMove(m.Move)

			_, found := hys[nCube]
			if found {
				continue
			}
			hys[nCube] = true

			var mvsCp = cur.moves
			var nNode = node{
				nCube,
				append(mvsCp, m.Move),
			}
			pile = append(pile, nNode)
		}
	}
	return nil
}
