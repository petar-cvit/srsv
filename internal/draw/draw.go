package draw

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/TwiN/go-color"

	"lab2/internal/utils"
)

type Drawer struct {
	SemaphoreChan chan *utils.SemaphoreMessage

	semaphores map[string]int
	waiting    map[string]int
}

func New() *Drawer {
	return &Drawer{
		SemaphoreChan: make(chan *utils.SemaphoreMessage),
		semaphores:    createDrawableSemaphors(),
		waiting:       createWaiting(),
	}
}

func (d *Drawer) Start() {
	for {
		select {
		case payload := <-d.SemaphoreChan:
			d.semaphores[payload.Position] = payload.State
			d.DrawCrossing(0, d.semaphores, map[string]string{})
		}
	}
}

func (d *Drawer) DrawCrossing(time int, semaphores map[string]int, crossing map[string]string) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Println(fmt.Sprintf("time: %v", time))

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   ¦\t|   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   ¦\t|\t\t\t")

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   ¦\t|   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   ¦\t|\t\t\t")

	fmt.Println("\t\t      " +
		drawSemaphore(semaphores[utils.PedestrianNorth]) +
		" |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   ¦\t|   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   ¦\t| " +
		drawSemaphore(semaphores[utils.PedestrianNorth]) +
		"\t\t\t")

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   ¦\t|   " +
		drawSingleVertical(utils.StraightVerticalToNorth, crossing) +
		"   ¦\t|\t\t\t")

	fmt.Println("\t\t  " +
		drawSemaphore(semaphores[utils.PedestrianWest]) +
		"     |   " + drawSemaphore(semaphores[utils.SouthLeft]) + "   " +
		"¦   " + drawSemaphore(semaphores[utils.StraightVertical]) + "   " +
		"¦   " + drawSemaphore(semaphores[utils.SouthRight]) + "   " +
		"|   " + drawSingleVertical(utils.StraightVerticalToSouth, crossing) + "   ¦\t|     " +
		drawSemaphore(semaphores[utils.PedestrianEast]) +
		" \t\t\t",
	)

	fmt.Println("------------------------" +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, crossing) +
		"    \t " +
		"------------------------")
	fmt.Println("\t\t\t\t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"\t\t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"\t\t " +
		drawSemaphore(semaphores[utils.WestRight]))
	fmt.Println("- - - - - - - - - - - - " +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, crossing) +
		"    \t " +
		"- - - - - - - - - - - - ")

	fmt.Println(drawHorizontalCrossing(utils.StraightHorizontalToEast, crossing, 32) + " " +
		drawSemaphore(semaphores[utils.StraightHorizontal]) +
		drawHorizontalCrossing(utils.StraightHorizontalToEast, crossing, 11))

	fmt.Println("------------------------" +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, crossing) +
		"    \t " +
		"- - - - - - - - - - - - ")

	fmt.Println("\t\t       " +
		drawSemaphore(semaphores[utils.WestLeft]) +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, crossing) +
		"    \t " +
		drawSemaphore(semaphores[utils.WestLeft]),
	)

	fmt.Println("- - - - - - - - - - - -" +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, crossing) +
		"    \t " +
		"------------------------")

	fmt.Println(drawHorizontalCrossing(utils.StraightHorizontalToWest, crossing, 11) + " " +
		drawSemaphore(semaphores[utils.StraightHorizontal]) + " " +
		drawHorizontalCrossing(utils.StraightHorizontalToWest, crossing, 32))

	fmt.Println("- - - - - - - - - - - - " +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, crossing) +
		"    \t " +
		"- - - - - - - - - - - - ")
	fmt.Println("\t\t       " + drawSemaphore(semaphores[utils.WestRight]) +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, crossing) +
		"    \t ")
	fmt.Println("------------------------" +
		"  \t    " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"    \t    " +
		drawSingleVertical(utils.StraightVerticalToNorth, crossing) +
		"    \t " +
		"------------------------")

	fmt.Println("\t\t  " +
		drawSemaphore(semaphores[utils.PedestrianWest]) + "     |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   |   " + drawSemaphore(semaphores[utils.SouthRight]) + "   " +
		"¦   " + drawSemaphore(semaphores[utils.StraightVertical]) + "   " +
		"¦   " + drawSemaphore(semaphores[utils.SouthLeft]) + "   " +
		"|     " +
		drawSemaphore(semaphores[utils.PedestrianEast]) +
		" \t\t\t",
	)

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   ¦\t|\t\t\t")

	fmt.Println("\t\t      " +
		drawSemaphore(semaphores[utils.PedestrianSouth]) +
		" |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   ¦\t|   " +
		drawSingleVertical(utils.StraightVerticalToNorth, crossing) +
		"   ¦\t| " +
		drawSemaphore(semaphores[utils.PedestrianSouth]) +
		"\t\t\t")

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   ¦\t|\t\t\t")

	fmt.Println("\t\t\t|\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
		"   |\t¦   " +
		drawSingleVertical(utils.StraightVerticalToSouth, crossing) +
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
