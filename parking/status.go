package parking

import "fmt"

func Status() {
	fmt.Println("Slot No.    Registration No    Colour")
	for _, val := range parking {
		if val.slotAvailable {
			continue
		} else {
			fmt.Printf("%-8d    %-15s    %-6s\n", val.ticket.parkingSlotNumber, val.ticket.car.registrationNumber, val.ticket.car.colour)
		}
	}
}
