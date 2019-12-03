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

func clockwiseR(cube Rubik) Rubik {
	return cube
}

func clockwiseL(cube Rubik) Rubik {
	return cube
}

func clockwiseD(cube Rubik) Rubik {
	return cube
}

func clockwiseU(cube Rubik) Rubik {
	return cube
}

func clockwiseF(cube Rubik) Rubik {
	return cube
}

func clockwiseB(cube Rubik) Rubik {
	return cube
}

func counterClockwiseR(cube Rubik) Rubik {
	return cube
}

func counterClockwiseL(cube Rubik) Rubik {
	return cube
}

func counterClockwiseD(cube Rubik) Rubik {
	return cube
}

func counterClockwiseU(cube Rubik) Rubik {
	return cube
}

func counterClockwiseF(cube Rubik) Rubik {
	return cube
}

func counterClockwiseB(cube Rubik) Rubik {
	return cube
}

const dispatcherLen = 12

var dispatcherTab = [dispatcherLen]dispatcher{
	dispatcher{
		move: RubikMove{
			face: U,
			turn: Clockwise,
		},
		fun: clockwiseU,
	},
	dispatcher{
		move: RubikMove{
			face: U,
			turn: CounterClockwise,
		},
		fun: counterClockwiseU,
	},
	dispatcher{
		move: RubikMove{
			face: L,
			turn: Clockwise,
		},
		fun: clockwiseL,
	},
	dispatcher{
		move: RubikMove{
			face: L,
			turn: CounterClockwise,
		},
		fun: counterClockwiseL,
	},
	dispatcher{
		move: RubikMove{
			face: F,
			turn: Clockwise,
		},
		fun: clockwiseF,
	},
	dispatcher{
		move: RubikMove{
			face: F,
			turn: CounterClockwise,
		},
		fun: counterClockwiseF,
	},
	dispatcher{
		move: RubikMove{
			face: R,
			turn: Clockwise,
		},
		fun: clockwiseR,
	},
	dispatcher{
		move: RubikMove{
			face: R,
			turn: CounterClockwise,
		},
		fun: counterClockwiseR,
	},
	dispatcher{
		move: RubikMove{
			face: B,
			turn: Clockwise,
		},
		fun: clockwiseB,
	},
	dispatcher{
		move: RubikMove{
			face: B,
			turn: CounterClockwise,
		},
		fun: counterClockwiseB,
	},
	dispatcher{
		move: RubikMove{
			face: D,
			turn: Clockwise,
		},
		fun: clockwiseD,
	},
	dispatcher{
		move: RubikMove{
			face: D,
			turn: CounterClockwise,
		},
		fun: counterClockwiseD,
	},
}

func (cube Rubik) Move(m RubikMove) Rubik {
	for i := 0; i < dispatcherLen; i++ {
		if m == dispatcherTab[i].move {
			return dispatcherTab[i].fun(cube)
		}
	}
	return cube
}
