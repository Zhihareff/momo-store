package app

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"gitlab.praktikum-services.ru/Stasyan/momo-store/internal/store/dumplings"
)

// Instance represents API app instance
type Instance struct {
	store dumplings.Store

	// metrics
	metricsRegistry         *prometheus.Registry
	responseTimings         *prometheus.HistogramVec
	ordersCounter           prometheus.Counter
	requestCounter          prometheus.Counter
	dumplingsListingCounter *prometheus.CounterVec
}

// NewInstance returns new app instance
func NewInstance(store dumplings.Store) (*Instance, error) {
	metricsRegistry := prometheus.NewRegistry()

	responseTimingsHistogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "response_timing_ms",
			Help:    "Response timings in milliseconds",
			Buckets: prometheus.LinearBuckets(0, 50, 10),
		},
		[]string{"handler"},
	)

	dumplingsListingCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "dumplings_listing_count",
			Help: "Number of times dumplings pack has been listed",
		},
		[]string{"id"},
	)

	ordersCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "orders_count",
			Help: "Number of dumplings orders",
		},
	)
	requestCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "requests_count",
			Help: "Number of HTTP requests made",
		},
	)

	// register metrics
	metricsRegistry.MustRegister(responseTimingsHistogram)
	metricsRegistry.MustRegister(ordersCounter)
	metricsRegistry.MustRegister(requestCounter)
	metricsRegistry.MustRegister(dumplingsListingCounter)

	return &Instance{
		store: store,

		metricsRegistry:         metricsRegistry,
		responseTimings:         responseTimingsHistogram,
		ordersCounter:           ordersCounter,
		requestCounter:          requestCounter,
		dumplingsListingCounter: dumplingsListingCounter,
	}, nil
}

func (i *Instance) HealthcheckController(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (i *Instance) MetricsHandler() http.Handler {
	return promhttp.HandlerFor(
		i.metricsRegistry,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	)
}
