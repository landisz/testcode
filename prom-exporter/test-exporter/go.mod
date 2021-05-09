module testcode/prom-exporter/test-expoter/v2

go 1.14

replace testcode/prom-exporter/test-exporter/collector => ./collector

require (
	github.com/prometheus/client_golang v1.10.0
	github.com/sirupsen/logrus v1.8.1
	testcode/prom-exporter/test-exporter/collector v0.0.0-00010101000000-000000000000
)
