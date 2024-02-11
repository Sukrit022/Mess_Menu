package main

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

func CheckMenuItems(day string, meal string, item string) bool {
	f, err := excelize.OpenFile("Menu.xlsx")
	defer f.Close()
	day = strings.ToUpper(day)
	meal = strings.ToUpper(meal)
	item = strings.ToUpper(item)

	if err != nil {
		fmt.Println(err)
		return false
	}

	row, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return false
	}

	col, err := f.GetCols("Sheet1")
	if err != nil {
		fmt.Println(err)
		return false
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
		return false
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
		return false
	}

	var menuItems []string
	var c int = 0
	for colIndex := mealRowIndex + 1; colIndex < len(row); colIndex++ {
		menuItem := strings.TrimSpace(col[dayColIndex][colIndex])
		if menuItem == day {
			break
		}
		if menuItem != "" {
			menuItems = append(menuItems, menuItem)
			c++
		}
	}

	for i := 0; i < c; i++ {
		if menuItems[i] == item {
			return true
		}
	}
	return false
}
