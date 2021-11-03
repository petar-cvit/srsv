package draw

import (
	"fmt"
	"lab2/internal/logger"
	"os"
	"os/exec"
	"strings"

	"github.com/TwiN/go-color"

	"lab2/internal/utils"
)

type Drawer struct {
	SemaphoreChan chan *utils.SemaphoreMessage
	WaitingChan   chan string
	CrossingChan  chan *utils.CrossingMessage

	time       int
	crossing   map[string]string
	semaphores map[string]int
	waiting    map[string]int
	logger     *logger.Logger
}

func New() *Drawer {
	crossing := make(map[string]string)

	return &Drawer{
		SemaphoreChan: make(chan *utils.SemaphoreMessage),
		WaitingChan:   make(chan string),
		CrossingChan:  make(chan *utils.CrossingMessage),
		semaphores:    createDrawableSemaphors(),
		waiting:       createWaiting(),
		crossing:      crossing,
	}
}

func (d *Drawer) Start() {
	for {
		select {
		case payload := <-d.SemaphoreChan:
			d.semaphores[payload.Position] = payload.State
			d.DrawCrossing()
			break
		case payload := <-d.WaitingChan:
			d.waiting[payload] = utils.Waiting
			d.DrawCrossing()
			break
		case payload := <-d.CrossingChan:
			fmt.Println("pedestrian crossing", payload.Position, payload.Crossing)
			d.crossing[payload.Position] = getRune(payload.Crossing, payload.Car)
			d.waiting = invalidateWaitingMap(payload, d.waiting)

			d.DrawCrossing()
		}
	}
}

func (d *Drawer) DrawCrossing() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Println(d.time)
	d.time++

	fmt.Println("\t\t\t|\t¦   " +
		drawWaitingCombined(utils.StraightVerticalToSouth, d.waiting, d.crossing) +
		"   ¦\t|   " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"   ¦\t|\t\t\t")

	fmt.Println("\t\t   " +
		drawWaitingPedestrians(utils.PedestrianWestToEastNorth, d.waiting) +
		"    | " +
		drawHorizontalCrossing(utils.PedestrianNorthDraw, d.crossing, 19) +
		"|   " +
		drawWaitingPedestrians(utils.PedestrianEastToWestNorth, d.waiting) +
		"    \t\t")

	fmt.Println("\t\t      " +
		drawSemaphore(d.semaphores[utils.PedestrianNorthLeft]) +
		" |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   ¦\t|   " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"   ¦\t| " +
		drawSemaphore(d.semaphores[utils.PedestrianNorthRight]) +
		"\t\t\t")

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   ¦\t|   " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"   ¦\t|\t\t\t")

	fmt.Println("\t     " +
		drawWaitingPedestrians(utils.PedestrianNorthToSouthEast, d.waiting) + "    " +
		drawSemaphore(d.semaphores[utils.PedestrianWestNorth]) +
		"     |   " + drawSemaphore(d.semaphores[utils.SouthLeft]) + "   " +
		"¦   " + drawSemaphore(d.semaphores[utils.StraightVerticalToSouth]) + "   " +
		"¦   " + drawSemaphore(d.semaphores[utils.SouthRight]) + "   " +
		"|   " + drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) + "   ¦\t|     " +
		drawSemaphore(d.semaphores[utils.PedestrianEastNorth]) + "     " +
		drawWaitingPedestrians(utils.PedestrianNorthToSouthWest, d.waiting) + "   " + "\t\t",
	)

	fmt.Println("------------------------" +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		"------------------------")
	fmt.Println("\t     " +
		drawSingleVertical(utils.PedestrianEastDraw, d.crossing) + "\t\t\t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"\t\t" +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"\t\t " +
		drawSemaphore(d.semaphores[utils.WestRight]) + "         " +
		drawSingleVertical(utils.PedestrianEastDraw, d.crossing))
	fmt.Println("- - - - - - -" +
		drawSingleVertical(utils.PedestrianEastDraw, d.crossing) +
		" - - - - - " +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		"- - - - - " +
		drawSingleVertical(utils.PedestrianWestDraw, d.crossing) + " - - - - - - ")

	fmt.Println(drawHorizontalCrossing(utils.StraightHorizontalToWest, d.crossing, 32) + " " +
		drawSemaphore(d.semaphores[utils.StraightHorizontalToWest]) +
		drawHorizontalCrossing(utils.StraightHorizontalToWest, d.crossing, 10) +
		drawWaitingCombined(utils.StraightHorizontalToWest, d.waiting, d.crossing))

	fmt.Println("------------------------" +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		"- - - - - " +
		drawSingleVertical(utils.PedestrianWestDraw, d.crossing) + " - - - - - - ")

	fmt.Println("\t     " +
		drawSingleVertical(utils.PedestrianEastDraw, d.crossing) + "         " +
		drawSemaphore(d.semaphores[utils.WestLeft]) +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		drawSemaphore(d.semaphores[utils.WestLeft]) + "         " +
		drawSingleVertical(utils.PedestrianEastDraw, d.crossing),
	)

	fmt.Println("- - - - - - -" +
		drawSingleVertical(utils.PedestrianEastDraw, d.crossing) +
		" - - - - - " +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		"------------------------")

	fmt.Println(drawWaitingCombined(utils.StraightHorizontalToEast, d.waiting, d.crossing) + " " +
		drawHorizontalCrossing(utils.StraightHorizontalToEast, d.crossing, 10) + " " +
		drawSemaphore(d.semaphores[utils.StraightHorizontalToEast]) + " " +
		drawHorizontalCrossing(utils.StraightHorizontalToEast, d.crossing, 32))

	fmt.Println("- - - - - - -" +
		drawSingleVertical(utils.PedestrianEastDraw, d.crossing) +
		" - - - - - " +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		"- - - - - " +
		drawSingleVertical(utils.PedestrianWestDraw, d.crossing) + " - - - - - - ")
	fmt.Println("\t     " +
		drawSingleVertical(utils.PedestrianEastDraw, d.crossing) + "         " +
		drawSemaphore(d.semaphores[utils.WestRight]) +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing))
	fmt.Println("------------------------" +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		"------------------------")

	fmt.Println("\t     " +
		drawWaitingPedestrians(utils.PedestrianSouthToNorthWest, d.waiting) + "    " +
		drawSemaphore(d.semaphores[utils.PedestrianWestSouth]) + "     |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   |   " + drawSemaphore(d.semaphores[utils.SouthRight]) + "   " +
		"¦   " + drawSemaphore(d.semaphores[utils.StraightVerticalToNorth]) + "   " +
		"¦   " + drawSemaphore(d.semaphores[utils.SouthLeft]) + "   " +
		"|     " +
		drawSemaphore(d.semaphores[utils.PedestrianEastSouth]) + "     " +
		drawWaitingPedestrians(utils.PedestrianSouthToNorthEast, d.waiting) + "   " + "\t\t",
	)

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"   ¦\t|\t\t\t")

	fmt.Println("\t\t      " +
		drawSemaphore(d.semaphores[utils.PedestrianSouthLeft]) +
		" |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"   ¦\t| " +
		drawSemaphore(d.semaphores[utils.PedestrianSouthRight]) +
		"\t\t\t")

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"   ¦\t| \t\t\t")

	fmt.Println("\t\t   " +
		drawWaitingPedestrians(utils.PedestrianWestToEastSouth, d.waiting) +
		"    | " +
		drawHorizontalCrossing(utils.PedestrianSouthDraw, d.crossing, 19) +
		"|   " +
		drawWaitingPedestrians(utils.PedestrianEastToWestSouth, d.waiting) +
		"    \t\t")

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   |\t¦   " +
		drawWaitingCombined(utils.StraightVerticalToNorth, d.waiting, d.crossing) +
		"   ¦\t|\t\t\t")

	for k, v := range d.crossing {
		fmt.Println(k, v)
	}

	fmt.Println()
	for k, v := range d.waiting {
		fmt.Println(k, v)
	}
}

