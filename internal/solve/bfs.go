package solve

import (
	"github.com/cepalle/rubik/internal/makemove"
)

type Node struct {
	cube  makemove.Rubik
	moves []makemove.RubikMoves
}

func Bfs(r makemove.Rubik) []makemove.RubikMoves {
	hys := make(map[makemove.Rubik]bool)
	var pile []Node

	pile = append(pile, Node{r, []makemove.RubikMoves{}})
	hys[r] = true

	for len(pile) > 0 {
		cur := pile[0]
		pile = pile[1:]
		if cur.cube.IsResolve() {
			return cur.moves
		}

		for i := uint8(0); i < makemove.NbRubikMoves; i++ {
			var nCube = cur.cube.Move(makemove.AllRubikMoves[i].move)

			_, found := hys[nCube]
			if found {
				continue
			}
			hys[nCube] = true

			var mvsCp = cur.moves
			var nNode = Node{
				nCube,
				append(mvsCp, makemove.AllRubikMoves[i].move),
			}
			pile = append(pile, nNode)
		}
	}
	return nil
}
