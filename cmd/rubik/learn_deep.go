package main

import (
	"encoding/gob"
	"fmt"
	"github.com/cepalle/rubik/internal/learn"
	"math/rand"
	"os"
	"github.com/cepalle/rubik/internal/solve"
	deep "github.com/patrikeh/go-deep"
	"github.com/patrikeh/go-deep/training"
)

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

	patterns := learn.MakePatterns(learn.Equalize(all, learn.Bfs_depth), learn.Bfs_depth)

	n := deep.NewNeural(&deep.Config{
		Inputs: 2,
		Layout: []int{48, 48, 36, 24, 12, learn.Bfs_depth + 1},
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
	optimizer := training.NewSGD(0.05, 0.1, 1e-6, true)
	// params: optimizer, verbosity (print stats at every 50th iteration)
	trainer := training.NewTrainer(optimizer, 50)

	training, heldout := data.Split(0.5)
	trainer.Train(n, training, heldout, 1000) // training, validation, iterations

}
