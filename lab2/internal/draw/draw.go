package draw

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"srsv/internal/elevator"
	"srsv/internal/passenger"
	"srsv/internal/utils"
)

type Drawer struct {
	numberOfFloors int

	elevator   *elevator.Elevator
	passengers []*passenger.Passenger

	waitingPassengers map[int]map[string]*passenger.Passenger
	leftPassengers    map[int][]*passenger.Passenger
	drivingPassengers []*passenger.Passenger

	elevatorChan  chan bool
	passengerChan chan *passenger.Passenger
}

func New(elevator *elevator.Elevator, numberOfFloors int, passengers []*passenger.Passenger,
	elevatorChan chan bool, passengerChan chan *passenger.Passenger) *Drawer {
	return &Drawer{
		numberOfFloors:    numberOfFloors,
		elevator:          elevator,
		passengers:        passengers,
		waitingPassengers: make(map[int]map[string]*passenger.Passenger),
		leftPassengers:    make(map[int][]*passenger.Passenger),
		drivingPassengers: make([]*passenger.Passenger, 0, 6),
		elevatorChan:      elevatorChan,
		passengerChan:     passengerChan,
	}
}

func (d *Drawer) Start() {
	for {
		select {
		case <-d.elevatorChan:
			open := false
			for _, p := range d.drivingPassengers {
				if p.ToFloor == d.elevator.CurrentFloor {
					d.leftPassengers[d.elevator.CurrentFloor] = append(d.leftPassengers[d.elevator.CurrentFloor], p)
					d.drivingPassengers = deletePassenger(p, d.drivingPassengers)
					d.elevator.Space++
					open = true
				}
			}

			getOn := make([]*passenger.Passenger, 0)
			for _, p := range d.waitingPassengers[d.elevator.CurrentFloor] {
				if d.elevator.Space == 0 {
					break
				}

				if p.Name == "P" {
					fmt.Println()
				}

				if d.elevator.Direction != utils.Waiting && p.Direction != d.elevator.Direction {
					continue
				}

				getOn = append(getOn, p)
				d.elevator.Space--
				d.elevator.GateState = utils.Open
				open = true
			}

			sort.Slice(getOn, func(i, j int) bool {
				if d.elevator.Direction == utils.Up {
					return getOn[i].ToFloor < getOn[j].ToFloor
				}

				return getOn[i].ToFloor > getOn[j].ToFloor
			})

			for _, p := range getOn {
				delete(d.waitingPassengers[d.elevator.CurrentFloor], p.Name)
				d.drivingPassengers = append(d.drivingPassengers, p)
				//p.ElevatorChan <- p.ToFloor
			}

			if open {
				d.elevator.GateState = utils.Open
			} else {
				d.elevator.GateState = utils.Closed
			}

			d.draw()
		case p := <-d.passengerChan:
			if d.waitingPassengers[p.FromFloor] == nil {
				d.waitingPassengers[p.FromFloor] = make(map[string]*passenger.Passenger)
			}

			d.waitingPassengers[p.FromFloor][p.Name] = p
			d.draw()
		}
	}
}

func (d *Drawer) draw() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Println("smjer/vrata:", d.elevator.Direction, d.elevator.GateState)

	fmt.Println("Stajanja:   ========== Izašli")

	for i := d.numberOfFloors; i > 0; i-- {
		fmt.Println(fmt.Sprintf("%v:%v|%v|%v", i, d.printWaiting(i), d.printElevator(i), d.printLeft(i)))
		if i != 1 {
			fmt.Println(fmt.Sprintf("============|%v|", "        "))
		}
	}

	d.printMetadata()
}

func (d *Drawer) printWaiting(floor int) string {
	return printPassengers(utils.MapToSlicePassengers(d.waitingPassengers[floor]), 10)
}

func (d *Drawer) printElevator(floor int) string {
	if floor != d.elevator.CurrentFloor {
		return "        "
	}

	return "[" + printPassengers(d.drivingPassengers, 6) + "]"
}

func (d *Drawer) printLeft(floor int) string {
	return printPassengers(d.leftPassengers[floor], len(d.passengers))
}

func printPassengers(passengers []*passenger.Passenger, maxSpace int) string {
	out := strings.Builder{}

	for _, p := range passengers {
		out.Write([]byte(p.Name))
	}

	for i := 0; i < maxSpace-len(passengers); i++ {
		out.Write([]byte(" "))
	}

	return out.String()
}

func (d *Drawer) printMetadata() {
	fmt.Print("Putnici: ")
	for _, p := range d.passengers {
		fmt.Print(p.Name)
	}
	fmt.Println()

	fmt.Print("     od: ")
	for _, p := range d.passengers {
		fmt.Print(p.FromFloor)
	}
	fmt.Println()

	fmt.Print("     do: ")
	for _, p := range d.passengers {
		fmt.Print(p.ToFloor)
	}
	fmt.Println()

	fmt.Print("  smjer: ")
	for _, p := range d.passengers {
		fmt.Print(p.Direction)
	}
	fmt.Println()
	fmt.Println(d.elevator.TargetFloor)
}

func deletePassenger(passenger *passenger.Passenger, passengers []*passenger.Passenger) []*passenger.Passenger {
	for i, p := range passengers {
		if p.Name == passenger.Name {
			return append(passengers[:i], passengers[i+1:]...)
		}
	}

	return passengers
}
