package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "chess.io_requests_total",
			Help: "Total number of requests processed by the web server.",
		},
		[]string{"path", "status"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "chess.io_requests_errors_total",
			Help: "Total number of error requests processed by the web server.",
		},
		[]string{"path", "status"},
	)
)

func PromInit() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)
}

func TrackMetrics() fiber.Handler {
	return func(c fiber.Ctx) error {
		path := c.Path()
		err := c.Next()
		status := c.Response().StatusCode()
		RequestCount.WithLabelValues(path, strconv.Itoa(status)).Inc()
		if status >= 400 {
			ErrorCount.WithLabelValues(path, strconv.Itoa(status)).Inc()
		}
		return err
	}
}
