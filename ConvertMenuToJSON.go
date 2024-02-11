package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func ConvertMenuToJSON(fileName string) error {
	f, err := excelize.OpenFile(fileName)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("error opening the file: %w", err)
	}

	menu, err := f.GetCols("Sheet1")
	if err != nil {
		return fmt.Errorf("error extracting data: %w", err)
	}

	jsonData, err := json.MarshalIndent(menu, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling data to JSON: %w", err)
	}

	jsonFileName := strings.TrimSuffix(fileName, ".xlsx") + ".json"

	file, err := os.OpenFile(jsonFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file for writing: %w", err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing data to file: %w", err)
	}

	fmt.Println("Menu successfully converted to JSON and saved as:", jsonFileName)
	return nil
}
