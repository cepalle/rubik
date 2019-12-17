package main

import (
	"encoding/gob"
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
	"github.com/cepalle/rubik/internal/solve"
	"github.com/patrikeh/go-deep"
	deep_training "github.com/patrikeh/go-deep/training"
	"math/rand"
	"os"
)

func makeExemple(bfsRes []solve.NodeExp, bfs_depth uint32) deep_training.Examples {
	var exemple deep_training.Examples

	for _, e := range bfsRes {
		var input []float64
		var output []float64

		input = makemove.Rubik_to_nn_input(&e.Cube)
		for i := uint32(0); i <= bfs_depth; i++ {
			output = append(output, 0)
		}
		output[e.Depth] = 1

		exemple = append(exemple, deep_training.Example{
			Input:    input,
			Response: output,
		})
	}
	return exemple
}

func main() {
	var all []solve.NodeExp
	dataFile, err1 := os.Open("node_exp.gob")

	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err1 = dataDecoder.Decode(&all)

	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}

	err1 = dataFile.Close()
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}

	rand.Seed(0)

	// ---

	exemples := makeExemple(solve.Equalize(all, solve.Bfs_depth), solve.Bfs_depth)
	rand.Shuffle(len(exemples), func(i, j int) { exemples[i], exemples[j] = exemples[j], exemples[i] })
	rand.Shuffle(len(exemples), func(i, j int) { exemples[i], exemples[j] = exemples[j], exemples[i] })
	rand.Shuffle(len(exemples), func(i, j int) { exemples[i], exemples[j] = exemples[j], exemples[i] })
	rand.Shuffle(len(exemples), func(i, j int) { exemples[i], exemples[j] = exemples[j], exemples[i] })

	n := deep.NewNeural(&deep.Config{
		Inputs: 48,
		Layout: []int{48, 96, 48, solve.Bfs_depth + 1},
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

	// params: learning rate, momentum, alpha decay, nesterov
	optimizer := deep_training.NewSGD(0.001, 0.001, 1e-6, true)
	// params: optimizer, verbosity (print stats at every 50th iteration)
	trainer := deep_training.NewTrainer(optimizer, 10)

	training, heldout := exemples.Split(0.25)
	trainer.Train(n, training, heldout, 500) // training, validation, iterations

	dataFile, err := os.Create(solve.NnDeepFilename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dataEncoder := gob.NewEncoder(dataFile)
	err = dataEncoder.Encode(*n)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = dataFile.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
