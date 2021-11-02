package utils

type SemaphoreMessage struct {
	Position string
	State    int
}

type CrossingMessage struct {
	Position string
	Crossing bool
}
