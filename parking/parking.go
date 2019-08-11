package parking

import (
	"fmt"
	"strconv"
	"strings"
)

func CreateParkingLot(totalslots int) string {
	if totalslots <= 0 {
		return "invalid parking slots"
	}
	for i := 1; i <= totalslots; i++ {
		var p parkingSlot
		p.slotAvailable = true
		p.ticket.parkingSlotNumber = i
		parking = append(parking, p)

	}
	return fmt.Sprintf("%s%d%s", "Created a parking lot with ", totalslots, " slots")
}

func Park(regno string, colour string) string {
	var c car
	c.registrationNumber = regno
	c.colour = strings.Title(colour)
	t, err := c.createTicket()
	if err != nil {
		return err.Error()
	}

	parking[t.parkingSlotNumber-1] = parkingSlot{t, false}

	return fmt.Sprintf("%s%d", "Allocated slot number: ", t.parkingSlotNumber)

}

func Leave(slotno int) string {
	parking[slotno-1].ticket.car = car{}
	parking[slotno-1].slotAvailable = true

	sn := strconv.Itoa(slotno)
	return "Slot number " + sn + " is free"
}
