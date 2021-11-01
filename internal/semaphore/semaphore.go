package semaphore

type Semaphore struct {
	Current   int
	StateChan chan int
	Running   bool
}

func New(state int) *Semaphore {
	return &Semaphore{
		Current:   state,
		StateChan: make(chan int),
		Running:   false,
	}
}

func (s *Semaphore) Start() {
	s.Running = true
	for {
		select {
		case next := <-s.StateChan:
			s.Current = next
		}
	}
}
