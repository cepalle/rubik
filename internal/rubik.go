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

func clockwiseWithPose(cube Rubik, ip2 [4] uint8, ip3 [4] uint8) Rubik {
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

func counterClockwiseWithPose(cube Rubik, ip2 [4] uint8, ip3 [4] uint8) Rubik {
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

	cube.rot_p3[cube.pos_p3[ip3[0]]] = (cube.rot_p3[cube.pos_p3[ip3[0]]] - 1) % 3
	cube.rot_p3[cube.pos_p3[ip3[1]]] = (cube.rot_p3[cube.pos_p3[ip3[1]]] - 1) % 3
	cube.rot_p3[cube.pos_p3[ip3[2]]] = (cube.rot_p3[cube.pos_p3[ip3[2]]] - 1) % 3
	cube.rot_p3[cube.pos_p3[ip3[3]]] = (cube.rot_p3[cube.pos_p3[ip3[3]]] - 1) % 3

	cube.rot_p2[cube.pos_p2[ip2[0]]] = (cube.rot_p2[cube.pos_p2[ip2[0]]] - 1) % 2
	cube.rot_p2[cube.pos_p2[ip2[1]]] = (cube.rot_p2[cube.pos_p2[ip2[1]]] - 1) % 2
	cube.rot_p2[cube.pos_p2[ip2[2]]] = (cube.rot_p2[cube.pos_p2[ip2[2]]] - 1) % 2
	cube.rot_p2[cube.pos_p2[ip2[3]]] = (cube.rot_p2[cube.pos_p2[ip2[3]]] - 1) % 2
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
					0, 0, 0, 0,
				}, [4]uint8{
					0, 0, 0, 0,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: L,
			turn: CounterClockwise,
		},
		fun: func(r Rubik) Rubik {
			return clockwiseWithPose(r,
				[4]uint8{
					0, 0, 0, 0,
				}, [4]uint8{
					0, 0, 0, 0,
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
					0, 0, 0, 0,
				}, [4]uint8{
					0, 0, 0, 0,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: F,
			turn: CounterClockwise,
		},
		fun: func(r Rubik) Rubik {
			return clockwiseWithPose(r,
				[4]uint8{
					0, 0, 0, 0,
				}, [4]uint8{
					0, 0, 0, 0,
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
					0, 0, 0, 0,
				}, [4]uint8{
					0, 0, 0, 0,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: R,
			turn: CounterClockwise,
		},
		fun: func(r Rubik) Rubik {
			return clockwiseWithPose(r,
				[4]uint8{
					0, 0, 0, 0,
				}, [4]uint8{
					0, 0, 0, 0,
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
					0, 0, 0, 0,
				}, [4]uint8{
					0, 0, 0, 0,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: B,
			turn: CounterClockwise,
		},
		fun: func(r Rubik) Rubik {
			return clockwiseWithPose(r,
				[4]uint8{
					0, 0, 0, 0,
				}, [4]uint8{
					0, 0, 0, 0,
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
					0, 0, 0, 0,
				}, [4]uint8{
					0, 0, 0, 0,
				})
		},
	},
	dispatcher{
		move: RubikMove{
			face: D,
			turn: CounterClockwise,
		},
		fun: func(r Rubik) Rubik {
			return clockwiseWithPose(r,
				[4]uint8{
					0, 0, 0, 0,
				}, [4]uint8{
					0, 0, 0, 0,
				})
		},
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
