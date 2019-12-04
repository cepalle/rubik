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

type RubikMoves struct {
	face   RubikFace
	turn   RubikFaceTurn
	nbTurn uint8
}

const NbRubikMoves = 18

var AllRubikMoves = [NbRubikMoves]RubikMoves{
	RubikMoves{U, Clockwise, 1},
	RubikMoves{L, Clockwise, 1},
	RubikMoves{F, Clockwise, 1},
	RubikMoves{R, Clockwise, 1},
	RubikMoves{B, Clockwise, 1},
	RubikMoves{D, Clockwise, 1},

	RubikMoves{U, CounterClockwise, 1},
	RubikMoves{L, CounterClockwise, 1},
	RubikMoves{F, CounterClockwise, 1},
	RubikMoves{R, CounterClockwise, 1},
	RubikMoves{B, CounterClockwise, 1},
	RubikMoves{D, CounterClockwise, 1},

	RubikMoves{U, Clockwise, 2},
	RubikMoves{L, Clockwise, 2},
	RubikMoves{F, Clockwise, 2},
	RubikMoves{R, Clockwise, 2},
	RubikMoves{B, Clockwise, 2},
	RubikMoves{D, Clockwise, 2},
}

type moveFunction func(Rubik) Rubik

type dispatcher struct {
	move RubikMove
	fun  moveFunction
}

func clockwiseWithPose(cube Rubik, ip2 [4]uint8, ip3 [4]uint8) Rubik {
	var tmp uint8 = 0

	tmp = cube.pos_p3[ip3[0]]
	cube.pos_p3[ip3[0]] = cube.pos_p3[ip3[3]]
	cube.pos_p3[ip3[3]] = cube.pos_p3[ip3[2]]
	cube.pos_p3[ip3[2]] = cube.pos_p3[ip3[1]]
	cube.pos_p3[ip3[1]] = tmp

	tmp = cube.pos_p2[ip2[0]]
	cube.pos_p2[ip2[0]] = cube.pos_p2[ip2[3]]
	cube.pos_p2[ip2[3]] = cube.pos_p2[ip2[2]]
	cube.pos_p2[ip2[2]] = cube.pos_p2[ip2[1]]
	cube.pos_p2[ip2[1]] = tmp

	cube.rot_p3[cube.pos_p3[ip3[0]]] = (cube.rot_p3[cube.pos_p3[ip3[0]]] + 1) % 3
	cube.rot_p3[cube.pos_p3[ip3[1]]] = (cube.rot_p3[cube.pos_p3[ip3[1]]] + 1) % 3
	cube.rot_p3[cube.pos_p3[ip3[2]]] = (cube.rot_p3[cube.pos_p3[ip3[2]]] + 1) % 3
	cube.rot_p3[cube.pos_p3[ip3[3]]] = (cube.rot_p3[cube.pos_p3[ip3[3]]] + 1) % 3

	cube.rot_p2[cube.pos_p2[ip2[0]]] = (cube.rot_p2[cube.pos_p2[ip2[0]]] + 1) % 2
	cube.rot_p2[cube.pos_p2[ip2[1]]] = (cube.rot_p2[cube.pos_p2[ip2[1]]] + 1) % 2
	cube.rot_p2[cube.pos_p2[ip2[2]]] = (cube.rot_p2[cube.pos_p2[ip2[2]]] + 1) % 2
	cube.rot_p2[cube.pos_p2[ip2[3]]] = (cube.rot_p2[cube.pos_p2[ip2[3]]] + 1) % 2
	return cube
}

func counterClockwiseWithPose(cube Rubik, ip2 [4]uint8, ip3 [4]uint8) Rubik {
	var tmp uint8 = 0

	tmp = cube.pos_p3[ip3[0]]
	cube.pos_p3[ip3[0]] = cube.pos_p3[ip3[1]]
	cube.pos_p3[ip3[1]] = cube.pos_p3[ip3[2]]
	cube.pos_p3[ip3[2]] = cube.pos_p3[ip3[3]]
	cube.pos_p3[ip3[3]] = tmp

	tmp = cube.pos_p2[ip2[0]]
	cube.pos_p2[ip2[0]] = cube.pos_p2[ip2[1]]
	cube.pos_p2[ip2[1]] = cube.pos_p2[ip2[2]]
	cube.pos_p2[ip2[2]] = cube.pos_p2[ip2[3]]
	cube.pos_p2[ip2[3]] = tmp

	cube.rot_p3[cube.pos_p3[ip3[0]]] = (cube.rot_p3[cube.pos_p3[ip3[0]]] + 2) % 3
	cube.rot_p3[cube.pos_p3[ip3[1]]] = (cube.rot_p3[cube.pos_p3[ip3[1]]] + 2) % 3
	cube.rot_p3[cube.pos_p3[ip3[2]]] = (cube.rot_p3[cube.pos_p3[ip3[2]]] + 2) % 3
	cube.rot_p3[cube.pos_p3[ip3[3]]] = (cube.rot_p3[cube.pos_p3[ip3[3]]] + 2) % 3

	cube.rot_p2[cube.pos_p2[ip2[0]]] = (cube.rot_p2[cube.pos_p2[ip2[0]]] + 1) % 2
	cube.rot_p2[cube.pos_p2[ip2[1]]] = (cube.rot_p2[cube.pos_p2[ip2[1]]] + 1) % 2
	cube.rot_p2[cube.pos_p2[ip2[2]]] = (cube.rot_p2[cube.pos_p2[ip2[2]]] + 1) % 2
	cube.rot_p2[cube.pos_p2[ip2[3]]] = (cube.rot_p2[cube.pos_p2[ip2[3]]] + 1) % 2
	return cube
}

