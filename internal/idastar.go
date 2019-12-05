package internal

import (
	"github.com/jupp0r/go-priority-queue"
	"math"
)

type nodeAStar struct {
	cube  Rubik
	moves []RubikMoves
}

func AStart(r Rubik, scoring func(Rubik) float64) []RubikMoves {
	return AStartWithScoreMax(r, scoring, math.MaxFloat64)
}

func AStartWithScoreMax(r Rubik, scoring func(Rubik) float64, scoreMax float64) []RubikMoves {
	hys := make(map[Rubik]bool)
	open := pq.New()

	for ; open.Len() > 0; {
		var cur, _ = open.Pop()
		curr := cur.(nodeAStar)
		if curr.cube.IsResolve() {
			return curr.moves
		}
		for _, m := range AllRubikMovesWithName {
			var nCube = curr.cube.Move(m.move)

			_, found := hys[nCube]
			if found {
				continue
			}
			hys[nCube] = true

			var mvsCp = curr.moves
			var nNode = node{
				nCube,
				append(mvsCp, m.move),
			}
			score := scoring(nNode.cube)

			if score < scoreMax {
				open.Insert(nNode, score)
			}
		}
	}
	return nil
}

func IdaStar(r Rubik, scoring func(Rubik) float64) []RubikMoves {
	var res []RubikMoves
	for i := float64(0); ; i += 10 {
		res = AStartWithScoreMax(r, scoring, i)
		if res != nil {
			return res
		}
	}
}
