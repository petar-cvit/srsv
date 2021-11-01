package utils

import "fmt"

// Semaphore states
const Red = 0
const Green = 1

// vehicle Semaphores
const StraightHorizontal = "straightHorizontal"
const StraightVertical = "straightVertical"

const WestRight = "westRight"
const SouthRight = "southRight"

const WestLeft = "westLeft"
const SouthLeft = "southLeft"

// pedestrian Semaphores
const PedestrianNorth = "pedestrianNorth"
const PedestrianEast = "pedestrianEast"
const PedestrianWest = "pedestrianWest"
const PedestrianSouth = "pedestrianSouth"

// crossing or waiting
const StraightHorizontalToWest = "straightHorizontalToWest"
const StraightHorizontalToEast = "straightHorizontalToEast"

const StraightVerticalToNorth = "straightVericalToNorth"
const StraightVerticalToSouth = "straightVericalToSouth"

// waiting status
const NotWaiting = 0
const Waiting = 1

func CarMessage(postition string) string {
	return fmt.Sprintf("car-%v", postition)
}
