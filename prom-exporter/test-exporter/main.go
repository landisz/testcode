package main

import (
  "net/http"
  "time"
  "math/rand"
  "fmt"

  log "github.com/sirupsen/logrus"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
  "github.com/prometheus/client_golang/prometheus/promhttp"
  c "testcode/prom-exporter/test-exporter/collector"
)

// define some auto metrics
var (
  thresholdDelay = 13.5
  counterInfo = "SDN_Delay_Over_"+ftostr(thresholdDelay)+"ms_total"
  countProcessed = promauto.NewCounter(prometheus.CounterOpts{
    Name: "MYAPP_DELAY_MONITOR_______",
    Help: "The total number received events",
  })
  countDelay = promauto.NewCounter(prometheus.CounterOpts{
    Name: "MYAPP_DELAY_COUNTER",
    Help: "The total number "+counterInfo+" events",
  })
  gaugeProcessed = promauto.NewGauge(prometheus.GaugeOpts{
    Name: "MYAPP_DELAY_VALUE",
    Help: "The total number of processed events",
  })
)

func ftostr(f float64) string {
  return fmt.Sprintf("%f", f)
}


func metricsMonitor() {
  //go func() {
  //  for {
  //    //auto increase counter
  //    countProcessed.Inc()
  //    time.Sleep(2 * time.Second)
  //  }
  //}()
  countProcessed.Inc()
}
func metricsCounter() {
  countDelay.Inc()
}
func metricsGauge() {
  go func() {
    for {
      // simulate random delay between 10.0-15.999 every 2 seconds
      fakeDelay := rand.Float64()*5+rand.Float64()+10
      gaugeProcessed.Set(fakeDelay)
      time.Sleep(2 * time.Second)
      metricsMonitor()
      if fakeDelay>thresholdDelay {
        metricsCounter()
      }
    }
  }()
}


func main() {
  rand.Seed(time.Now().UnixNano())
  // start simulating metrics
  //metricsCounter()
  //metricsGauge()

  //Create a new instance of the foocollector and
  //register it with the prometheus client.
  smaemCollector := c.NewSmaemCollector()
  testCollector1 := c.NewTestCollector("669652d0-89e3-4aa4-81fd-a1ccb23497ab", "Metrics_for_ID_a1ccb23497ab:")
  testCollector2 := c.NewTestCollector("08bb4fe9-f497-4e78-b153-d34cc97feaaf", "Metrics_for_ID_d34cc97feaaf:")

  prometheus.MustRegister(smaemCollector)
  prometheus.MustRegister(testCollector1)
  prometheus.MustRegister(testCollector2)


  //This section will start the HTTP server and expose
  //any metrics on the /metrics endpoint.
  http.Handle("/metrics", promhttp.Handler())
  log.Info("Beginning to serve on port :9414")
  log.Fatal(http.ListenAndServe(":9414", nil))
}
