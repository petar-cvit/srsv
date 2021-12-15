package passenger

type Passenger struct {
	Name      string
	ToFloor   int
	FromFloor int
	Direction string

	ElevatorChan chan int
}

func New(name string, to, from int, direction string, elevatorChan chan int) *Passenger {
	return &Passenger{
		Name:         name,
		ToFloor:      to,
		FromFloor:    from,
		Direction:    direction,
		ElevatorChan: elevatorChan,
	}
}

func (p *Passenger) Start() {
	p.ElevatorChan <- p.FromFloor
}
