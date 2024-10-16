package models

import "heroes-cube-web/internal/database"

type CharacterViewModel struct {
	ID     int64
	Name   string
	Race   string
	Class  string
	Damage int64
	Level  int64
	Points int64 // You'll need to fetch this from the external API
}

func CharacterToViewModel(c database.Character) CharacterViewModel {
	return CharacterViewModel{
		ID:     c.ID,
		Name:   c.Name,
		Race:   c.Race,
		Class:  c.Class,
		Damage: c.Damage,
		Level:  c.Level,
		Points: fetchPointsFromAPI(c.ID), // Implement this function
	}
}

type InventoryViewModel struct {
	ID       int64
	Name     string
	Damage   int64
	Price    int64
	Class    string
	Quantity int64
}

func InventoryToViewModel(inventory []database.Inventory) []InventoryViewModel {
	var result []InventoryViewModel
	for _, item := range inventory {
		result = append(result, InventoryViewModel{
			ID:       item.ID,
			Name:     item.Name,
			Damage:   item.Damage,
			Price:    item.Price,
			Class:    item.Class,
			Quantity: item.Quantity,
		})
	}
	return result
}

type SlotViewModel struct {
	Slot   string
	Name   string
	Damage int64
	Price  int64
	Class  string
}

func SlotsToViewModel(slots []database.Slot) []SlotViewModel {
	var result []SlotViewModel
	for _, slot := range slots {
		result = append(result, SlotViewModel{
			Slot:   slot.Slot,
			Name:   slot.Name,
			Damage: slot.Damage,
			Price:  slot.Price,
			Class:  slot.Class,
		})
	}
	return result
}

// You'll need to implement this function to fetch points from your external API
func fetchPointsFromAPI(characterID int64) int64 {
	// Implement the API call here
	return 0 // Placeholder return
}
