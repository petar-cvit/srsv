package utils

import (
	"srsv/internal/passenger"
)

func MapToSlicePassengers(m map[string]*passenger.Passenger) []*passenger.Passenger {
	s := make([]*passenger.Passenger, 0, len(m))

	for _, p := range m {
		s = append(s, p)
	}

	return s
}

func GetDirection(to, from int) string {
	if to > from {
		return Up
	}

	return Down
}
