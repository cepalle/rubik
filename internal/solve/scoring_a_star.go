package solve

import (
	"github.com/cepalle/rubik/internal/makemove"
	"github.com/goml/gobrain/persist"
	"log"
)
import "github.com/goml/gobrain"

func ScoringHamming(cube *makemove.Rubik) float64 {
	var i uint8
	tot := float64(0)

	for i = 0; i < 12; i++ {
		if cube.RotP2[i] != 0 {
			tot++
		}
	}
	for i = 0; i < 12; i++ {
		if cube.PosP2[i] != i {
			tot++
		}
	}
	for i = 0; i < 24; i++ {
		if cube.PosFP3[i] != i {
			tot++
		}
	}
	return tot
}

func nnOutputToScoring(out []float64) float64 {
	var res float64 = 0

	for i := 0; i < len(out); i++ {
		res = res + float64(i)*out[i]
	}
	return res
}

func MakeNNScoring(filename string) func(cube *makemove.Rubik) float64 {
	ff := &gobrain.FeedForward{}
	err := persist.Load(filename, &ff)
	if err != nil {
		log.Println("impossible to save network on file: ", err.Error())
	}

	return func(cube *makemove.Rubik) float64 {
		return nnOutputToScoring(ff.Update(makemove.Rubik_to_nn_input(cube)))
	}
}
