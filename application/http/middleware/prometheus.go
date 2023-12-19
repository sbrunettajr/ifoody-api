package middleware

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sbrunettajr/ifoody-api/domain/service"
)

type prometheusMiddleware struct {
	metricsService service.MetricsService
}

func NewPrometheusMiddleware(
	metricsService service.MetricsService,
) prometheusMiddleware {
	return prometheusMiddleware{
		metricsService: metricsService,
	}
}

func (m prometheusMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		statusCode := strconv.Itoa(c.Response().Status)
		m.metricsService.RegisterRequest(
			c.Request().Method,
			c.Request().URL.Path,
			statusCode,
		)

		return err
	}
}
