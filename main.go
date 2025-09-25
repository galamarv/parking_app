package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Car represents a vehicle with its registration number.
type Car struct {
	RegistrationNumber string
}

// ParkingLot manages the parking slots and car information.
type ParkingLot struct {
	Capacity  int
	Slots     []*Car
	RegToSlot map[string]int
}

var lot *ParkingLot

// createParkingLot initializes a new parking lot of a given size.
func createParkingLot(capacity int) {
	lot = &ParkingLot{
		Capacity:  capacity,
		Slots:     make([]*Car, capacity), // Index 0 is Slot 1
		RegToSlot: make(map[string]int),
	}
}

// park finds the nearest empty slot and allocates it to a car.
func park(regNumber string) {
	if lot == nil {
		fmt.Println("Error: Parking lot not created yet.")
		return
	}

	for i := 0; i < lot.Capacity; i++ {
		if lot.Slots[i] == nil {
			slotNumber := i + 1
			lot.Slots[i] = &Car{RegistrationNumber: regNumber}
			lot.RegToSlot[regNumber] = slotNumber
			fmt.Printf("Allocated slot number: %d\n", slotNumber)
			return
		}
	}
	fmt.Println("Sorry, parking lot is full")
}

// leave removes a car from a slot, calculates the charge, and frees the slot.
func leave(regNumber string, hours int) {
	if lot == nil {
		fmt.Println("Error: Parking lot not created yet.")
		return
	}

	slotNumber, exists := lot.RegToSlot[regNumber]
	if !exists {
		fmt.Printf("Registration number %s not found\n", regNumber)
		return
	}

	// Calculate charge
	charge := 10
	if hours > 2 {
		charge += (hours - 2) * 10
	}

	// Free the slot
	lot.Slots[slotNumber-1] = nil
	delete(lot.RegToSlot, regNumber)

	fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", regNumber, slotNumber, charge)
}

// status prints the current occupancy of the parking lot.
func status() {
	if lot == nil {
		fmt.Println("Error: Parking lot not created yet.")
		return
	}

	fmt.Println("Slot No. Registration No.")
	for i, car := range lot.Slots {
		if car != nil {
			fmt.Printf("%d %s\n", i+1, car.RegistrationNumber)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_filename>")
		return
	}
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		command := strings.ToLower(parts[0])
		args := parts[1:]

		switch command {
		case "create_parking_lot":
			capacity, _ := strconv.Atoi(args[0])
			createParkingLot(capacity)
		case "park":
			park(args[0])
		case "leave":
			regNumber := args[0]
			hours, _ := strconv.Atoi(args[1])
			leave(regNumber, hours)
		case "status":
			status()
		default:
			fmt.Println("Unknown command:", command)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}
}
