package main

import (
	"encoding/gob"
	"fmt"
	"github.com/cepalle/rubik/internal/makemove"
	"github.com/cepalle/rubik/internal/solve"
	"github.com/patrikeh/go-deep"
	deeptraining "github.com/patrikeh/go-deep/training"
	"math/rand"
	"os"
)

func makeExample(bfsRes []solve.NodeExp, bfsDepth uint32) deeptraining.Examples {
	var example deeptraining.Examples

	for _, e := range bfsRes {
		var input []float64
		var output []float64

		input = makemove.RubikToNnInput(&e.Cube)
		for i := uint32(0); i <= bfsDepth; i++ {
			output = append(output, 0)
		}
		output[e.Depth] = 1

		example = append(example, deeptraining.Example{
			Input:    input,
			Response: output,
		})
	}
	return example
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

	examples := makeExample(solve.Equalize(all, solve.bfsDepth), solve.bfsDepth)
	rand.Shuffle(len(examples), func(i, j int) { examples[i], examples[j] = examples[j], examples[i] })
	rand.Shuffle(len(examples), func(i, j int) { examples[i], examples[j] = examples[j], examples[i] })
	rand.Shuffle(len(examples), func(i, j int) { examples[i], examples[j] = examples[j], examples[i] })
	rand.Shuffle(len(examples), func(i, j int) { examples[i], examples[j] = examples[j], examples[i] })

	n := deep.NewNeural(&deep.Config{
		Inputs: 48,
		Layout: []int{48, 96, 48, solve.bfsDepth + 1},
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
	optimizer := deeptraining.NewSGD(0.001, 0.001, 1e-6, true)
	// params: optimizer, verbosity (print stats at every 50th iteration)
	trainer := deeptraining.NewTrainer(optimizer, 10)

	training, heldout := examples.Split(0.25)
	trainer.Train(n, training, heldout, 500) // training, validation, iterations

	dataFile, err := os.Create(solve.NnDeepFilename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dataEncoder := gob.NewEncoder(dataFile)
	b, _ := n.Marshal()
	err = dataEncoder.Encode(b)
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
