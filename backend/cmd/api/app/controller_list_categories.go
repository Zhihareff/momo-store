package app

import (
	"encoding/json"
	"net/http"
)

func (i *Instance) ListCategoriesController(w http.ResponseWriter, r *http.Request) {
	type category struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	categories := []category{
		{ID: 1, Name: "Сырые", Description: "Для особых ценителей"},
		{ID: 2, Name: "Аль денте", Description: "В лучших традициях итальянской кухни"},
		{ID: 3, Name: "Сваренные", Description: "Неувядающая классика"},
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(categories)
}
