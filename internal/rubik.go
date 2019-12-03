package internal

type RubikFace uint8

const (
	U RubikFace = 0
	L RubikFace = 1
	F RubikFace = 2
	R RubikFace = 3
	B RubikFace = 4
	D RubikFace = 5
)

type RubikFaceTurn uint8

const (
	Right RubikFaceTurn = 0
	Left  RubikFaceTurn = 1
)

type Rubik struct {
	g1 [8]    uint8
	g2 [12]    uint8
	g3 [8]    uint8
	g4 [12]    uint8
}

type RubikMove struct {
	face RubikFace
	turn RubikFaceTurn
}

func (p Rubik) move(m RubikMove) {
	if m.turn == Right {
		if m.face == U {
		} else if m.face == L {
		} else if m.face == F {
		} else if m.face == R {
		} else if m.face == B {
		} else if m.face == D {
		}
	} else {
		if m.face == U {
		} else if m.face == L {
		} else if m.face == F {
		} else if m.face == R {
		} else if m.face == B {
		} else if m.face == D {
		}
	}
}
