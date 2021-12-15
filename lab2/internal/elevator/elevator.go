package elevator

import (
	"time"

	"srsv/internal/utils"
)

type Elevator struct {
	Direction string
	GateState string

	CurrentFloor int
	TargetFloor  int

	Space int

	passengerEnterChan chan int
	passengerChan      chan int
	drawChan           chan bool

	operator int
}

func New(drawChan chan bool, passengerChan chan int, passengerEnterChan chan int) *Elevator {
	return &Elevator{
		Direction:          utils.Waiting,
		GateState:          utils.Closed,
		operator:           0,
		CurrentFloor:       1,
		TargetFloor:        1,
		Space:              6,
		passengerEnterChan: passengerEnterChan,
		passengerChan:      passengerChan,
		drawChan:           drawChan,
	}
}

func (e *Elevator) Start() {
	for {
		select {
		case called := <-e.passengerChan:
			if called < e.CurrentFloor {
				e.Direction = utils.Down
				e.GateState = utils.Closed
				e.TargetFloor = called

				e.operator = -1
				e.move()
			} else if called > e.CurrentFloor {
				e.Direction = utils.Up
				e.GateState = utils.Closed
				e.TargetFloor = called

				e.operator = 1
				e.move()
			} else {
				e.Direction = utils.Waiting
				e.GateState = utils.Open
			}
		}
	}
}

func (e *Elevator) move() {
	for e.TargetFloor != e.CurrentFloor {
		e.CurrentFloor += e.operator
		if e.CurrentFloor == e.TargetFloor {
			e.Direction = utils.Waiting
		}
		e.drawChan <- true
		time.Sleep(time.Millisecond * 1500)
	}
}
