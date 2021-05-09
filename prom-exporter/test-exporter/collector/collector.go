package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type smaemCollector struct {
	nfiMetric *prometheus.Desc
	tsiMetric *prometheus.Desc
}

type SmaemCollector interface {
	Describe(ch chan<- *prometheus.Desc)
	Collect(ch chan<- prometheus.Metric)
}
//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func NewSmaemCollector() SmaemCollector {
	return &smaemCollector{
		nfiMetric: prometheus.NewDesc("nfi_metric",
			"Shows whether a nfi has occurred in our cluster",
			nil, nil,
		),
		tsiMetric: prometheus.NewDesc("tsi_metric",
			"Shows whether a tsi has occurred in our cluster",
			nil, nil,
		),
	}
}
type testCollector struct {
	TestID	string
	TestMetric *prometheus.Desc
	Calls float64
}

type TestCollector interface {
	Describe(ch chan<- *prometheus.Desc)
	Collect(ch chan<- prometheus.Metric)
}
func NewTestCollector(id string, descName string) TestCollector{
	return &testCollector{
		TestID: id,
		TestMetric: prometheus.NewDesc(descName,
			"Shows test_metric to test multiple collector registration",
			nil, nil,
		),
		Calls: 0,
	}
}
//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *smaemCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.nfiMetric
	ch <- collector.tsiMetric
}

//Collect implements required collect function for all promehteus collectors
func (collector *smaemCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	var metricValue float64
	if 1 == 1 {
		metricValue = 1
	}

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	ch <- prometheus.MustNewConstMetric(collector.nfiMetric, prometheus.CounterValue, metricValue)
	ch <- prometheus.MustNewConstMetric(collector.tsiMetric, prometheus.CounterValue, metricValue)

}

func (c *testCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- c.TestMetric
	//ch <- c.tsiMetric
}

//Collect implements required collect function for all promehteus collectors
func (c *testCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	var metricValue float64
	if strings.Contains(c.TestID, "a1ccb23497ab") {
		metricValue = 11111+c.Calls
	} else { metricValue = 22222+c.Calls }
	c.Calls+=1
	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	ch <- prometheus.MustNewConstMetric(c.TestMetric, prometheus.CounterValue, metricValue)
	//ch <- prometheus.MustNewConstMetric(c.tsiMetric, prometheus.CounterValue, metricValue)

}