package main

import (
	//"encoding/json"
	"fmt"
	//"os"
	"strings"
)

func main() {

	day := "Tuesday"
	meal := "Lunch"
	item := "Cornflakes"
	menuItems := GetMenuItems(day, meal)
	count := GetMenuItemscount(day, meal)

	if menuItems != nil {
		fmt.Println("Menu items for", strings.ToUpper(day), strings.ToUpper(meal), ": \n", menuItems)
		fmt.Println("Number of menu items for", strings.ToUpper(day), strings.ToUpper(meal), ":", count)
		fmt.Printf("Is %s present in the menu for the specified day and meal: %t", item, CheckMenuItems(day, meal, item))
	}

	err := ConvertMenuToJSON("Menu.xlsx")
	if err != nil {
		fmt.Println("Error converting menu to JSON:", err)
	}

	details(day, meal)

}
