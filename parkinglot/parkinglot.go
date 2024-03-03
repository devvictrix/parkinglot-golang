package parkinglot

import (
	"strconv"
	"strings"
	"fmt"
)

type ParkingLot struct {
	Slots    []*Car
	Capacity int
}

var parkingLot *ParkingLot

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		Slots:    make([]*Car, capacity),
		Capacity: capacity,
	}
}

func (pl *ParkingLot) Park(car *Car) int {
	for i := 0; i < pl.Capacity; i++ {
		if pl.Slots[i] == nil {
			pl.Slots[i] = car   
			return i + 1
		}
	}
	return -1 
}

func (pl *ParkingLot) Leave(slotNumber int) bool {
	if slotNumber-1 < pl.Capacity && pl.Slots[slotNumber-1] != nil {
		pl.Slots[slotNumber-1] = nil
		return true
	}
	return false
}

func (pl *ParkingLot) Status() {
	fmt.Println("Slot No. Registration No Color")
	for i, car := range pl.Slots {
		if car != nil {
			fmt.Printf("%d %s %s\n", i+1, car.RegistrationNumber, car.Color)
		}
	}
}

func (pl *ParkingLot) GetRegistrationNumbersForColor(color string) []string {
    var registrationNumbers []string
    for _, car := range pl.Slots {
        if car != nil && strings.EqualFold(car.Color, color) {
            registrationNumbers = append(registrationNumbers, car.RegistrationNumber)
        }
    }
    return registrationNumbers
}

func (pl *ParkingLot) GetSlotNumbersForColor(color string) []int {
    var slotNumbers []int
    for i, car := range pl.Slots {
        if car != nil && strings.EqualFold(car.Color, color) {
            slotNumbers = append(slotNumbers, i+1)
        }
    }
    return slotNumbers
}

func (pl *ParkingLot) GetSlotNumberForRegistrationNumber(registrationNumber string) int {
    for i, car := range pl.Slots {
        if car != nil && strings.EqualFold(car.RegistrationNumber, registrationNumber) {
            return i + 1
        }
    }
    return -1
}

func ProcessCommand(commandLine string) {
	commands := strings.Fields(commandLine)
	if len(commands) == 0 {
		return
	}

	switch commands[0] {
	case "create_parking_lot":
		size, err := strconv.Atoi(commands[1])
		if err != nil {
			fmt.Println("Invalid input for parking lot size.")
			return
		}
		parkingLot = NewParkingLot(size)
		fmt.Printf("Created a parking lot with %d slots\n", size)
	case "park":
		if parkingLot == nil {
			fmt.Println("Parking lot not created.")
			return
		}
		regNumber := commands[1]
		color := commands[2]
		slotNumber := parkingLot.Park(&Car{RegistrationNumber: regNumber, Color: color})
		if slotNumber != -1 {
			fmt.Printf("Allocated slot number: %d\n", slotNumber)
		} else {
			fmt.Println("Sorry, parking lot is full")
		}
	case "leave":
		slotNumber, err := strconv.Atoi(commands[1])
		if err != nil {
			fmt.Println("Invalid input for slot number.")
			return
		}
		if parkingLot.Leave(slotNumber) {
			fmt.Printf("Slot number %d is free\n", slotNumber)
		} else {
			fmt.Println("Leave operation failed")
		}
	case "status":
		if parkingLot == nil {
			fmt.Println("Parking lot not created.")
			return
		}
		parkingLot.Status()
    case "registration_numbers_for_cars_with_color":
        if len(commands) < 2 {
            fmt.Println("Missing color parameter.")
            return
        }
        color := commands[1]
        registrationNumbers := parkingLot.GetRegistrationNumbersForColor(color)
        fmt.Println(strings.Join(registrationNumbers, ", "))
    case "slot_numbers_for_cars_with_color":
        if len(commands) < 2 {
            fmt.Println("Missing color parameter.")
            return
        }
        color := commands[1]
        slotNumbers := parkingLot.GetSlotNumbersForColor(color)
        for _, slot := range slotNumbers {
            fmt.Print(slot, " ")
        }
        fmt.Println()
    case "slot_number_for_registration_number":
        if len(commands) < 2 {
            fmt.Println("Missing registration number parameter.")
            return
        }
        registrationNumber := commands[1]
        slotNumber := parkingLot.GetSlotNumberForRegistrationNumber(registrationNumber)
        if slotNumber != -1 {
            fmt.Println(slotNumber)
        } else {
            fmt.Println("Not found")
        }
	default:
		fmt.Println("Unknown command:", commands[0])
	}
}
