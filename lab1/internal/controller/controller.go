package controller

import (
	"time"

	"github.com/petar-cvit/srsv.lab1/internal/draw"
	"github.com/petar-cvit/srsv.lab1/internal/semaphore"
	"github.com/petar-cvit/srsv.lab1/internal/utils"
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
			t++

			semaphores[utils.StraightVerticalToSouth].StateChan <- utils.Red
			semaphores[utils.StraightVerticalToNorth].StateChan <- utils.Red
			semaphores[utils.StraightHorizontalToEast].StateChan <- utils.Green
			semaphores[utils.StraightHorizontalToWest].StateChan <- utils.Green

			semaphores[utils.WestRight].StateChan <- utils.Red
			semaphores[utils.SouthRight].StateChan <- utils.Red

			semaphores[utils.WestLeft].StateChan <- utils.Red
			semaphores[utils.SouthLeft].StateChan <- utils.Red

			// north pedestrian
			semaphores[utils.PedestrianNorthLeft].StateChan <- utils.Green
			semaphores[utils.PedestrianNorthRight].StateChan <- utils.Green

			// south pedestrian
			semaphores[utils.PedestrianSouthLeft].StateChan <- utils.Green
			semaphores[utils.PedestrianSouthRight].StateChan <- utils.Green

			// east pedestrians
			semaphores[utils.PedestrianEastNorth].StateChan <- utils.Red
			semaphores[utils.PedestrianEastSouth].StateChan <- utils.Red

			// west pedestrians
			semaphores[utils.PedestrianWestNorth].StateChan <- utils.Red
			semaphores[utils.PedestrianWestSouth].StateChan <- utils.Red

			time.Sleep(time.Second * 10)
			t++

			semaphores[utils.StraightVerticalToSouth].StateChan <- utils.Red
			semaphores[utils.StraightVerticalToNorth].StateChan <- utils.Red
			semaphores[utils.StraightHorizontalToEast].StateChan <- utils.Red
			semaphores[utils.StraightHorizontalToWest].StateChan <- utils.Red

			semaphores[utils.WestRight].StateChan <- utils.Red
			semaphores[utils.SouthRight].StateChan <- utils.Red

			semaphores[utils.WestLeft].StateChan <- utils.Green
			semaphores[utils.SouthLeft].StateChan <- utils.Green

			// north pedestrian
			semaphores[utils.PedestrianNorthLeft].StateChan <- utils.Red
			semaphores[utils.PedestrianNorthRight].StateChan <- utils.Red

			// south pedestrian
			semaphores[utils.PedestrianSouthLeft].StateChan <- utils.Red
			semaphores[utils.PedestrianSouthRight].StateChan <- utils.Red

			// east pedestrians
			semaphores[utils.PedestrianEastNorth].StateChan <- utils.Red
			semaphores[utils.PedestrianEastSouth].StateChan <- utils.Red

			// west pedestrians
			semaphores[utils.PedestrianWestNorth].StateChan <- utils.Red
			semaphores[utils.PedestrianWestSouth].StateChan <- utils.Red

			time.Sleep(time.Second * 10)
			t++

			semaphores[utils.StraightVerticalToSouth].StateChan <- utils.Green
			semaphores[utils.StraightVerticalToNorth].StateChan <- utils.Green
			semaphores[utils.StraightHorizontalToEast].StateChan <- utils.Red
			semaphores[utils.StraightHorizontalToWest].StateChan <- utils.Red

			semaphores[utils.WestRight].StateChan <- utils.Red
			semaphores[utils.SouthRight].StateChan <- utils.Red

			semaphores[utils.WestLeft].StateChan <- utils.Red
			semaphores[utils.SouthLeft].StateChan <- utils.Red

			// north pedestrian
			semaphores[utils.PedestrianNorthLeft].StateChan <- utils.Red
			semaphores[utils.PedestrianNorthRight].StateChan <- utils.Red

			// south pedestrian
			semaphores[utils.PedestrianSouthLeft].StateChan <- utils.Red
			semaphores[utils.PedestrianSouthRight].StateChan <- utils.Red

			// east pedestrians
			semaphores[utils.PedestrianEastNorth].StateChan <- utils.Green
			semaphores[utils.PedestrianEastSouth].StateChan <- utils.Green

			// west pedestrians
			semaphores[utils.PedestrianWestNorth].StateChan <- utils.Green
			semaphores[utils.PedestrianWestSouth].StateChan <- utils.Green

			time.Sleep(time.Second * 10)
			t++

			semaphores[utils.StraightVerticalToSouth].StateChan <- utils.Red
			semaphores[utils.StraightVerticalToNorth].StateChan <- utils.Red
			semaphores[utils.StraightHorizontalToEast].StateChan <- utils.Red
			semaphores[utils.StraightHorizontalToWest].StateChan <- utils.Red

			semaphores[utils.WestRight].StateChan <- utils.Green
			semaphores[utils.SouthRight].StateChan <- utils.Green

			semaphores[utils.WestLeft].StateChan <- utils.Red
			semaphores[utils.SouthLeft].StateChan <- utils.Red

			// north pedestrian
			semaphores[utils.PedestrianNorthLeft].StateChan <- utils.Red
			semaphores[utils.PedestrianNorthRight].StateChan <- utils.Red

			// south pedestrian
			semaphores[utils.PedestrianSouthLeft].StateChan <- utils.Red
			semaphores[utils.PedestrianSouthRight].StateChan <- utils.Red

			// east pedestrians
			semaphores[utils.PedestrianEastNorth].StateChan <- utils.Red
			semaphores[utils.PedestrianEastSouth].StateChan <- utils.Red

			// west pedestrians
			semaphores[utils.PedestrianWestNorth].StateChan <- utils.Red
			semaphores[utils.PedestrianWestSouth].StateChan <- utils.Red

			time.Sleep(time.Second * 10)
		}
	}()
}

