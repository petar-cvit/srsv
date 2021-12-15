package main

import (
	"os"
	"strconv"

	"srsv/internal"
	"srsv/internal/draw"
	"srsv/internal/elevator"
	"srsv/internal/generate"
	"srsv/internal/passenger"
)

func main() {
	traffic := int64(2)
	if len(os.Args) > 1 {
		traffic, _ = strconv.ParseInt(os.Args[1], 10, 64)
	}

	numOfFloors := 5

	passengersToDrawChan := make(chan *passenger.Passenger)
	elevatorToDrawChan := make(chan bool)
	passengersToElevatorChan := make(chan int)
	passengersEnterChan := make(chan int)

	passengers := generate.GeneratePassangers(15, numOfFloors, passengersToElevatorChan)
	elevator := elevator.New(elevatorToDrawChan, passengersToElevatorChan, passengersEnterChan)
	drawer := draw.New(elevator, numOfFloors, passengers, elevatorToDrawChan, passengersToDrawChan)

	go drawer.Start()
	go elevator.Start()

	go internal.Simulate(passengers, passengersToDrawChan, traffic)

	// invoke drawing
	elevatorToDrawChan <- true

	for {
	}
}
