package parking

import "errors"

func (car car) createTicket() (ticket, error) {
	slot := searchNearestSlot()
	if slot == -1 {
		return ticket{}, errors.New("Sorry, parking lot is full")
	}

	t := ticket{car, slot}

	return t, nil
}

func searchNearestSlot() int {
	for _, val := range parking {
		if val.slotAvailable {
			return val.ticket.parkingSlotNumber
		}
	}
	return -1
}
