package internal

type RubikFace uint8

const (
	U RubikFace = iota + 1
	L
	F
	R
	B
	D
)

type RubikFaceTurn uint8

const (
	Clockwise RubikFaceTurn = iota + 1
	CounterClockwise
)

type Rubik struct {
	pos_p3 [8]uint8
	rot_p3 [8]uint8
	pos_p2 [12]uint8
	rot_p2 [12]uint8
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

func clockwiseU(cube Rubik) Rubik {
	var tmp uint8 = 0

	tmp = cube.pos_p3[0]
	cube.pos_p3[0] = cube.pos_p3[3]
	cube.pos_p3[3] = cube.pos_p3[2]
	cube.pos_p3[2] = cube.pos_p3[1]
	cube.pos_p3[0] = tmp

	tmp = cube.pos_p2[0]
	cube.pos_p2[0] = cube.pos_p2[3]
	cube.pos_p2[3] = cube.pos_p2[2]
	cube.pos_p2[2] = cube.pos_p2[1]
	cube.pos_p2[0] = tmp

	cube.rot_p2[cube.pos_p2[0]] = (cube.rot_p2[cube.pos_p2[0]] + 1) % 2
	cube.rot_p2[cube.pos_p2[1]] = (cube.rot_p2[cube.pos_p2[1]] + 1) % 2
	cube.rot_p2[cube.pos_p2[2]] = (cube.rot_p2[cube.pos_p2[2]] + 1) % 2
	cube.rot_p2[cube.pos_p2[3]] = (cube.rot_p2[cube.pos_p2[3]] + 1) % 2

	cube.rot_p3[cube.pos_p3[0]] = (cube.rot_p3[cube.pos_p3[0]] + 1) % 3
	cube.rot_p3[cube.pos_p3[1]] = (cube.rot_p3[cube.pos_p3[1]] + 1) % 3
	cube.rot_p3[cube.pos_p3[2]] = (cube.rot_p3[cube.pos_p3[2]] + 1) % 3
	cube.rot_p3[cube.pos_p3[3]] = (cube.rot_p3[cube.pos_p3[3]] + 1) % 3
	return cube
}

func counterClockwiseU(cube Rubik) Rubik {
	var tmp uint8 = 0

	tmp = cube.pos_p3[0]
	cube.pos_p3[0] = cube.pos_p3[1]
	cube.pos_p3[1] = cube.pos_p3[2]
	cube.pos_p3[2] = cube.pos_p3[3]
	cube.pos_p3[3] = tmp

	tmp = cube.pos_p2[0]
	cube.pos_p2[0] = cube.pos_p2[1]
	cube.pos_p2[1] = cube.pos_p2[2]
	cube.pos_p2[2] = cube.pos_p2[3]
	cube.pos_p2[3] = tmp

	cube.rot_p2[cube.pos_p2[0]] = (cube.rot_p2[cube.pos_p2[0]] - 1) % 2
	cube.rot_p2[cube.pos_p2[1]] = (cube.rot_p2[cube.pos_p2[1]] - 1) % 2
	cube.rot_p2[cube.pos_p2[2]] = (cube.rot_p2[cube.pos_p2[2]] - 1) % 2
	cube.rot_p2[cube.pos_p2[3]] = (cube.rot_p2[cube.pos_p2[3]] - 1) % 2

	cube.rot_p3[cube.pos_p3[0]] = (cube.rot_p3[cube.pos_p3[0]] - 1) % 3
	cube.rot_p3[cube.pos_p3[1]] = (cube.rot_p3[cube.pos_p3[1]] - 1) % 3
	cube.rot_p3[cube.pos_p3[2]] = (cube.rot_p3[cube.pos_p3[2]] - 1) % 3
	cube.rot_p3[cube.pos_p3[3]] = (cube.rot_p3[cube.pos_p3[3]] - 1) % 3
	return cube
}

func clockwiseL(cube Rubik) Rubik {
	return cube
}

func counterClockwiseL(cube Rubik) Rubik {
	return cube
}

func clockwiseF(cube Rubik) Rubik {
	return cube
}

func counterClockwiseF(cube Rubik) Rubik {
	return cube
}

func clockwiseR(cube Rubik) Rubik {
	return cube
}

func counterClockwiseR(cube Rubik) Rubik {
	return cube
}

func clockwiseB(cube Rubik) Rubik {
	return cube
}

func counterClockwiseB(cube Rubik) Rubik {
	return cube
}

func clockwiseD(cube Rubik) Rubik {
	return cube
}

func counterClockwiseD(cube Rubik) Rubik {
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
