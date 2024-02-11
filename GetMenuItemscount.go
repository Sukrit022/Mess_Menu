package main

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

func GetMenuItemscount(day string, meal string) int {
	f, err := excelize.OpenFile("Menu.xlsx")
	defer f.Close()
	day = strings.ToUpper(day)
	meal = strings.ToUpper(meal)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	row, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	col, err := f.GetCols("Sheet1")
	if err != nil {
		fmt.Println(err)
		return 0
	}

	dayColIndex := -1
	for colIndex, colName := range row[0] {
		if colName == day {
			dayColIndex = colIndex
			break
		}
	}
	if dayColIndex == -1 {
		fmt.Println("Day not found:", day)
		return 0
	}

	mealRowIndex := -1
	for rowIndex, rowName := range col[dayColIndex] {
		if strings.EqualFold(rowName, meal) {
			mealRowIndex = rowIndex
			break
		}
	}
	if mealRowIndex == -1 {
		fmt.Println("Meal not found:", meal)
		return 0
	}

	var c int = 0
	for colIndex := mealRowIndex + 1; colIndex < len(row); colIndex++ {
		menuItem := strings.TrimSpace(col[dayColIndex][colIndex])
		if menuItem == day {
			break
		}
		if menuItem != "" {
			c++
		}
	}
	return c
}
