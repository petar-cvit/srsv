package actors

import (
	"fmt"

	"lab2/internal/utils"
)

type Car struct {
	Position  string
	register  chan string
	semaphore chan string
	draw      chan string
}

func New(position string, register chan string, semaphore chan string, draw chan string) *Car {
	return &Car{
		Position:  position,
		register:  register,
		semaphore: semaphore,
		draw:      draw,
	}
}

func (c *Car) Start() {
	c.draw <- c.Position

	fmt.Println("started car")

	go func() {
		for {
			select {
			case <-c.semaphore:
				c.draw <- utils.CarMessage(c.Position)
				return
			}
		}
	}()
}
