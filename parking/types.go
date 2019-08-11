package parking

var parking []parkingSlot

type car struct {
	registrationNumber string
	colour             string
}

type ticket struct {
	car               car
	parkingSlotNumber int
}

type parkingSlot struct {
	ticket        ticket
	slotAvailable bool
}
