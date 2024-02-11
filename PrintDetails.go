package main

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

type menu struct {
	day   string
	date  string
	meal  string
	items []string
}

func (m menu) printDetails() {
	fmt.Println("Day:", m.day)
	fmt.Println("Date:", m.date)
	fmt.Println("Meal:", m.meal)
	fmt.Println("Items:", m.items)
}

func details(day string, meal string) {
	f, err := excelize.OpenFile("Menu.xlsx")
	defer f.Close()
	day = strings.ToUpper(day)
	meal = strings.ToUpper(meal)
	if err != nil {
		fmt.Println(err)
		return
	}

	row, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	var date string

	switch day {
	case "MONDAY":
		date = row[1][0]
	case "TUESDAY":
		date = row[1][1]
	case "WEDNESDAY":
		date = row[1][2]
	case "THURDAY":
		date = row[1][3]
	case "FRIDAY":
		date = row[1][4]
	case "SATURDAY":
		date = row[1][5]
	case "SUNDAY":
		date = row[1][6]
	}

	details := menu{
		day:   day,
		date:  date,
		meal:  meal,
		items: GetMenuItems(day, meal),
	}
	details.printDetails()
}
