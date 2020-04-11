package parking

var parking [][]parkingSlot

var dispatchRule string

type car struct {
	registrationNumber string
	colour             string
}

type ticket struct {
	car               car
	parkingSlotNumber [2]int
}

type parkingSlot struct {
	ticket        ticket
	slotAvailable bool
}
