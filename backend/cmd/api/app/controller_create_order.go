package app

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"gitlab.praktikum-services.ru/Stasyan/momo-store/internal/logger"
	"gitlab.praktikum-services.ru/Stasyan/momo-store/internal/store/dumplings"
)

func (i *Instance) CreateOrderController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// TODO: add order items later
	var items []dumplings.OrderItem
	id, err := i.store.CreateOrder(ctx, items...)
	if err != nil {
		logger.Log.Error("cannot create order", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).
		Encode(map[string]interface{}{
			"id": id,
		})

	i.ordersCounter.Inc()
}
