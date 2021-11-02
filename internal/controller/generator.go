package controller

import (
	"time"

	"lab2/internal/actors"
	"lab2/internal/logger"
	"lab2/internal/semaphore"
	"lab2/internal/utils"
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
		for {
			time.Sleep(time.Second * 15)

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
		for {
			time.Sleep(time.Second * 12)

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
		for {
			time.Sleep(time.Second * 18)

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
		for {
			time.Sleep(time.Second * 21)

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

	//go func() {
	//	for {
	//		time.Sleep(time.Second * 15)
	//
	//		car := actors.NewCar(
	//			utils.StraightVerticalToNorth,
	//			g.register,
	//			g.semaphores[utils.StraightVertical].OutputChan,
	//			g.draw,
	//			g.crossingChan,
	//		)
	//		car.StartCar()
	//	}
	//}()

	//go func() {
	//	for {
	//		time.Sleep(time.Second * 10)
	//
	//		car := actors.NewCar(
	//			utils.StraightVerticalToSouth,
	//			g.register,
	//			g.semaphores[utils.StraightVertical].OutputChan,
	//			g.draw,
	//			g.crossingChan,
	//		)
	//		car.StartCar()
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		time.Sleep(time.Second * 5)
	//
	//		car := actors.NewCar(
	//			utils.StraightHorizontalToWest,
	//			g.register,
	//			g.semaphores[utils.StraightHorizontal].OutputChan,
	//			g.draw,
	//			g.crossingChan,
	//		)
	//		car.StartCar()
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		time.Sleep(time.Second * 10)
	//
	//		car := actors.NewCar(
	//			utils.StraightHorizontalToEast,
	//			g.register,
	//			g.semaphores[utils.StraightHorizontal].OutputChan,
	//			g.draw,
	//			g.crossingChan,
	//		)
	//		car.StartCar()
	//	}
	//}()
}