func createSemaphores(drawChan chan *utils.SemaphoreMessage) map[string]*semaphore.Semaphore {
	semaphores := make(map[string]*semaphore.Semaphore)
	semaphores[utils.StraightVerticalToNorth] = semaphore.New(utils.Red, utils.StraightVerticalToNorth, drawChan)
	semaphores[utils.StraightVerticalToSouth] = semaphore.New(utils.Red, utils.StraightVerticalToSouth, drawChan)

	semaphores[utils.StraightHorizontalToWest] = semaphore.New(utils.Red, utils.StraightHorizontalToWest, drawChan)
	semaphores[utils.StraightHorizontalToEast] = semaphore.New(utils.Red, utils.StraightHorizontalToEast, drawChan)

	semaphores[utils.WestRight] = semaphore.New(utils.Red, utils.WestRight, drawChan)
	semaphores[utils.SouthRight] = semaphore.New(utils.Red, utils.SouthRight, drawChan)

	semaphores[utils.WestLeft] = semaphore.New(utils.Red, utils.WestLeft, drawChan)
	semaphores[utils.SouthLeft] = semaphore.New(utils.Red, utils.SouthLeft, drawChan)

	semaphores[utils.PedestrianNorthLeft] = semaphore.New(utils.Red, utils.PedestrianNorthLeft, drawChan)
	semaphores[utils.PedestrianNorthRight] = semaphore.New(utils.Red, utils.PedestrianNorthRight, drawChan)
	semaphores[utils.PedestrianSouthLeft] = semaphore.New(utils.Red, utils.PedestrianSouthLeft, drawChan)
	semaphores[utils.PedestrianSouthRight] = semaphore.New(utils.Red, utils.PedestrianSouthRight, drawChan)

	semaphores[utils.PedestrianEastNorth] = semaphore.New(utils.Red, utils.PedestrianEastNorth, drawChan)
	semaphores[utils.PedestrianEastSouth] = semaphore.New(utils.Red, utils.PedestrianEastSouth, drawChan)
	semaphores[utils.PedestrianWestNorth] = semaphore.New(utils.Red, utils.PedestrianWestNorth, drawChan)
	semaphores[utils.PedestrianWestSouth] = semaphore.New(utils.Red, utils.PedestrianWestSouth, drawChan)

	return semaphores
}

func spinSemaphores(semaphores map[string]*semaphore.Semaphore) {
	go semaphores[utils.StraightVerticalToNorth].Start()
	go semaphores[utils.StraightVerticalToSouth].Start()
	go semaphores[utils.StraightHorizontalToEast].Start()
	go semaphores[utils.StraightHorizontalToWest].Start()

	go semaphores[utils.WestRight].Start()
	go semaphores[utils.SouthRight].Start()
	go semaphores[utils.WestLeft].Start()
	go semaphores[utils.SouthLeft].Start()

	go semaphores[utils.PedestrianNorthRight].Start()
	go semaphores[utils.PedestrianNorthLeft].Start()
	go semaphores[utils.PedestrianSouthRight].Start()
	go semaphores[utils.PedestrianSouthLeft].Start()

	go semaphores[utils.PedestrianEastNorth].Start()
	go semaphores[utils.PedestrianEastSouth].Start()
	go semaphores[utils.PedestrianWestNorth].Start()
	go semaphores[utils.PedestrianWestSouth].Start()
}
