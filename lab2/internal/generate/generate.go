package generate

import (
	"math/rand"

	"srsv/internal/passenger"
	"srsv/internal/utils"
)

func GeneratePassangers(numPassengers, numFloors int, elevatorChan chan int) []*passenger.Passenger {
	passengers := make([]*passenger.Passenger, 0, numPassengers)

	letters := lower + upper

	inRune := []rune(letters)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	for i := 0; i < numPassengers; i++ {
		toFLoor := rand.Intn(numFloors)
		fromFloor := rand.Intn(numFloors)
		for fromFloor == toFLoor {
			fromFloor = rand.Intn(numFloors)
		}

		passengers = append(passengers,
			passenger.New(
				string(inRune[i]),
				toFLoor+1,
				fromFloor+1,
				utils.GetDirection(toFLoor, fromFloor),
				elevatorChan,
			),
		)
	}

	return passengers
}

const lower = "abcdefghijklmnoprstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPRSTUVWXYZ"
