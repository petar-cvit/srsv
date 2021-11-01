package utils

import "lab2/internal/semaphore"

type Payload struct {
	Time       int
	Semaphores map[string]*semaphore.Semaphore
	Crossing   map[string]string
}
