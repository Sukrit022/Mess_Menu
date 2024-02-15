package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var day, meal, item string
	var count int
	var menuItems []string

	fmt.Println("Enter the day:")
	n, err := fmt.Scanln(&day)
	if err != nil || n != 1 {
		fmt.Println("Error:", err)
	}

	fmt.Println("Enter the meal:")
	n, err = fmt.Scanln(&meal)
	if err != nil || n != 1 {
		fmt.Println("Error:", err)
	}

	
	fmt.Println("Enter the item(full name):")
	item, err = reader.ReadString('\n')
	if err != nil || n != 1 {
		fmt.Println("Error:", err)
	}
	item = strings.TrimSpace(item)

	menuItems = GetMenuItems(day, meal)
	count = GetMenuItemscount(day, meal)

	if menuItems != nil {
		fmt.Println("Menu items for", strings.ToUpper(day), strings.ToUpper(meal), ": \n", menuItems)
		fmt.Println("Number of menu items for", strings.ToUpper(day), strings.ToUpper(meal), ":", count)
		fmt.Printf("Is %s present in the menu for the specified day and meal: %t \n", item, CheckMenuItems(day, meal, item))
	}

	err1 := ConvertMenuToJSON("Menu.xlsx")
	if err1 != nil {
		fmt.Println("Error converting menu to JSON: \n", err1)
	}

	fmt.Printf("Here are the final details for the menu of %s %s : \n", strings.ToUpper(day), strings.ToUpper(meal))
	details(day, meal)

}
