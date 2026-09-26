package main
import (
	"fmt"
	"restaurant-menu/restaurant"
)
func main() {
	menu := restaurant.Menu{}
	item1 := restaurant.MenuItem{
		ItemID:       1,
		Name:         "Dosa",
		Category:     "Breakfast",
		Price:        50,
		IsVegetarian: true,
	}

	item2 := restaurant.MenuItem{
		ItemID:       2,
		Name:         "Chicken Biryani",
		Category:     "Main Course",
		Price:        180,
		IsVegetarian: false,
	}
	menu.AddMenuItem(item1)
	menu.AddMenuItem(item2)
	fmt.Println("Menu items added successfully")
	item, err := menu.FindItemByName("chicken biryani")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found Item:", item.Name)
		fmt.Println("Price:", item.Price)
	}
	err = menu.UpdatePrice(1, 60)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Price updated successfully")
	}
	err = menu.RemoveMenuItem(2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Menu item removed successfully")
	}
}
