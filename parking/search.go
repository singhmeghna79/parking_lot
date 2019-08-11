package parking

import (
	"strconv"
	"strings"
)

func RegistrationNumbersForCarsWithColour(colour string) string {
	var s string
	for _, val := range parking {
		if strings.EqualFold(val.ticket.car.colour, colour) {
			s = s + val.ticket.car.registrationNumber + ", "

		}
	}

	if len(s) == 0 {
		return "Not found"
	}
	return strings.TrimSuffix(s, ", ")

}

func SlotNumbersForCarsWithColour(colour string) string {
	var s string
	for _, val := range parking {
		if strings.EqualFold(val.ticket.car.colour, colour) {
			s = s + strconv.Itoa(val.ticket.parkingSlotNumber) + ", "

		}
	}
	if len(s) == 0 {
		return "Not found"
	}
	return strings.TrimSuffix(s, ", ")
}

func SlotNumberForRegistrationNumber(registrationNumber string) string {
	var s string
	for _, val := range parking {
		if strings.EqualFold(val.ticket.car.registrationNumber, registrationNumber) {
			s = strconv.Itoa(val.ticket.parkingSlotNumber)
			break
		}
	}
	if len(s) == 0 {
		return "Not found"
	}
	return s

}
