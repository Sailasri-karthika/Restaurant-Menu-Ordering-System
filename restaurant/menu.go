package restaurant

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type MenuItem struct {
	ItemID       int
	Name         string
	Category     string
	Price        float64
	IsVegetarian bool
}

var items []MenuItem

func AddMenuItem(item MenuItem) error {

	for _, i := range items {
		if i.ItemID == item.ItemID {
			return errors.New("item ID already exists")
		}
	}

	if strings.TrimSpace(item.Name) == "" {
		return errors.New("item name cannot be empty")
	}

	if item.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	items = append(items, item)

	return nil
}

func FindItemByName(name string) (*MenuItem, error) {

	for i := range items {
		if strings.EqualFold(items[i].Name, name) {
			return &items[i], nil
		}
	}

	return nil, errors.New("item not found")
}

func UpdatePrice(id int, newPrice float64) error {

	if newPrice <= 0 {
		return errors.New("price must be greater than 0")
	}

	for i := range items {
		if items[i].ItemID == id {
			items[i].Price = newPrice
			return nil
		}
	}

	return errors.New("item not found")
}

func RemoveMenuItem(id int) error {

	for i := range items {
		if items[i].ItemID == id {

			items = append(items[:i], items[i+1:]...)

			return nil
		}
	}

	return errors.New("item not found")
}

func DisplayMenu() {

	if len(items) == 0 {
		fmt.Println("No menu items available.")
		return
	}

	fmt.Println("\n===== ALL MENU ITEMS =====")

	for _, item := range items {

		fmt.Println("-------------------------")
		fmt.Println("Item ID    :", item.ItemID)
		fmt.Println("Name       :", item.Name)
		fmt.Println("Category   :", item.Category)
		fmt.Println("Price      :", item.Price)
		fmt.Println("Vegetarian :", item.IsVegetarian)
	}
}

func LoadMenu(filename string) error {

	file, err := os.Open(filename)

	if err != nil {

		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	defer file.Close()

	items = nil

	err = json.NewDecoder(file).Decode(&items)

	if err != nil {
		return err
	}

	return nil
}

func SaveMenu(filename string) error {

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "    ")

	return encoder.Encode(items)
}
