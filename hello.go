package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// NoOfStudents just checks if rollNumber is valid or not
// Can be 1 or 0
func NoOfStudents(rollNumber string) int {
	if rollNumber != "-" {
		return 1
	}
	return 0
}

// AddCount adds the student in department count
func AddCount(rollNumber string, department string, m map[string]int) {
	m[department] = m[department] + 1
}

func displayData(currentFloor int, m map[string]int) {
	totalCount := 0
	for department, numberOfStudents := range m {
		if department == "- -" {
			continue
		}
		fmt.Printf("[%s]: [%d] ", department, numberOfStudents)
		totalCount += numberOfStudents
		delete(m, department)
	}
	fmt.Printf("\nTotal students in Floor %d: %d ", currentFloor, totalCount)
	fmt.Println()
}

func main() {
	FILE := os.Args[1]
	STARTINGFLOOR := 1

	fmt.Printf("-------- Analysing %s starting with floor %d -------- \n", FILE, STARTINGFLOOR)
	csvFile, _ := os.Open(FILE)
	data := csv.NewReader(bufio.NewReader(csvFile))

	currentFloor := STARTINGFLOOR
	m := make(map[string]int)
	for {
		record, err := data.Read()
		if err == io.EOF {
			break
		}

		roomNo := record[2]
		rollNumber := record[3]
		department := record[5] + " " + record[6]

		val, _ := strconv.Atoi(roomNo)
		if val/100 == currentFloor {
			AddCount(rollNumber, department, m)
		} else {
			// Floor has been changed, display the data and reset
			displayData(currentFloor, m)

			// Actulally change the floor
			currentFloor = val / 100
			AddCount(rollNumber, department, m)
		}
	}
	// for the last loop
	displayData(currentFloor, m)
	fmt.Printf("-------- DONE! -------- \n")
}