const dispatcherLen = 12

var dispatcherTab = [dispatcherLen]dispatcher{
	dispatcher{
		move: RubikMove{
			face: U,
			turn: Clockwise,
		},
		fun: func(r Rubik) Rubik {
			return clockwiseWithPose(r,
				[4]uint8{
					0, 1, 2, 3,
				}, [4]uint8{
					0, 1, 2, 3,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: U,
			turn: CounterClockwise,
		},
		fun: func(r Rubik) Rubik {
			return counterClockwiseWithPose(r,
				[4]uint8{
					0, 1, 2, 3,
				}, [4]uint8{
					0, 1, 2, 3,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: L,
			turn: Clockwise,
		},
		fun: func(r Rubik) Rubik {
			return clockwiseWithPose(r,
				[4]uint8{
					3, 6, 11, 7,
				}, [4]uint8{
					2, 6, 7, 3,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: L,
			turn: CounterClockwise,
		},
		fun: func(r Rubik) Rubik {
			return counterClockwiseWithPose(r,
				[4]uint8{
					3, 6, 11, 7,
				}, [4]uint8{
					2, 6, 7, 3,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: F,
			turn: Clockwise,
		},
		fun: func(r Rubik) Rubik {
			return clockwiseWithPose(r,
				[4]uint8{
					2, 5, 10, 6,
				}, [4]uint8{
					1, 5, 6, 2,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: F,
			turn: CounterClockwise,
		},
		fun: func(r Rubik) Rubik {
			return counterClockwiseWithPose(r,
				[4]uint8{
					2, 5, 10, 6,
				}, [4]uint8{
					1, 5, 6, 2,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: R,
			turn: Clockwise,
		},
		fun: func(r Rubik) Rubik {
			return clockwiseWithPose(r,
				[4]uint8{
					1, 4, 9, 5,
				}, [4]uint8{
					1, 0, 4, 5,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: R,
			turn: CounterClockwise,
		},
		fun: func(r Rubik) Rubik {
			return counterClockwiseWithPose(r,
				[4]uint8{
					1, 4, 9, 5,
				}, [4]uint8{
					1, 0, 4, 5,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: B,
			turn: Clockwise,
		},
		fun: func(r Rubik) Rubik {
			return clockwiseWithPose(r,
				[4]uint8{
					0, 7, 8, 4,
				}, [4]uint8{
					0, 3, 7, 4,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: B,
			turn: CounterClockwise,
		},
		fun: func(r Rubik) Rubik {
			return counterClockwiseWithPose(r,
				[4]uint8{
					0, 7, 8, 4,
				}, [4]uint8{
					0, 3, 7, 4,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: D,
			turn: Clockwise,
		},
		fun: func(r Rubik) Rubik {
			return clockwiseWithPose(r,
				[4]uint8{
					10, 9, 8, 11,
				}, [4]uint8{
					6, 5, 4, 7,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: D,
			turn: CounterClockwise,
		},
		fun: func(r Rubik) Rubik {
			return counterClockwiseWithPose(r,
				[4]uint8{
					10, 9, 8, 11,
				}, [4]uint8{
					6, 5, 4, 7,
				})
		},
	},
}

func (cube Rubik) Move(m RubikMoves) Rubik {
	for i := 0; i < dispatcherLen; i++ {
		if dispatcherTab[i].move.face == m.face && dispatcherTab[i].move.turn == m.turn {
			for j := uint8(0); j < m.nbTurn; j++ {
				cube = dispatcherTab[j].fun(cube)
			}
			return cube
		}
	}
	return cube
}

func (cube Rubik) IsResolve() bool {
	var i uint8
	for i = 0; i < 12; i++ {
		if cube.rot_p2[i] != 0 {
			return false
		}
	}
	for i = 0; i < 8; i++ {
		if cube.rot_p3[i] != 0 {
			return false
		}
	}
	for i = 0; i < 12; i++ {
		if cube.pos_p2[i] != i {
			return false
		}
	}
	for i = 0; i < 8; i++ {
		if cube.pos_p3[i] != i {
			return false
		}
	}
	return true
}
