package draw

import (
	"fmt"
	"os"
	"os/exec"

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
			d.DrawCrossing(payload.Time, payload.Semaphores)
		}
	}
}

func (d *Drawer) DrawCrossing(time int, semaphores map[string]*semaphore.Semaphore) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Println(fmt.Sprintf("time: %v", time))

	fmt.Println("\t\t\t|\t¦\t¦\t|\t¦\t|\t\t\t")
	fmt.Println("\t\t\t|\t¦\t¦\t|\t¦\t|\t\t\t")
	fmt.Println("\t\t      " +
		drawSemaphore(semaphores[utils.PedestrianNorth]) +
		" |\t¦\t¦\t|\t¦\t| " +
		drawSemaphore(semaphores[utils.PedestrianNorth]) +
		"\t\t\t")
	fmt.Println("\t\t\t|\t¦\t¦\t|\t¦\t|\t\t\t")

	fmt.Println("\t\t  " +
		drawSemaphore(semaphores[utils.PedestrianWest]) +
		"     |   " + drawSemaphore(semaphores[utils.SouthLeft]) + "   " +
		"¦   " + drawSemaphore(semaphores[utils.StraightVertical]) + "   " +
		"¦   " + drawSemaphore(semaphores[utils.SouthRight]) + "   " +
		"|\t¦\t|     " +
		drawSemaphore(semaphores[utils.PedestrianEast]) +
		" \t\t\t",
	)

	fmt.Println("------------------------\t\t\t\t\t------------------------")
	fmt.Println("\t\t\t\t\t\t\t\t" + drawSemaphore(semaphores[utils.WestRight]))
	fmt.Println("- - - - - - - - - - - - \t\t\t\t\t- - - - - - - - - - - - ")
	fmt.Println("\t\t\t\t\t\t\t\t" + drawSemaphore(semaphores[utils.StraightHorizontal]))
	fmt.Println("------------------------\t\t\t\t\t- - - - - - - - - - - - ")
	fmt.Println("\t\t       " +
		drawSemaphore(semaphores[utils.WestLeft]) +
		"\t\t\t\t\t" + drawSemaphore(semaphores[utils.WestLeft]),
	)
	fmt.Println("- - - - - - - - - - - - \t\t\t\t\t------------------------")
	fmt.Println("\t\t       " + drawSemaphore(semaphores[utils.StraightHorizontal]))
	fmt.Println("- - - - - - - - - - - - \t\t\t\t\t- - - - - - - - - - - - ")
	fmt.Println("\t\t       " + drawSemaphore(semaphores[utils.WestRight]))
	fmt.Println("------------------------\t\t\t\t\t------------------------")

	fmt.Println("\t\t  " +
		drawSemaphore(semaphores[utils.PedestrianWest]) + "     |\t¦\t" +
		"|   " + drawSemaphore(semaphores[utils.SouthRight]) + "   " +
		"¦   " + drawSemaphore(semaphores[utils.StraightVertical]) + "   " +
		"¦   " + drawSemaphore(semaphores[utils.SouthLeft]) + "   " +
		"|     " +
		drawSemaphore(semaphores[utils.PedestrianEast]) +
		" \t\t\t",
	)

	fmt.Println("\t\t\t|\t¦\t|\t¦\t¦\t|\t\t\t")
	fmt.Println("\t\t      " +
		drawSemaphore(semaphores[utils.PedestrianSouth]) +
		" |\t¦\t|\t¦\t¦\t| " +
		drawSemaphore(semaphores[utils.PedestrianSouth]) +
		"\t\t\t")
	fmt.Println("\t\t\t|\t¦\t|\t¦\t¦\t|\t\t\t")
	fmt.Println("\t\t\t|\t¦\t|\t¦\t¦\t|\t\t\t")
}

func drawSemaphore(sem *semaphore.Semaphore) string {
	if sem.Current == utils.Red {
		return color.Ize(color.Red, "R")
	}

	return color.Ize(color.Green, "G")
}
