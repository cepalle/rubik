package solve

import (
	"github.com/cepalle/rubik/internal/makemove"
)

const BfsDepth = 3
const Nnfilename = "./ff.network"
const NnDeepFilename = "./deep.gob"

func MakePatterns(bfsRes []NodeExp, bfsDepth uint32) [][][]float64 {
	var patterns [][][]float64

	for _, e := range bfsRes {
		var input []float64
		var output []float64

		input = makemove.RubikToNnInput(&e.Cube)
		for i := uint32(0); i <= bfsDepth; i++ {
			output = append(output, 0)
		}
		output[e.Depth] = 1

		line := [][]float64{input, output}
		patterns = append(patterns, line)
	}
	return patterns
}

func Equalize(bfsRes []NodeExp, bfsDepth uint32) []NodeExp {
	var nbByDepth []uint32

	for i := uint32(0); i <= bfsDepth; i++ {
		nbByDepth = append(nbByDepth, 0)
	}
	for _, e := range bfsRes {
		nbByDepth[e.Depth]++
	}
	var mx uint32
	mx = 0

	for _, e := range nbByDepth {
		if e > mx {
			mx = e
		}
	}

	for i := uint32(0); i < uint32(len(nbByDepth)); i++ {
		for nbByDepth[i] < mx {
			for j := int(0); j < len(bfsRes) && nbByDepth[i] < mx; j++ {
				if bfsRes[j].Depth == i {
					nbByDepth[i]++
					bfsRes = append(bfsRes, bfsRes[j])
				}
			}
		}

	}
	return bfsRes
}
