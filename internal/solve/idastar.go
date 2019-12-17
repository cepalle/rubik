package solve

import (
	"github.com/cepalle/rubik/internal/makemove"
	"math"
)

func AStart(r makemove.Rubik, scoring func(*makemove.Rubik) float64) []makemove.RubikMoves {
	return aStartWithScoreMax(r, scoring, math.MaxFloat64)
}

func aStartWithScoreMax(r makemove.Rubik, scoring func(*makemove.Rubik) float64, scoreMax float64) []makemove.RubikMoves {
	hys := make(map[makemove.Rubik]bool)
	open := New()
	open.Insert(node{r, []makemove.RubikMoves{}}, 0)

	for open.Len() > 0 {
		var cur, _ = open.Pop()
		curr := cur.(node)
		if curr.cube.IsResolve() {
			return curr.moves
		}
		for _, m := range makemove.AllRubikMovesWithName {
			var nCube = curr.cube.DoMove(m.Move)

			_, found := hys[nCube]
			if found {
				continue
			}
			hys[nCube] = true

			mvsCp := make([]makemove.RubikMoves, len(curr.moves))
			copy(mvsCp, curr.moves)
			var nNode = node{
				nCube,
				append(mvsCp, m.Move),
			}
			score := float64(len(nNode.moves)) + scoring(&nNode.cube)

			if score < scoreMax {
				open.Insert(nNode, score)
			}
		}
	}
	return nil
}

func IdaStar(r makemove.Rubik, scoring func(*makemove.Rubik) float64) []makemove.RubikMoves {
	var res []makemove.RubikMoves
	for i := float64(0); ; i += 10 {
		res = aStartWithScoreMax(r, scoring, i)
		if res != nil {
			return res
		}
	}
}
