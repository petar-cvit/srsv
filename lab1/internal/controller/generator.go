package controller

import (
	"time"

	"lab2/lab1/internal/actors"
	"lab2/lab1/internal/logger"
	"lab2/lab1/internal/semaphore"
	"lab2/lab1/internal/utils"
)

type Generator struct {
	semaphores   map[string]*semaphore.Semaphore
	register     chan string
	crossingChan chan *utils.CrossingMessage
	draw         chan string
}

func NewGenerator(register chan string, semaphores map[string]*semaphore.Semaphore, draw chan string,
	crossingChan chan *utils.CrossingMessage) *Generator {
	return &Generator{
		register:     register,
		semaphores:   semaphores,
		crossingChan: crossingChan,
		draw:         draw,
	}
}

func (g *Generator) Start() {
	logger := logger.New()

	go func() {
		time.Sleep(10)

		for {
			time.Sleep(time.Second * 45)

			pedestrian := actors.NewPedestrian(
				utils.PedestrianWestToEastNorth,
				g.register,
				g.semaphores[utils.PedestrianNorthLeft].OutputChan,
				g.draw,
				g.crossingChan,
				utils.PedestrianNorthDraw,
				logger,
			)
			pedestrian.StartPedestrian()
		}
	}()

	go func() {
		time.Sleep(time.Second * 6)

		for {
			time.Sleep(time.Second * 50)

			pedestrian := actors.NewPedestrian(
				utils.PedestrianEastToWestNorth,
				g.register,
				g.semaphores[utils.PedestrianNorthRight].OutputChan,
				g.draw,
				g.crossingChan,
				utils.PedestrianNorthDraw,
				logger,
			)
			pedestrian.StartPedestrian()
		}
	}()

	go func() {
		time.Sleep(time.Second * 35)

		for {
			time.Sleep(time.Second * 62)

			pedestrian := actors.NewPedestrian(
				utils.PedestrianWestToEastSouth,
				g.register,
				g.semaphores[utils.PedestrianSouthLeft].OutputChan,
				g.draw,
				g.crossingChan,
				utils.PedestrianSouthDraw,
				logger,
			)
			pedestrian.StartPedestrian()
		}
	}()

	go func() {
		time.Sleep(time.Second * 70)

		for {
			time.Sleep(time.Second * 71)

			pedestrian := actors.NewPedestrian(
				utils.PedestrianEastToWestSouth,
				g.register,
				g.semaphores[utils.PedestrianSouthRight].OutputChan,
				g.draw,
				g.crossingChan,
				utils.PedestrianSouthDraw,
				logger,
			)
			pedestrian.StartPedestrian()
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 60)

			pedestrian := actors.NewPedestrian(
				utils.PedestrianNorthToSouthEast,
				g.register,
				g.semaphores[utils.PedestrianEastNorth].OutputChan,
				g.draw,
				g.crossingChan,
				utils.PedestrianEastDraw,
				logger,
			)
			pedestrian.StartPedestrian()
		}
	}()

	go func() {
		time.Sleep(time.Second * 25)

		for {
			time.Sleep(time.Second * 65)

			pedestrian := actors.NewPedestrian(
				utils.PedestrianNorthToSouthWest,
				g.register,
				g.semaphores[utils.PedestrianWestNorth].OutputChan,
				g.draw,
				g.crossingChan,
				utils.PedestrianWestDraw,
				logger,
			)
			pedestrian.StartPedestrian()
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 55)

			pedestrian := actors.NewPedestrian(
				utils.PedestrianSouthToNorthEast,
				g.register,
				g.semaphores[utils.PedestrianEastSouth].OutputChan,
				g.draw,
				g.crossingChan,
				utils.PedestrianEastDraw,
				logger,
			)
			pedestrian.StartPedestrian()
		}
	}()

	go func() {
		time.Sleep(time.Second * 55)

		for {
			time.Sleep(time.Second * 51)

			pedestrian := actors.NewPedestrian(
				utils.PedestrianSouthToNorthWest,
				g.register,
				g.semaphores[utils.PedestrianWestSouth].OutputChan,
				g.draw,
				g.crossingChan,
				utils.PedestrianWestDraw,
				logger,
			)
			pedestrian.StartPedestrian()
		}
	}()

	go func() {
		time.Sleep(time.Second * 15)

		for {
			time.Sleep(time.Second * 80)

			car := actors.NewCar(
				utils.StraightVerticalToNorth,
				g.register,
				g.semaphores[utils.StraightVerticalToNorth].OutputChan,
				g.draw,
				g.crossingChan,
			)
			car.StartCar()
		}
	}()

	go func() {
		time.Sleep(time.Second * 10)

		for {
			time.Sleep(time.Second * 90)

			car := actors.NewCar(
				utils.StraightVerticalToSouth,
				g.register,
				g.semaphores[utils.StraightVerticalToSouth].OutputChan,
				g.draw,
				g.crossingChan,
			)
			car.StartCar()
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)
		for {
			time.Sleep(time.Second * 40)

			car := actors.NewCar(
				utils.StraightHorizontalToWest,
				g.register,
				g.semaphores[utils.StraightHorizontalToWest].OutputChan,
				g.draw,
				g.crossingChan,
			)
			car.StartCar()
		}
	}()

	go func() {
		time.Sleep(time.Second * 10)

		for {
			time.Sleep(time.Second * 70)

			car := actors.NewCar(
				utils.StraightHorizontalToEast,
				g.register,
				g.semaphores[utils.StraightHorizontalToEast].OutputChan,
				g.draw,
				g.crossingChan,
			)
			car.StartCar()
		}
	}()
}
