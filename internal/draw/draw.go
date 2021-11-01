package draw

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/TwiN/go-color"

	"lab2/internal/semaphore"
	"lab2/internal/utils"
)

type Drawer struct {
	Input chan *utils.Payload
}

func New() *Drawer {
	return &Drawer{
		Input: make(chan *utils.Payload),
	}
}

func (d *Drawer) Start() {
	for {
		select {
		case payload := <-d.Input:
			d.DrawCrossing(payload.Time, payload.Semaphores, payload.Crossing)
		}
	}
}

func (d *Drawer) DrawCrossing(time int, semaphores map[string]*semaphore.Semaphore, crossing map[string]string) {
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

func drawSemaphore(sem *semaphore.Semaphore) string {
	if sem.Current == utils.Red {
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
