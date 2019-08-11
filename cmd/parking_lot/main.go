package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	strings "strings"

	"github.com/singhmeghna79/parking_lot/parking"
)

func main() {

	arg := os.Args
	if len(arg) > 1 {

		filename := arg[1]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("cannot read file")
			return
		}
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				log.Fatalln(err)
			}
			if err == io.EOF {
				break
			}

			cmd := strings.Fields(line)
			executeCommands(cmd)

		}

	} else {
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("$ ")
			cmdString, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			cmdString = strings.TrimSuffix(cmdString, "\n")
			cmd := strings.Fields(cmdString)
			if len(cmd) == 0 {
				continue
			}
			executeCommands(cmd)

		}
	}
}

// All the commands to be executed
func executeCommands(cmd []string) {
	switch cmd[0] {
	case "create_parking_lot":
		totalSlots, err := strconv.Atoi(cmd[1])
		if err != nil {
			fmt.Println("cannot conver total slot to int")
		}
		cpl := parking.CreateParkingLot(totalSlots)
		fmt.Println(cpl)
	case "park":
		regno := cmd[1]
		colour := cmd[2]
		park := parking.Park(string(regno), string(colour))
		fmt.Println(park)
	case "leave":
		slotNo := cmd[1]
		sni, _ := strconv.Atoi(string(slotNo))
		leave := parking.Leave(sni)
		fmt.Println(leave)
	case "status":
		parking.Status()
	case "registration_numbers_for_cars_with_colour":
		carColor := cmd[1]
		registrationNumberWithColor := parking.RegistrationNumbersForCarsWithColour(string(carColor))
		fmt.Println(registrationNumberWithColor)
	case "slot_numbers_for_cars_with_colour":
		carColor := cmd[1]
		slotNumberOfCarWithColour := parking.SlotNumbersForCarsWithColour(string(carColor))
		fmt.Println(slotNumberOfCarWithColour)
	case "slot_number_for_registration_number":
		carRegisrationNumber := cmd[1]
		slotNumberWithregistrationNumber := parking.SlotNumberForRegistrationNumber(string(carRegisrationNumber))
		fmt.Println(slotNumberWithregistrationNumber)

	case "exit":
		os.Exit(0)
	}
}
