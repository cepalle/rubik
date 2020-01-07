package solve

import "github.com/cepalle/rubik/internal/makemove"

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
