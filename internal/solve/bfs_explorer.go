package solve

import (
	"github.com/cepalle/rubik/internal/makemove"
)

type NodeExp struct {
	Depth uint32
	Cube  makemove.Rubik
}

func Bfs_explorer(depth uint32) []NodeExp {
	var res []NodeExp
	hys := make(map[makemove.Rubik]bool)
	var pile []NodeExp

	solveCube := makemove.InitRubik()
	pile = append(pile, NodeExp{0, solveCube})
	hys[solveCube] = true

	for len(pile) > 0 {
		cur := pile[0]
		pile = pile[1:]

		res = append(res, cur)

		if cur.Depth >= depth {
			continue
		}

		for _, m := range makemove.AllRubikMovesWithName {
			var nCube = cur.Cube.DoMove(m.Move)

			_, found := hys[nCube]
			if found {
				continue
			}
			hys[nCube] = true

			var nNode = NodeExp{
				cur.Depth + 1,
				nCube,
			}
			pile = append(pile, nNode)
		}
	}
	return res
}
