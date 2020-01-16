package solve

import (
	"github.com/cepalle/rubik/internal/makemove"
)

type nodeBfs struct {
	cube  makemove.Rubik
	depth float64
}

func BfsScore(r makemove.Rubik, depth uint32, scoring func(*makemove.Rubik) float64) float64 {
	hys := make(map[makemove.Rubik]bool)
	var pile []nodeBfs
	var min float64

	pile = append(pile, nodeBfs{r, 0})
	hys[r] = true
	min = scoring(&r)

	for len(pile) > 0 {
		cur := pile[0]
		pile = pile[1:]
		if cur.depth == float64(depth) {
			continue
		}
		if cur.cube.IsResolve() {
			return cur.depth
		}

		for _, m := range makemove.AllRubikMovesWithName {
			var nCube = cur.cube.DoMove(m.Move)

			_, found := hys[nCube]
			if found {
				continue
			}
			hys[nCube] = true

			score := scoring(&nCube)
			if score < min {
				min = score
			}
			var nNode = nodeBfs{
				nCube,
				cur.depth + 1,
			}
			pile = append(pile, nNode)
		}
	}
	return min
}
