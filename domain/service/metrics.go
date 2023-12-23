package service

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sbrunettajr/ifoody-api/domain/constant"
)

type MetricsService struct {
	ordersTotal   prometheus.Counter
	requestsTotal *prometheus.CounterVec
}

func NewMetricsService() MetricsService {
	ordersTotal := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: constant.OrdersTotal,
			Help: "Total of orders",
		},
	)
	prometheus.Register(ordersTotal)

	requestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: constant.HTTPRequestsTotal,
			Help: "Total of HTTP requests",
		},
		[]string{"method", "path", "statusCode"},
	)
	prometheus.Register(requestsTotal)

	return MetricsService{
		ordersTotal:   ordersTotal,
		requestsTotal: requestsTotal,
	}
}

func (s MetricsService) RegisterOrder() {
	s.ordersTotal.Inc()
}

func (s MetricsService) RegisterRequest(method, path, statusCode string) {
	s.requestsTotal.WithLabelValues(method, path, statusCode).Inc()
}
