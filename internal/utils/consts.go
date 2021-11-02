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

// crossing or waiting
const StraightHorizontalToWest = "straightHorizontalToWest"
const StraightHorizontalToEast = "straightHorizontalToEast"

const StraightVerticalToNorth = "straightVericalToNorth"
const StraightVerticalToSouth = "straightVericalToSouth"

// pedestrian Semaphores
const PedestrianNorth = "pedestrianNorth"
const PedestrianEast = "pedestrianEast"
const PedestrianWest = "pedestrianWest"
const PedestrianSouth = "pedestrianSouth"

// pedestrian SemaphoresDraw
const PedestrianNorthDraw = "pedestrianNorthDraw"
const PedestrianEastDraw = "pedestrianEastDraw"
const PedestrianWestDraw = "pedestrianWestDraw"
const PedestrianSouthDraw = "pedestrianSouthDraw"

// pedestrians
const PedestrianNorthToSouth = "pedestrianNorthToSouth"
const PedestrianSouthToNorth = "pedestrianSouthTONorth"
const PedestrianEastToWest = "pedestrianEastToWest"
const PedestrianWestToEast = "pedestrianWestToEast"

// waiting status
const NotWaiting = 0
const Waiting = 1

func CarMessage(postition string) string {
	return fmt.Sprintf("car-%v", postition)
}
