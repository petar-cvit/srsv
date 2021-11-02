package actors

import (
	"lab2/internal/logger"
	"lab2/internal/utils"
	"time"
)

type Pedestrian struct {
	Position          string
	register          chan string
	semaphore         chan int
	semaphoreLocation string
	crossingChan      chan *utils.CrossingMessage
	draw              chan string
	logger            *logger.Logger
}

func NewPedestrian(position string, register chan string, semaphore chan int,
	draw chan string, crossingChan chan *utils.CrossingMessage, semaphoreLocation string, logger *logger.Logger) *Pedestrian {
	return &Pedestrian{
		Position:          position,
		register:          register,
		semaphore:         semaphore,
		semaphoreLocation: semaphoreLocation,
		crossingChan:      crossingChan,
		draw:              draw,
		logger:            logger,
	}
}

func (p *Pedestrian) StartPedestrian() {
	p.draw <- p.Position

	go func() {
		for {
			select {
			case semaphoreState := <-p.semaphore:
				if semaphoreState != utils.Green {
					continue
				}

				p.crossingChan <- &utils.CrossingMessage{
					Position: p.semaphoreLocation,
					Crossing: true,
					Car:      false,
				}

				for {
					select {
					case state := <-p.semaphore:
						if state == utils.Red {
							time.Sleep(time.Millisecond * 30)
							p.crossingChan <- &utils.CrossingMessage{
								Position: p.semaphoreLocation,
								Crossing: false,
								Car:      false,
							}

							return
						}
					}
				}
			}
		}
	}()
}
