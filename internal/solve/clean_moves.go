package solve

import (
	"github.com/cepalle/rubik/internal/makemove"
)

var order = [3][2]makemove.RubikFace{
	{makemove.U, makemove.D},
	{makemove.F, makemove.B},
	{makemove.L, makemove.R},
}

func nbTM(m makemove.RubikMoves) uint8 {
	if m.Turn == makemove.CounterClockwise {
		return -m.NbTurn
	}
	return m.NbTurn
}

func CleanMoves(m []makemove.RubikMoves) []makemove.RubikMoves {

	mCp := make([]makemove.RubikMoves, len(m))
	copy(mCp, m)

Boucle1:
	for i := 0; i+1 < len(mCp); i++ {
		for j := 0; j < 3; j++ {
			if mCp[i].Face == order[j][0] && mCp[i+1].Face == order[j][1] {
				tmp := mCp[i]
				mCp[i] = mCp[i+1]
				mCp[i+1] = tmp
				goto Boucle1
			}
		}
	}

Boucle2:
	for i := 0; i+1 < len(mCp); i++ {
		if mCp[i].Face == mCp[i+1].Face {
			var n1 = nbTM(mCp[i])
			var n2 = nbTM(mCp[i+1])
			var nTot = (n1 + n2) % 4
			if nTot == 0 {
				copy(mCp[i:], mCp[i+2:])
				mCp = mCp[:len(mCp)-2]
			} else {
				copy(mCp[i:], mCp[i+1:]) // Shift a[i+1:] left one index.
				mCp = mCp[:len(mCp)-1]   // Truncate slice.
				if nTot == 3 {
					mCp[i].Turn = makemove.CounterClockwise
					mCp[i].NbTurn = 1
				} else {
					mCp[i].Turn = makemove.Clockwise
					mCp[i].NbTurn = nTot
				}
			}
			goto Boucle2
		}
	}
	return mCp
}
