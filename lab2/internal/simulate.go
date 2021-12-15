package internal

import (
	"time"

	"srsv/internal/passenger"
)

func Simulate(passengers []*passenger.Passenger, passengersToDrawChan chan *passenger.Passenger) {
	for _, p := range passengers {
		time.Sleep(time.Millisecond * 400)
		passengersToDrawChan <- p
		p.Start()
	}
}
