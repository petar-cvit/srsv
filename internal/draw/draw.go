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
			d.waiting[payload.Position] = utils.NotWaiting
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
		drawWaitingPedestrians(utils.PedestrianWestToEast, d.waiting) +
		"    | " +
		drawHorizontalCrossing(utils.PedestrianNorthDraw, d.crossing, 19) +
		"|\t\t\t")

	fmt.Println("\t\t      " +
		drawSemaphore(d.semaphores[utils.PedestrianNorth]) +
		" |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   ¦\t|   " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"   ¦\t| " +
		drawSemaphore(d.semaphores[utils.PedestrianNorth]) +
		"\t\t\t")

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   ¦\t|   " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"   ¦\t|\t\t\t")

	fmt.Println("\t\t  " +
		drawSemaphore(d.semaphores[utils.PedestrianWest]) +
		"     |   " + drawSemaphore(d.semaphores[utils.SouthLeft]) + "   " +
		"¦   " + drawSemaphore(d.semaphores[utils.StraightVertical]) + "   " +
		"¦   " + drawSemaphore(d.semaphores[utils.SouthRight]) + "   " +
		"|   " + drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) + "   ¦\t|     " +
		drawSemaphore(d.semaphores[utils.PedestrianEast]) +
		" \t\t\t",
	)

	fmt.Println("------------------------" +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		"------------------------")
	fmt.Println("\t\t\t\t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"\t\t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"\t\t " +
		drawSemaphore(d.semaphores[utils.WestRight]))
	fmt.Println("- - - - - - - - - - - - " +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		"- - - - - - - - - - - - ")

	fmt.Println(drawHorizontalCrossing(utils.StraightHorizontalToWest, d.crossing, 32) + " " +
		drawSemaphore(d.semaphores[utils.StraightHorizontal]) +
		drawHorizontalCrossing(utils.StraightHorizontalToWest, d.crossing, 10) +
		drawWaitingCombined(utils.StraightHorizontalToWest, d.waiting, d.crossing))

	fmt.Println("------------------------" +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		"- - - - - - - - - - - - ")

	fmt.Println("\t\t       " +
		drawSemaphore(d.semaphores[utils.WestLeft]) +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		drawSemaphore(d.semaphores[utils.WestLeft]),
	)

	fmt.Println("- - - - - - - - - - - -" +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		"------------------------")

	fmt.Println(drawWaitingCombined(utils.StraightHorizontalToEast, d.waiting, d.crossing) + " " +
		drawHorizontalCrossing(utils.StraightHorizontalToEast, d.crossing, 10) + " " +
		drawSemaphore(d.semaphores[utils.StraightHorizontal]) + " " +
		drawHorizontalCrossing(utils.StraightHorizontalToEast, d.crossing, 32))

	fmt.Println("- - - - - - - - - - - - " +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		"- - - - - - - - - - - - ")
	fmt.Println("\t\t       " + drawSemaphore(d.semaphores[utils.WestRight]) +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t ")
	fmt.Println("------------------------" +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"    \t " +
		"------------------------")

	fmt.Println("\t\t  " +
		drawSemaphore(d.semaphores[utils.PedestrianWest]) + "     |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   |   " + drawSemaphore(d.semaphores[utils.SouthRight]) + "   " +
		"¦   " + drawSemaphore(d.semaphores[utils.StraightVertical]) + "   " +
		"¦   " + drawSemaphore(d.semaphores[utils.SouthLeft]) + "   " +
		"|     " +
		drawSemaphore(d.semaphores[utils.PedestrianEast]) +
		" \t\t\t",
	)

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"   ¦\t|\t\t\t")

	fmt.Println("\t\t      " +
		drawSemaphore(d.semaphores[utils.PedestrianSouth]) +
		" |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToNorth, d.crossing) +
		"   ¦\t| " +
		drawSemaphore(d.semaphores[utils.PedestrianSouth]) +
		"\t\t\t")

	fmt.Println("\t\t\t| " +
		drawHorizontalCrossing(utils.PedestrianSouth, d.crossing, 19) +
		"|\t\t\t")

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, d.crossing) +
		"   |\t¦   " +
		drawWaitingCombined(utils.StraightVerticalToNorth, d.waiting, d.crossing) +
		"   ¦\t|\t\t\t")
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
	if _, exists := waiting[key]; exists {
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
		utils.StraightHorizontal: 0,
		utils.StraightVertical:   0,
		utils.WestRight:          0,
		utils.SouthRight:         0,
		utils.WestLeft:           0,
		utils.SouthLeft:          0,
		utils.PedestrianNorth:    0,
		utils.PedestrianEast:     0,
		utils.PedestrianWest:     0,
		utils.PedestrianSouth:    0,
	}
}

func createWaiting() map[string]int {
	return map[string]int{
		utils.StraightHorizontalToWest: 0,
		utils.StraightHorizontalToEast: 0,
		utils.StraightVerticalToNorth:  0,
		utils.StraightVerticalToSouth:  0,
	}
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
