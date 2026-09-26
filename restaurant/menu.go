package restaurant

import (
	"errors"
	"strings"
)

type MenuItem struct {
	ItemID       int
	Name         string
	Category     string
	Price        float64
	IsVegetarian bool
}

type Menu struct {
	items []MenuItem
}
// AddMenuItem adds a new item to the menu
func (m *Menu) AddMenuItem(item MenuItem) error {

	if strings.TrimSpace(item.Name) == "" {
		return errors.New("item name cannot be empty")
	}

	if item.Price <= 0 {
		return errors.New("price must be positive")
	}

	m.items = append(m.items, item)

	return nil
}

// FindItemByName searches for an item by name
func (m *Menu) FindItemByName(name string) (*MenuItem, error) {

	for i := range m.items {
		if strings.EqualFold(m.items[i].Name, name) {
			return &m.items[i], nil
		}
	}

	return nil, errors.New("item not found")
}

// UpdatePrice updates the price of an item
func (m *Menu) UpdatePrice(id int, newPrice float64) error {

	if newPrice <= 0 {
		return errors.New("price must be positive")
	}

	for i := range m.items {
		if m.items[i].ItemID == id {
			m.items[i].Price = newPrice
			return nil
		}
	}

	return errors.New("item not found")
}

// RemoveMenuItem removes an item from the menu
func (m *Menu) RemoveMenuItem(id int) error {

	for i := range m.items {
		if m.items[i].ItemID == id {

			m.items = append(m.items[:i], m.items[i+1:]...)

			return nil
		}
	}

	return errors.New("item not found")
}