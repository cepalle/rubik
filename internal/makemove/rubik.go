package makemove

import (
	"log"
)

type RubikFace uint8

const (
	U RubikFace = iota + 1
	F
	R
	B
	L
	D
)

type RubikFaceTurn uint8

const (
	Clockwise RubikFaceTurn = iota + 1
	CounterClockwise
)

type Rubik struct {
	PosP2  [12]uint8
	RotP2  [12]uint8
	PosFP3 [24]uint8
}

func InitRubik() Rubik {
	var rubik Rubik

	rubik.PosFP3 = [24]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	rubik.PosP2 = [12]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	return rubik
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

type RubikMovesWithName struct {
	Name string
	Move RubikMoves
	Rev  RubikMoves
}

const NbRubikMoves = 18

var AllRubikMovesWithName = [NbRubikMoves]RubikMovesWithName{
	RubikMovesWithName{
		Name: "U",
		Move: RubikMoves{
			U, Clockwise, 1,
		},
		Rev: RubikMoves{
			U, CounterClockwise, 1,
		},
	},
	RubikMovesWithName{
		Name: "U2",
		Move: RubikMoves{
			U, Clockwise, 2,
		},
		Rev: RubikMoves{
			U, Clockwise, 2,
		},
	},
	RubikMovesWithName{
		Name: "U'",
		Move: RubikMoves{
			U, CounterClockwise, 1,
		},
		Rev: RubikMoves{
			U, Clockwise, 1,
		},
	},
	RubikMovesWithName{
		Name: "D",
		Move: RubikMoves{
			D, Clockwise, 1,
		},
		Rev: RubikMoves{
			D, CounterClockwise, 1,
		},
	},
	RubikMovesWithName{
		Name: "D2",
		Move: RubikMoves{
			D, Clockwise, 2,
		},
		Rev: RubikMoves{
			D, Clockwise, 2,
		},
	},
	RubikMovesWithName{
		Name: "D'",
		Move: RubikMoves{
			D, CounterClockwise, 1,
		},
		Rev: RubikMoves{
			D, Clockwise, 1,
		},
	},
	RubikMovesWithName{
		Name: "B",
		Move: RubikMoves{
			B, Clockwise, 1,
		},
		Rev: RubikMoves{
			B, CounterClockwise, 1,
		},
	},
	RubikMovesWithName{
		Name: "B2",
		Move: RubikMoves{
			B, Clockwise, 2,
		},
		Rev: RubikMoves{
			B, Clockwise, 2,
		},
	},
	RubikMovesWithName{
		Name: "B'",
		Move: RubikMoves{
			B, CounterClockwise, 1,
		},
		Rev: RubikMoves{
			B, Clockwise, 1,
		},
	},
	RubikMovesWithName{
		Name: "R",
		Move: RubikMoves{
			R, Clockwise, 1,
		},
		Rev: RubikMoves{
			R, CounterClockwise, 1,
		},
	},
	RubikMovesWithName{
		Name: "R2",
		Move: RubikMoves{
			R, Clockwise, 2,
		},
		Rev: RubikMoves{
			R, Clockwise, 2,
		},
	},
	RubikMovesWithName{
		Name: "R'",
		Move: RubikMoves{
			R, CounterClockwise, 1,
		},
		Rev: RubikMoves{
			R, Clockwise, 1,
		},
	},
	RubikMovesWithName{
		Name: "F",
		Move: RubikMoves{
			F, Clockwise, 1,
		},
		Rev: RubikMoves{
			F, CounterClockwise, 1,
		},
	},
	RubikMovesWithName{
		Name: "F2",
		Move: RubikMoves{
			F, Clockwise, 2,
		},
		Rev: RubikMoves{
			F, Clockwise, 2,
		},
	},
	RubikMovesWithName{
		Name: "F'",
		Move: RubikMoves{
			F, CounterClockwise, 1,
		},
		Rev: RubikMoves{
			F, Clockwise, 1,
		},
	},
	RubikMovesWithName{
		Name: "L",
		Move: RubikMoves{
			L, Clockwise, 1,
		},
		Rev: RubikMoves{
			L, CounterClockwise, 1,
		},
	},
	RubikMovesWithName{
		Name: "L2",
		Move: RubikMoves{
			L, Clockwise, 2,
		},
		Rev: RubikMoves{
			L, Clockwise, 2,
		},
	},
	RubikMovesWithName{
		Name: "L'",
		Move: RubikMoves{
			L, CounterClockwise, 1,
		},
		Rev: RubikMoves{
			L, Clockwise, 1,
		},
	},
}

type moveFunction func(*Rubik) *Rubik

type dispatcher struct {
	move RubikMove
	fun  moveFunction
}

type poseSwap struct {
	ip2 [4]uint8
	ip3 [12]uint8
}

func clockwiseWithPose(ps poseSwap) moveFunction {
	return func(cube *Rubik) *Rubik {
		var tmp uint8 = 0

		tmp = cube.PosP2[ps.ip2[0]]
		cube.PosP2[ps.ip2[0]] = cube.PosP2[ps.ip2[3]]
		cube.PosP2[ps.ip2[3]] = cube.PosP2[ps.ip2[2]]
		cube.PosP2[ps.ip2[2]] = cube.PosP2[ps.ip2[1]]
		cube.PosP2[ps.ip2[1]] = tmp

		cube.RotP2[cube.PosP2[ps.ip2[0]]] = (cube.RotP2[cube.PosP2[ps.ip2[0]]] + 1) % 2
		cube.RotP2[cube.PosP2[ps.ip2[1]]] = (cube.RotP2[cube.PosP2[ps.ip2[1]]] + 1) % 2
		cube.RotP2[cube.PosP2[ps.ip2[2]]] = (cube.RotP2[cube.PosP2[ps.ip2[2]]] + 1) % 2
		cube.RotP2[cube.PosP2[ps.ip2[3]]] = (cube.RotP2[cube.PosP2[ps.ip2[3]]] + 1) % 2

		tmp = ps.ip3[0]
		ps.ip3[0] = ps.ip3[3]
		ps.ip3[3] = ps.ip3[2]
		ps.ip3[2] = ps.ip3[1]
		ps.ip3[1] = tmp

		tmp = ps.ip3[0+4]
		ps.ip3[0+4] = ps.ip3[3+4]
		ps.ip3[3+4] = ps.ip3[2+4]
		ps.ip3[2+4] = ps.ip3[1+4]
		ps.ip3[1+4] = tmp

		tmp = ps.ip3[0+8]
		ps.ip3[0+8] = ps.ip3[3+8]
		ps.ip3[3+8] = ps.ip3[2+8]
		ps.ip3[2+8] = ps.ip3[1+8]
		ps.ip3[1+8] = tmp

		return cube
	}
}

func counterClockwiseWithPose(ps poseSwap) moveFunction {
	return func(cube *Rubik) *Rubik {
		var tmp uint8 = 0

		tmp = cube.PosP2[ps.ip2[0]]
		cube.PosP2[ps.ip2[0]] = cube.PosP2[ps.ip2[1]]
		cube.PosP2[ps.ip2[1]] = cube.PosP2[ps.ip2[2]]
		cube.PosP2[ps.ip2[2]] = cube.PosP2[ps.ip2[3]]
		cube.PosP2[ps.ip2[3]] = tmp

		cube.RotP2[cube.PosP2[ps.ip2[0]]] = (cube.RotP2[cube.PosP2[ps.ip2[0]]] + 1) % 2
		cube.RotP2[cube.PosP2[ps.ip2[1]]] = (cube.RotP2[cube.PosP2[ps.ip2[1]]] + 1) % 2
		cube.RotP2[cube.PosP2[ps.ip2[2]]] = (cube.RotP2[cube.PosP2[ps.ip2[2]]] + 1) % 2
		cube.RotP2[cube.PosP2[ps.ip2[3]]] = (cube.RotP2[cube.PosP2[ps.ip2[3]]] + 1) % 2

		tmp = ps.ip3[0]
		ps.ip3[0] = ps.ip3[1]
		ps.ip3[1] = ps.ip3[2]
		ps.ip3[2] = ps.ip3[3]
		ps.ip3[3] = tmp

		tmp = ps.ip3[4]
		ps.ip3[4] = ps.ip3[5]
		ps.ip3[5] = ps.ip3[6]
		ps.ip3[6] = ps.ip3[7]
		ps.ip3[7] = tmp

		tmp = ps.ip3[8]
		ps.ip3[8] = ps.ip3[9]
		ps.ip3[9] = ps.ip3[10]
		ps.ip3[10] = ps.ip3[11]
		ps.ip3[11] = tmp

		return cube
	}
}

var poseSwapU = poseSwap{
	[4]uint8{
		0, 1, 2, 3,
	},
	[12]uint8{
		0, 1, 2, 3,
		11, 9, 7, 5,
		10, 8, 6, 4,
	},
}

var poseSwapF = poseSwap{
	[4]uint8{
		2, 5, 10, 6,
	},
	[12]uint8{
		6, 7, 15, 14,
		5, 3, 16, 20,
		21, 13, 2, 8,
	},
}

var poseSwapR = poseSwap{
	[4]uint8{
		1, 4, 9, 5,
	},
	[12]uint8{
		8, 9, 17, 16,
		7, 1, 18, 21,
		3, 10, 23, 15,
	},
}

var poseSwapB = poseSwap{
	[4]uint8{
		0, 7, 8, 4,
	},
	[12]uint8{
		10, 11, 19, 18,
		17, 1, 4, 22,
		23, 9, 0, 12,
	},
}

var poseSwapL = poseSwap{
	[4]uint8{
		3, 6, 11, 7,
	},
	[12]uint8{
		4, 5, 13, 12,
		11, 2, 14, 22,
		19, 0, 6, 20,
	},
}

var poseSwapD = poseSwap{
	[4]uint8{
		8, 11, 10, 9,
	},
	[12]uint8{
		20, 21, 23, 22,
		12, 14, 16, 18,
		13, 15, 17, 19,
	},
}

const dispatcherLen int = 12

var dispatcherTab = [dispatcherLen]dispatcher{
	dispatcher{
		move: RubikMove{
			face: U,
			turn: Clockwise,
		},
		fun: clockwiseWithPose(poseSwapU),
	},
	dispatcher{
		move: RubikMove{
			face: U,
			turn: CounterClockwise,
		},
		fun: counterClockwiseWithPose(poseSwapU),
	},
	dispatcher{
		move: RubikMove{
			face: L,
			turn: Clockwise,
		},
		fun: clockwiseWithPose(poseSwapL),
	},
	dispatcher{
		move: RubikMove{
			face: L,
			turn: CounterClockwise,
		},
		fun: counterClockwiseWithPose(poseSwapL),
	},
	dispatcher{
		move: RubikMove{
			face: F,
			turn: Clockwise,
		},
		fun: clockwiseWithPose(poseSwapF),
	},
	dispatcher{
		move: RubikMove{
			face: F,
			turn: CounterClockwise,
		},
		fun: counterClockwiseWithPose(poseSwapF),
	},
	dispatcher{
		move: RubikMove{
			face: R,
			turn: Clockwise,
		},
		fun: clockwiseWithPose(poseSwapR),
	},
	dispatcher{
		move: RubikMove{
			face: R,
			turn: CounterClockwise,
		},
		fun: counterClockwiseWithPose(poseSwapR),
	},
	dispatcher{
		move: RubikMove{
			face: B,
			turn: Clockwise,
		},
		fun: clockwiseWithPose(poseSwapB),
	},
	dispatcher{
		move: RubikMove{
			face: B,
			turn: CounterClockwise,
		},
		fun: counterClockwiseWithPose(poseSwapB),
	},
	dispatcher{
		move: RubikMove{
			face: D,
			turn: Clockwise,
		},
		fun: clockwiseWithPose(poseSwapD),
	},
	dispatcher{
		move: RubikMove{
			face: D,
			turn: CounterClockwise,
		},
		fun: counterClockwiseWithPose(poseSwapD),
	},
}

func (cube Rubik) DoMove(m RubikMoves) Rubik {
	for i := 0; i < dispatcherLen; i++ {
		if dispatcherTab[i].move.face == m.face && dispatcherTab[i].move.turn == m.turn {
			for j := uint8(0); j < m.nbTurn; j++ {
				dispatcherTab[i].fun(&cube)
			}
			return cube
		}
	}
	log.Fatal("You should not reach this code")
	return cube
}

func (cube *Rubik) DoMovePtr(m RubikMoves) *Rubik {
	for i := 0; i < dispatcherLen; i++ {
		if dispatcherTab[i].move.face == m.face && dispatcherTab[i].move.turn == m.turn {
			for j := uint8(0); j < m.nbTurn; j++ {
				dispatcherTab[i].fun(cube)
			}
			return cube
		}
	}
	log.Fatal("You should not reach this code")
	return nil
}

func (cube Rubik) DoMoves(m []RubikMoves) Rubik {
	for _, move := range m {
		cube = cube.DoMove(move)
	}
	return cube
}

func (cube Rubik) IsResolve() bool {
	var i uint8
	for i = 0; i < 12; i++ {
		if cube.RotP2[i] != 0 {
			return false
		}
	}
	for i = 0; i < 12; i++ {
		if cube.PosP2[i] != i {
			return false
		}
	}
	for i = 0; i < 24; i++ {
		if cube.PosFP3[i] != i {
			return false
		}
	}
	return true
}
