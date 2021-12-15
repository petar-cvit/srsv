package main

import (
	"srsv/internal"
	"srsv/internal/draw"
	"srsv/internal/elevator"
	"srsv/internal/generate"
	"srsv/internal/passenger"
)

func main() {
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

	go internal.Simulate(passengers, passengersToDrawChan)

	// invoke drawing
	elevatorToDrawChan <- true

	for {
	}
}
