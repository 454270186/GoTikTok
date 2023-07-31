package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		httpRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
		c.Next()
	}
}