package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	filePath := "decode_message_2.xlsx"

	err := decodeMessageFromFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func decodeMessageFromFile(filePath string) error {
	// 1. Open the Excel file from the local file system
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer f.Close()

	// 2. Get all rows from the active sheet
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return fmt.Errorf("could not get rows from sheet: %w", err)
	}

	points := make(map[string]string)
	maxX, maxY := 0, 0

	// Iterate over rows, skipping the header
	for i, row := range rows {
		if i == 0 {
			continue // Skip the header row
		}

		if len(row) < 3 {
			continue // Skip malformed rows
		}

		x, err := strconv.Atoi(row[0])
		if err != nil {
			continue
		}

		y, err := strconv.Atoi(row[2])
		if err != nil {
			continue
		}

		char := row[1]
		points[fmt.Sprintf("%d,%d", x, y)] = char

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	// 3. Print the grid
	for y := 0; y <= maxY; y++ {
		line := ""
		for x := 0; x <= maxX; x++ {
			key := fmt.Sprintf("%d,%d", x, y)
			if val, ok := points[key]; ok {
				line += val
			} else {
				line += " "
			}
		}
		fmt.Println(line)
	}

	return nil
}
