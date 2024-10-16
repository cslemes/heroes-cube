package handlers

import (
	"heroes-cube-web/internal/database"
	"heroes-cube-web/internal/models"
	"heroes-cube-web/internal/templates"
	"net/http"
)

func Characters(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			if err := templates.RenderCharacters(w, nil); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}

		character, err := db.GetCharacter(r.Context(), name)
		if err != nil {
			http.Error(w, "Character not found", http.StatusNotFound)
			return
		}

		inventory, err := db.GetInventory(r.Context(), character.ID)
		if err != nil {
			http.Error(w, "Failed to fetch inventory", http.StatusInternalServerError)
			return
		}

		slots, err := db.GetSlots(r.Context(), character.ID)
		if err != nil {
			http.Error(w, "Failed to fetch slots", http.StatusInternalServerError)
			return
		}

		data := struct {
			Character models.CharacterViewModel
			Inventory []models.InventoryViewModel
			Slots     []models.SlotViewModel
		}{
			Character: models.CharacterToViewModel(character),
			Inventory: models.InventoryToViewModel(inventory),
			Slots:     models.SlotsToViewModel(slots),
		}

		if err := templates.RenderCharacters(w, data); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