func drawSemaphore(sem int) string {
	if sem == utils.Red {
		return color.Ize(color.Red, "R")
	}

	return color.Ize(color.Green, "G")
}

func drawHorizontalCrossing(key string, crossing map[string]string, times int) string {
	if val, exists := crossing[key]; exists {
		return strings.Repeat(val+" ", times)
	}

	return strings.Repeat("  ", times)
}

func drawSingleVertical(key string, crossing map[string]string) string {
	if val, exists := crossing[key]; exists {
		return val
	}

	return " "
}

func drawWaitingPedestrians(key string, waiting map[string]int) string {
	if val, exists := waiting[key]; exists && val == utils.Waiting {
		return color.Ize(color.Blue, "p")
	}

	return " "
}

func drawWaitingCombined(key string, waiting map[string]int, crossing map[string]string) string {
	singleVertical := drawSingleVertical(key, crossing)

	if singleVertical != " " {
		return singleVertical
	}

	if val, exists := waiting[key]; exists && val == utils.Waiting {
		return color.Ize(color.Yellow, "A")
	}

	return " "
}

func createDrawableSemaphors() map[string]int {
	return map[string]int{
		utils.StraightVerticalToNorth:  0,
		utils.StraightVerticalToSouth:  0,
		utils.StraightHorizontalToEast: 0,
		utils.StraightHorizontalToWest: 0,
		utils.WestRight:                0,
		utils.SouthRight:               0,
		utils.WestLeft:                 0,
		utils.SouthLeft:                0,
		utils.PedestrianNorthLeft:      0,
		utils.PedestrianNorthRight:     0,
		utils.PedestrianSouthLeft:      0,
		utils.PedestrianSouthRight:     0,
		utils.PedestrianEastNorth:      0,
		utils.PedestrianEastSouth:      0,
		utils.PedestrianWestNorth:      0,
		utils.PedestrianWestSouth:      0,
	}
}

func createWaiting() map[string]int {
	return map[string]int{
		utils.StraightHorizontalToWest: 0,
		utils.StraightHorizontalToEast: 0,
		utils.StraightVerticalToNorth:  0,
		utils.StraightVerticalToSouth:  0,

		utils.PedestrianEastToWestNorth: 0,
		utils.PedestrianWestToEastNorth: 0,
		utils.PedestrianEastToWestSouth: 0,
		utils.PedestrianWestToEastSouth: 0,
	}
}

func invalidateWaitingMap(payload *utils.CrossingMessage, waiting map[string]int) map[string]int {
	if payload.Position == utils.PedestrianNorthDraw {
		waiting[utils.PedestrianEastToWestNorth] = utils.NotWaiting
		waiting[utils.PedestrianWestToEastNorth] = utils.NotWaiting
	}

	if payload.Position == utils.PedestrianSouthDraw {
		waiting[utils.PedestrianEastToWestSouth] = utils.NotWaiting
		waiting[utils.PedestrianWestToEastSouth] = utils.NotWaiting
	}

	return waiting
}

func getRune(crossing bool, car bool) string {
	if crossing {
		if car {
			return color.Ize(color.Yellow, "A")
		}

		return color.Ize(color.Blue, "p")
	}

	return " "
}
