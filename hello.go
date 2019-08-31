package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// NoOfStudents finds find the number of students in the present line.
// Can be 1 or 0
func NoOfStudents(record []string) int {

	if record[3] != "-" {
		return 1
	}
	return 0
}

func main() {
	FILE := "16a.csv"
	STARTINGFLOOR := 1
	csvFile, _ := os.Open(FILE)
	data := csv.NewReader(bufio.NewReader(csvFile))

	count := 0
	currentFloor := STARTINGFLOOR
	for {
		record, err := data.Read()
		if err == io.EOF {
			break
		}

		roomNo := record[2]
		val, _ := strconv.Atoi(roomNo)
		if val/100 == currentFloor {
			count += NoOfStudents(record)
		} else {
			fmt.Printf("Floor %d: %d\n", currentFloor, count)
			// change floor
			currentFloor = val / 100
			count = NoOfStudents(record)
		}
	}
	// for the last loop
	fmt.Printf("Floor %d: %d\n", currentFloor, count)

}
