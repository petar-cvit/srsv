package controller

import (
	"github.com/TwiN/go-color"
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
	semaphores := createSemaphores()
	spinSemaphores(semaphores)

	drawer := draw.New()
	go drawer.Start()

	go func() {
		t := 0

		for {
			time.Sleep(time.Second * 3)
			t++

			semaphores[utils.StraightHorizontal].StateChan <- utils.Green
			semaphores[utils.StraightVertical].StateChan <- utils.Red

			semaphores[utils.WestRight].StateChan <- utils.Red
			semaphores[utils.SouthRight].StateChan <- utils.Red

			semaphores[utils.WestLeft].StateChan <- utils.Red
			semaphores[utils.SouthLeft].StateChan <- utils.Red

			semaphores[utils.PedestrianEast].StateChan <- utils.Red
			semaphores[utils.PedestrianWest].StateChan <- utils.Red
			semaphores[utils.PedestrianNorth].StateChan <- utils.Green
			semaphores[utils.PedestrianSouth].StateChan <- utils.Green

			drawer.Input <- &utils.Payload{
				Time:       t,
				Semaphores: semaphores,
				Crossing: map[string]string{
					utils.StraightHorizontalToWest: color.Ize(color.Yellow, "A"),
					utils.StraightHorizontalToEast: color.Ize(color.Yellow, "A"),
				},
			}

			time.Sleep(time.Second * 3)
			t++

			semaphores[utils.StraightHorizontal].StateChan <- utils.Red
			semaphores[utils.StraightVertical].StateChan <- utils.Red

			semaphores[utils.WestRight].StateChan <- utils.Red
			semaphores[utils.SouthRight].StateChan <- utils.Red

			semaphores[utils.WestLeft].StateChan <- utils.Green
			semaphores[utils.SouthLeft].StateChan <- utils.Green

			semaphores[utils.PedestrianEast].StateChan <- utils.Red
			semaphores[utils.PedestrianWest].StateChan <- utils.Red
			semaphores[utils.PedestrianNorth].StateChan <- utils.Red
			semaphores[utils.PedestrianSouth].StateChan <- utils.Red

			drawer.Input <- &utils.Payload{
				Time:       t,
				Semaphores: semaphores,
			}

			time.Sleep(time.Second * 3)
			t++

			semaphores[utils.StraightHorizontal].StateChan <- utils.Red
			semaphores[utils.StraightVertical].StateChan <- utils.Green

			semaphores[utils.WestRight].StateChan <- utils.Red
			semaphores[utils.SouthRight].StateChan <- utils.Red

			semaphores[utils.WestLeft].StateChan <- utils.Red
			semaphores[utils.SouthLeft].StateChan <- utils.Red

			semaphores[utils.PedestrianEast].StateChan <- utils.Green
			semaphores[utils.PedestrianWest].StateChan <- utils.Green
			semaphores[utils.PedestrianNorth].StateChan <- utils.Red
			semaphores[utils.PedestrianSouth].StateChan <- utils.Red

			drawer.Input <- &utils.Payload{
				Time:       t,
				Semaphores: semaphores,
				Crossing: map[string]string{
					utils.StraightVerticalToSouth: color.Ize(color.Yellow, "A"),
					utils.StraightVerticalToNorth: color.Ize(color.Yellow, "A"),
				},
			}

			time.Sleep(time.Second * 3)
			t++

			semaphores[utils.StraightHorizontal].StateChan <- utils.Red
			semaphores[utils.StraightVertical].StateChan <- utils.Red

			semaphores[utils.WestRight].StateChan <- utils.Green
			semaphores[utils.SouthRight].StateChan <- utils.Green

			semaphores[utils.WestLeft].StateChan <- utils.Red
			semaphores[utils.SouthLeft].StateChan <- utils.Red

			semaphores[utils.PedestrianEast].StateChan <- utils.Red
			semaphores[utils.PedestrianWest].StateChan <- utils.Red
			semaphores[utils.PedestrianNorth].StateChan <- utils.Red
			semaphores[utils.PedestrianSouth].StateChan <- utils.Red

			drawer.Input <- &utils.Payload{
				Time:       t,
				Semaphores: semaphores,
			}

		}
	}()
}

func createSemaphores() map[string]*semaphore.Semaphore {
	semaphores := make(map[string]*semaphore.Semaphore)
	semaphores[utils.StraightHorizontal] = semaphore.New(utils.Green)
	semaphores[utils.StraightVertical] = semaphore.New(utils.Green)

	semaphores[utils.WestRight] = semaphore.New(utils.Red)
	semaphores[utils.SouthRight] = semaphore.New(utils.Red)

	semaphores[utils.WestLeft] = semaphore.New(utils.Red)
	semaphores[utils.SouthLeft] = semaphore.New(utils.Red)

	semaphores[utils.PedestrianNorth] = semaphore.New(utils.Red)
	semaphores[utils.PedestrianEast] = semaphore.New(utils.Red)
	semaphores[utils.PedestrianWest] = semaphore.New(utils.Red)
	semaphores[utils.PedestrianSouth] = semaphore.New(utils.Red)

	return semaphores
}

func spinSemaphores(semaphores map[string]*semaphore.Semaphore) {
	go semaphores[utils.StraightHorizontal].Start()
	go semaphores[utils.StraightVertical].Start()
	go semaphores[utils.WestRight].Start()
	go semaphores[utils.SouthRight].Start()
	go semaphores[utils.WestLeft].Start()
	go semaphores[utils.SouthLeft].Start()

	go semaphores[utils.PedestrianNorth].Start()
	go semaphores[utils.PedestrianEast].Start()
	go semaphores[utils.PedestrianWest].Start()
	go semaphores[utils.PedestrianSouth].Start()
}
