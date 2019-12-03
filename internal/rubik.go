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
	Clockwise        RubikFaceTurn = 0
	CounterClockwise RubikFaceTurn = 1
)

type Rubik struct {
	g1 [8]uint8
	g2 [12]uint8
	g3 [8]uint8
	g4 [12]uint8
}

type RubikMove struct {
	face RubikFace
	turn RubikFaceTurn
}

type moveFunction func(Rubik) Rubik

type dispatcher struct {
	move RubikMove
	fun  moveFunction
}

const dispatcherTab = [12]dispatcher{
	dispatcher{
		RubikMove{
			face: RubikFace.U,
			turn: RubikFaceTurn.Clockwise,
		},
		fun: clockwiseU,
	},
	dispatcher{
		RubikMove{
			face: RubikFace.U,
			turn: RubikFaceTurn.CounterClockwise,
		},
		fun: clockwiseU,
	},
	dispatcher{
		RubikMove{
			face: RubikFace.L,
			turn: RubikFaceTurn.Clockwise,
		},
		fun: clockwiseU,
	},
	dispatcher{
		RubikMove{
			face: RubikFace.L,
			turn: RubikFaceTurn.CounterClockwise,
		},
		fun: clockwiseU,
	},
	dispatcher{
		RubikMove{
			face: RubikFace.F,
			turn: RubikFaceTurn.Clockwise,
		},
		fun: clockwiseU,
	},
	dispatcher{
		RubikMove{
			face: RubikFace.F,
			turn: RubikFaceTurn.CounterClockwise,
		},
		fun: clockwiseU,
	},
	dispatcher{
		RubikMove{
			face: RubikFace.R,
			turn: RubikFaceTurn.Clockwise,
		},
		fun: clockwiseU,
	},
	dispatcher{
		RubikMove{
			face: RubikFace.R,
			turn: RubikFaceTurn.CounterClockwise,
		},
		fun: clockwiseU,
	},
	dispatcher{
		RubikMove{
			face: RubikFace.B,
			turn: RubikFaceTurn.Clockwise,
		},
		fun: clockwiseU,
	},
	dispatcher{
		RubikMove{
			face: RubikFace.B,
			turn: RubikFaceTurn.CounterClockwise,
		},
		fun: clockwiseU,
	},
	dispatcher{
		RubikMove{
			face: RubikFace.D,
			turn: RubikFaceTurn.Clockwise,
		},
		fun: clockwiseU,
	},
	dispatcher{
		RubikMove{
			face: RubikFace.D,
			turn: RubikFaceTurn.CounterClockwise,
		},
		fun: clockwiseU,
	},
}

func (cube Rubik) Move(m RubikMove) {
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

func clockwiseR(cube Rubik) {
	return cube
}

func clockwiseL(cube Rubik) {
	return cube
}

func clockwiseD(cube Rubik) {
	return cube
}

func clockwiseU(cube Rubik) {
	return cube
}

func clockwiseF(cube Rubik) {
	return cube
}

func clockwiseB(cube Rubik) {
	return cube
}

func counterClockwiseR(cube Rubik) {
	return cube
}

func counterClockwiseL(cube Rubik) {
	return cube
}

func counterClockwiseD(cube Rubik) {
	return cube
}

func counterClockwiseU(cube Rubik) {
	return cube
}

func counterClockwiseF(cube Rubik) {
	return cube
}

func counterClockwiseB(cube Rubik) {
	return cube
}
