package controller

import (
	"time"

	"lab2/internal/draw"
	"lab2/internal/semaphore"
	"lab2/internal/utils"
)

type Controller struct {
	Semaphores map[string]*semaphore.Semaphore
	drawer     *draw.Drawer
}

func New() {
	drawer := draw.New()
	go drawer.Start()

	semaphores := createSemaphores(drawer.SemaphoreChan)
	spinSemaphores(semaphores)

	generator := NewGenerator(nil, semaphores, drawer.WaitingChan, drawer.CrossingChan)
	generator.Start()

	go func() {
		t := 0

		for {
			time.Sleep(time.Second * 5)
			t++

			semaphores[utils.StraightHorizontal].StateChan <- utils.Green
			semaphores[utils.StraightVertical].StateChan <- utils.Red

			semaphores[utils.WestRight].StateChan <- utils.Red
			semaphores[utils.SouthRight].StateChan <- utils.Red

			semaphores[utils.WestLeft].StateChan <- utils.Red
			semaphores[utils.SouthLeft].StateChan <- utils.Red

			semaphores[utils.PedestrianEast].StateChan <- utils.Red
			semaphores[utils.PedestrianWest].StateChan <- utils.Red

			// north pedestrian
			semaphores[utils.PedestrianNorthLeft].StateChan <- utils.Green
			semaphores[utils.PedestrianNorthRight].StateChan <- utils.Green

			// south pedestrian
			semaphores[utils.PedestrianSouthLeft].StateChan <- utils.Green
			semaphores[utils.PedestrianSouthRight].StateChan <- utils.Green

			time.Sleep(time.Second * 5)
			t++

			semaphores[utils.StraightHorizontal].StateChan <- utils.Red
			semaphores[utils.StraightVertical].StateChan <- utils.Red

			semaphores[utils.WestRight].StateChan <- utils.Red
			semaphores[utils.SouthRight].StateChan <- utils.Red

			semaphores[utils.WestLeft].StateChan <- utils.Green
			semaphores[utils.SouthLeft].StateChan <- utils.Green

			semaphores[utils.PedestrianEast].StateChan <- utils.Red
			semaphores[utils.PedestrianWest].StateChan <- utils.Red

			// north pedestrian
			semaphores[utils.PedestrianNorthLeft].StateChan <- utils.Red
			semaphores[utils.PedestrianNorthRight].StateChan <- utils.Red

			// south pedestrian
			semaphores[utils.PedestrianSouthLeft].StateChan <- utils.Red
			semaphores[utils.PedestrianSouthRight].StateChan <- utils.Red

			time.Sleep(time.Second * 5)
			t++

			semaphores[utils.StraightHorizontal].StateChan <- utils.Red
			semaphores[utils.StraightVertical].StateChan <- utils.Green

			semaphores[utils.WestRight].StateChan <- utils.Red
			semaphores[utils.SouthRight].StateChan <- utils.Red

			semaphores[utils.WestLeft].StateChan <- utils.Red
			semaphores[utils.SouthLeft].StateChan <- utils.Red

			semaphores[utils.PedestrianEast].StateChan <- utils.Green
			semaphores[utils.PedestrianWest].StateChan <- utils.Green

			// north pedestrian
			semaphores[utils.PedestrianNorthLeft].StateChan <- utils.Red
			semaphores[utils.PedestrianNorthRight].StateChan <- utils.Red

			// south pedestrian
			semaphores[utils.PedestrianSouthLeft].StateChan <- utils.Red
			semaphores[utils.PedestrianSouthRight].StateChan <- utils.Red

			time.Sleep(time.Second * 5)
			t++

			semaphores[utils.StraightHorizontal].StateChan <- utils.Red
			semaphores[utils.StraightVertical].StateChan <- utils.Red

			semaphores[utils.WestRight].StateChan <- utils.Green
			semaphores[utils.SouthRight].StateChan <- utils.Green

			semaphores[utils.WestLeft].StateChan <- utils.Red
			semaphores[utils.SouthLeft].StateChan <- utils.Red

			semaphores[utils.PedestrianEast].StateChan <- utils.Red
			semaphores[utils.PedestrianWest].StateChan <- utils.Red

			// north pedestrian
			semaphores[utils.PedestrianNorthLeft].StateChan <- utils.Red
			semaphores[utils.PedestrianNorthRight].StateChan <- utils.Red

			// south pedestrian
			semaphores[utils.PedestrianSouthLeft].StateChan <- utils.Red
			semaphores[utils.PedestrianSouthRight].StateChan <- utils.Red
		}
	}()
}

func createSemaphores(drawChan chan *utils.SemaphoreMessage) map[string]*semaphore.Semaphore {
	semaphores := make(map[string]*semaphore.Semaphore)
	semaphores[utils.StraightHorizontal] = semaphore.New(utils.Red, utils.StraightHorizontal, drawChan)
	semaphores[utils.StraightVertical] = semaphore.New(utils.Red, utils.StraightVertical, drawChan)

	semaphores[utils.WestRight] = semaphore.New(utils.Red, utils.WestRight, drawChan)
	semaphores[utils.SouthRight] = semaphore.New(utils.Red, utils.SouthRight, drawChan)

	semaphores[utils.WestLeft] = semaphore.New(utils.Red, utils.WestLeft, drawChan)
	semaphores[utils.SouthLeft] = semaphore.New(utils.Red, utils.SouthLeft, drawChan)

	semaphores[utils.PedestrianEast] = semaphore.New(utils.Red, utils.PedestrianEast, drawChan)
	semaphores[utils.PedestrianWest] = semaphore.New(utils.Red, utils.PedestrianWest, drawChan)

	semaphores[utils.PedestrianNorthLeft] = semaphore.New(utils.Red, utils.PedestrianNorthLeft, drawChan)
	semaphores[utils.PedestrianNorthRight] = semaphore.New(utils.Red, utils.PedestrianNorthRight, drawChan)
	semaphores[utils.PedestrianSouthLeft] = semaphore.New(utils.Red, utils.PedestrianSouthLeft, drawChan)
	semaphores[utils.PedestrianSouthRight] = semaphore.New(utils.Red, utils.PedestrianSouthRight, drawChan)

	return semaphores
}

func spinSemaphores(semaphores map[string]*semaphore.Semaphore) {
	go semaphores[utils.StraightHorizontal].Start()
	go semaphores[utils.StraightVertical].Start()
	go semaphores[utils.WestRight].Start()
	go semaphores[utils.SouthRight].Start()
	go semaphores[utils.WestLeft].Start()
	go semaphores[utils.SouthLeft].Start()

	go semaphores[utils.PedestrianEast].Start()
	go semaphores[utils.PedestrianWest].Start()

	go semaphores[utils.PedestrianNorthRight].Start()
	go semaphores[utils.PedestrianNorthLeft].Start()
	go semaphores[utils.PedestrianSouthRight].Start()
	go semaphores[utils.PedestrianSouthLeft].Start()
}
