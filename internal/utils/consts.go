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
const PedestrianNorthLeft = "pedestrianNorthLeft"
const PedestrianNorthRight = "pedestrianNorthRight"
const PedestrianSouthLeft = "pedestrianSouthLeft"
const PedestrianSouthRight = "pedestrianSouthRight"

const PedestrianWest = "pedestrianWest"
const PedestrianEast = "pedestrianEast"

// pedestrian SemaphoresDraw
const PedestrianEastDraw = "pedestrianEastDraw"
const PedestrianWestDraw = "pedestrianWestDraw"

const PedestrianNorthDraw = "pedestrianNorthDraw"
const PedestrianSouthDraw = "pedestrianSouthDraw"

// pedestrians
const PedestrianNorthToSouth = "pedestrianNorthToSouth"
const PedestrianSouthToNorth = "pedestrianSouthToNorth"

// east-west pedestrians
const PedestrianEastToWestNorth = "pedestrianEastToWestNorth"
const PedestrianWestToEastNorth = "pedestrianWestToEastNorth"
const PedestrianEastToWestSouth = "pedestrianEastToWestSouth"
const PedestrianWestToEastSouth = "pedestrianWestToEastSouth"

// waiting status
const NotWaiting = 0
const Waiting = 1

func CarMessage(postition string) string {
	return fmt.Sprintf("car-%v", postition)
}
