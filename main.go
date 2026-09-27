package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"restaurant-menu/restaurant"
)
func main() {

	filename := "menu.json"

	err := restaurant.LoadMenu(filename)

	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	fmt.Println("Previous menu data loaded successfully.")

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("\n===== RESTAURANT MENU SYSTEM =====")
		fmt.Println("1. Add Menu Item")
		fmt.Println("2. Find Menu Item")
		fmt.Println("3. Update Price")
		fmt.Println("4. Remove Menu Item")
		fmt.Println("5. Display All Items")
		fmt.Println("6. Exit")

		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Invalid choice.")
			continue
		}

		switch choice {

		case 1:

			fmt.Print("Enter Item ID: ")
			input, _ = reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(input))

			if err != nil {
				fmt.Println("Invalid Item ID.")
				continue
			}

			fmt.Print("Enter Item Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter Category: ")
			category, _ := reader.ReadString('\n')
			category = strings.TrimSpace(category)

			fmt.Print("Enter Price: ")
			input, _ = reader.ReadString('\n')
			price, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

			if err != nil {
				fmt.Println("Invalid price.")
				continue
			}

			fmt.Print("Is Vegetarian? (true/false): ")
			input, _ = reader.ReadString('\n')
			vegetarian, err := strconv.ParseBool(strings.TrimSpace(input))

			if err != nil {
				fmt.Println("Enter true or false.")
				continue
			}

			item := restaurant.MenuItem{
				ItemID:       id,
				Name:         name,
				Category:     category,
				Price:        price,
				IsVegetarian: vegetarian,
			}

			err = restaurant.AddMenuItem(item)

			if err != nil {
				fmt.Println("Error:", err)
			} else {

				err = restaurant.SaveMenu(filename)

				if err != nil {
					fmt.Println("Error saving data:", err)
				} else {
					fmt.Println("Menu item added and saved successfully!")
				}
			}

		case 2:

			fmt.Print("Enter Item Name: ")

			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			item, err := restaurant.FindItemByName(name)

			if err != nil {
				fmt.Println("Error:", err)
			} else {

				fmt.Println("\n===== ITEM FOUND =====")
				fmt.Println("Item ID    :", item.ItemID)
				fmt.Println("Name       :", item.Name)
				fmt.Println("Category   :", item.Category)
				fmt.Println("Price      :", item.Price)
				fmt.Println("Vegetarian :", item.IsVegetarian)
			}

		case 3:

			fmt.Print("Enter Item ID: ")

			input, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(input))

			if err != nil {
				fmt.Println("Invalid Item ID.")
				continue
			}

			fmt.Print("Enter New Price: ")

			input, _ = reader.ReadString('\n')
			price, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

			if err != nil {
				fmt.Println("Invalid price.")
				continue
			}

			err = restaurant.UpdatePrice(id, price)

			if err != nil {
				fmt.Println("Error:", err)
			} else {

				restaurant.SaveMenu(filename)

				fmt.Println("Price updated and saved successfully!")
			}

		case 4:

			fmt.Print("Enter Item ID to remove: ")

			input, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(input))

			if err != nil {
				fmt.Println("Invalid Item ID.")
				continue
			}

			err = restaurant.RemoveMenuItem(id)

			if err != nil {
				fmt.Println("Error:", err)
			} else {

				restaurant.SaveMenu(filename)

				fmt.Println("Item removed and saved successfully!")
			}

		case 5:

			restaurant.DisplayMenu()

		case 6:

			restaurant.SaveMenu(filename)

			fmt.Println("Data saved successfully.")
			fmt.Println("Thank you!")

			return

		default:

			fmt.Println("Invalid choice.")
		}
	}
}
