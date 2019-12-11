package solve

import (
	"github.com/cepalle/rubik/internal/makemove"
	"github.com/jupp0r/go-priority-queue"
	"math"
)

func AStart(r makemove.Rubik, scoring func(*makemove.Rubik) float64) []makemove.RubikMoves {
	return aStartWithScoreMax(r, scoring, math.MaxFloat64)
}

func aStartWithScoreMax(r makemove.Rubik, scoring func(*makemove.Rubik) float64, scoreMax float64) []makemove.RubikMoves {
	hys := make(map[makemove.Rubik]bool)
	open := pq.New()

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
			score := scoring(&nNode.cube)

			if score < scoreMax {
				open.Insert(nNode, -(float64(len(nNode.moves)) + score))
			}
		}
	}
	return nil
}

func ScoringHamming(cube *makemove.Rubik) float64 {
	var i uint8
	tot := float64(0)

	for i = 0; i < 12; i++ {
		if cube.RotP2[i] != 0 {
			tot++
		}
	}
	for i = 0; i < 8; i++ {
		if cube.RotP3[i] != 0 {
			tot++
		}
	}
	for i = 0; i < 12; i++ {
		if cube.PosP2[i] != i {
			tot++
		}
	}
	for i = 0; i < 8; i++ {
		if cube.PosP3[i] != i {
			tot++
		}
	}
	return tot
}

func IdaStar(r makemove.Rubik, scoring func(*makemove.Rubik) float64) []makemove.RubikMoves {
	var res []makemove.RubikMoves
	for i := float64(0); ; i += 30 {
		res = aStartWithScoreMax(r, scoring, i)
		if res != nil {
			return res
		}
	}
}
