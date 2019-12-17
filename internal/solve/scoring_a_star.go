package solve

import (
	"encoding/gob"
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
	"github.com/goml/gobrain"
	"github.com/goml/gobrain/persist"
	"github.com/patrikeh/go-deep"
	"log"
	"os"
)

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

func MakeNNDeepScoring(filename string) func(cube *makemove.Rubik) float64 {
	n := deep.NewNeural(&deep.Config{
		Inputs: 48,
		Layout: []int{48, 48, Bfs_depth + 1},
		/* Activation functions: Sigmoid, Tanh, ReLU, Linear */
		Activation: deep.ActivationSigmoid,
		/* Determines output layer activation & loss function:
		ModeRegression: linear outputs with MSE loss
		ModeMultiClass: softmax output with Cross Entropy loss
		ModeMultiLabel: sigmoid output with Cross Entropy loss
		ModeBinary: sigmoid output with binary CE loss */
		Mode: deep.ModeMultiClass,
		/* Weight initializers: {deep.NewNormal(μ, σ), deep.NewUniform(μ, σ)} */
		Weight: deep.NewNormal(1.0, 0.0),
		Bias:   true,
	})

	dataFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&n)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = dataFile.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return func(cube *makemove.Rubik) float64 {
		return nnOutputToScoring(n.Predict(makemove.Rubik_to_nn_input(cube)))
	}
}
