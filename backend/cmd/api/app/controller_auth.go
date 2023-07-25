package app

import (
	"encoding/json"
	"net/http"
)

func (i *Instance) WhoAmIController(w http.ResponseWriter, r *http.Request) {
	type user struct {
		ID        int64  `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).
		Encode(user{
			ID:        1,
			FirstName: "Иван",
			LastName:  "Иванов",
			Email:     "momolover@mail.ru",
		})
}
