package controller

import (
	"fmt"
	"lab2/internal/actors"
	"lab2/internal/utils"
	"time"
)

type Generator struct {
	register  chan string
	semaphore chan string
	draw      chan string
}

func NewGenerator(register chan string, semaphore chan string, draw chan string) *Generator {
	return &Generator{
		register:  register,
		semaphore: semaphore,
		draw:      draw,
	}
}

func (g *Generator) Start() {
	go func() {
		fmt.Println("generator start")
		time.Sleep(time.Second * 5)

		car := actors.New(utils.StraightVerticalToNorth, g.register, g.semaphore, g.draw)
		car.Start()

		time.Sleep(time.Second * 5)

		car = actors.New(utils.StraightVerticalToSouth, g.register, g.semaphore, g.draw)
		car.Start()
	}()
}
