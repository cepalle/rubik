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

func MakeBfsScore(depth uint32, scoring func(c *makemove.Rubik) float64) func(*makemove.Rubik) float64 {
	return func(c *makemove.Rubik) float64 {
		return BfsScore(*c, depth, scoring)
	}
}

func ScoringHuman(cube *makemove.Rubik) float64 {
	return float64(len(MechanicalHuman(*cube, false)))
}

func nnOutputToScoring(out []float64) float64 {
	var res float64

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
		return nnOutputToScoring(ff.Update(makemove.RubikToNnInput(cube)))
	}
}

func MakeNNDeepScoring(filename string) func(cube *makemove.Rubik) float64 {
	var b []byte
	var n *deep.Neural
	dataFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&b)
	n, _ = deep.Unmarshal(b)

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
		return nnOutputToScoring(n.Predict(makemove.RubikToNnInput(cube)))
	}
}
