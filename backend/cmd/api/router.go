package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"go.uber.org/zap"

	"gitlab.praktikum-services.ru/Stasyan/momo-store/cmd/api/app"
	"gitlab.praktikum-services.ru/Stasyan/momo-store/internal/logger"
)

func newRouter(app *app.Instance) (http.Handler, error) {
	r := chi.NewRouter()

	corses := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	r.Use(
		middleware.StripSlashes,
		logMiddleware,
		corses.Handler,
	)

	r.Group(func(r chi.Router) {
		r.Use(
			app.TimingsMiddleware,
			app.RequestsMiddleware,
		)

		r.Get("/products", app.ListDumplingsController)
		r.Get("/categories", app.ListCategoriesController)
		r.Post("/orders", app.CreateOrderController)

		r.Get("/auth/whoami", app.WhoAmIController)
	})

	r.Get("/health", app.HealthcheckController)
	r.Method(http.MethodGet, "/metrics", app.MetricsHandler())

	return r, nil
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Debug("got request",
			zap.String("method", r.Method),
			zap.String("uri", r.RequestURI),
		)
		next.ServeHTTP(w, r)
	})
}
