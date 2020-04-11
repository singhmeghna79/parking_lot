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
	for i, _ := range parking {
		for j, val := range parking[i] {
			if val.slotAvailable {
				return val.slotAvailable{i, j}
			}
		}
	}
	return -1
}
