package semaphore

type Semaphore struct {
	Current   int
	StateChan chan int
}

func New(state int) *Semaphore {
	return &Semaphore{
		Current:   state,
		StateChan: make(chan int),
	}
}

func (s *Semaphore) Start() {
	for {
		select {
		case next := <-s.StateChan:
			s.Current = next
		}
	}
}
