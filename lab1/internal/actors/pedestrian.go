package actors

import (
	"github.com/petar-cvit/srsv.lab1/internal/utils"
)

type Pedestrian struct {
	Position          string
	register          chan string
	semaphore         chan int
	semaphoreLocation string
	crossingChan      chan *utils.CrossingMessage
	draw              chan string
}

func NewPedestrian(position string, register chan string, semaphore chan int,
	draw chan string, crossingChan chan *utils.CrossingMessage, semaphoreLocation string) *Pedestrian {
	return &Pedestrian{
		Position:          position,
		register:          register,
		semaphore:         semaphore,
		semaphoreLocation: semaphoreLocation,
		crossingChan:      crossingChan,
		draw:              draw,
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
