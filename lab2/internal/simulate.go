package internal

import (
	"fmt"
	"time"

	"srsv/internal/passenger"
)

func Simulate(passengers []*passenger.Passenger, passengersToDrawChan chan *passenger.Passenger, traffic int64) {
	fmt.Println(traffic)
	fmt.Println((time.Millisecond * 500) * time.Duration(traffic))
	for _, p := range passengers {
		time.Sleep((time.Millisecond * 500) * time.Duration(traffic))
		passengersToDrawChan <- p
		p.Start()
	}
}
