package semaphore

import "lab2/internal/utils"

type Semaphore struct {
	Current   int
	StateChan chan int
	position  string
	drawChan  chan *utils.SemaphoreMessage
}

func New(state int, position string, drawChan chan *utils.SemaphoreMessage) *Semaphore {
	return &Semaphore{
		Current:   state,
		StateChan: make(chan int),
		position:  position,
		drawChan:  drawChan,
	}
}

func (s *Semaphore) Start() {
	for {
		select {
		case next := <-s.StateChan:
			s.Current = next
			s.drawChan <- &utils.SemaphoreMessage{
				Position: s.position,
				State:    next,
			}
		}
	}
}
