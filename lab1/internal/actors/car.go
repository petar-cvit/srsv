package actors

import (
	"lab2/lab1/internal/utils"
)

type Car struct {
	Position     string
	register     chan string
	semaphore    chan int
	crossingChan chan *utils.CrossingMessage
	draw         chan string
}

func NewCar(position string, register chan string, semaphore chan int,
	draw chan string, crossingChan chan *utils.CrossingMessage) *Car {
	return &Car{
		Position:     position,
		register:     register,
		semaphore:    semaphore,
		crossingChan: crossingChan,
		draw:         draw,
	}
}

func (c *Car) StartCar() {
	c.draw <- c.Position

	go func() {
		for {
			select {
			case semaphoreState := <-c.semaphore:
				if semaphoreState != utils.Green {
					continue
				}

				go func() {
					c.crossingChan <- &utils.CrossingMessage{
						Position: c.Position,
						Crossing: true,
						Car:      true,
					}
				}()

				for {
					select {
					case state := <-c.semaphore:
						if state == utils.Red {
							c.crossingChan <- &utils.CrossingMessage{
								Position: c.Position,
								Crossing: false,
								Car:      true,
							}

							return
						}
					}
				}
			}
		}
	}()
}
