package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	//  labels to print url in prometheus output
	dynamicLabels = []string{"url"}
	// GaugeVec is a Collector that bundles a set of Gauges that all share the same Desc, but have different values for their variable labels.
	//for response time
	responseTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sample_external_url_response_ms",
			Help: "Services http response.",
		},
		dynamicLabels,
	)
	// for service status
	urlUp = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sample_external_url_up",
			Help: "Services http status.",
		},
		dynamicLabels,
	)
)

func main() {
	// NewRegistry creates a new vanilla Registry without any Collectors pre-registered.
	// this avoid the default metrics
	registry := prometheus.NewRegistry()
	//Registry registers Prometheus collectors, collects their metrics,
	_ = registry.Register(responseTime)
	_ = registry.Register(urlUp)
	//this will call endpoints and collect the metrics
	go runAlways()
	// Thus, HandlerFor is useful to create http.Handlers for custom Gatherers, with non-default HandlerOpts,
	gwHandler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	//start the server with /metircs endpoint
	http.Handle("/metrics", gwHandler)
	//server will host the metrics with 2222 port
	log.Fatal(http.ListenAndServe(":2222", nil))
}

// this is the continuous loop wich calls the endpoints
func runAlways() {
	for {
		urls := []string{"https://httpstat.us/503", "https://httpstat.us/200"}
		for _, url := range urls {
			instrumentedHandler(url)
		}
	}
}

func instrumentedHandler(url string) int {
	// start the time to calcalute response time
	start := time.Now()
	// set the labels
	labels := prometheus.Labels{"url": url}
	// call the endpoints
	resp, _ := http.Get(url)
	//as https://httpstat.us/503 is always returns 503 to get count we are extracting expected response code from url
	respCode, _ := strconv.Atoi(url[strings.LastIndex(url, "/")+1:])
	if respCode != 0 {
		// we are calling https://httpstat.us/200
		if respCode == 200 {
			if resp.StatusCode == 200 {
				//set value is 1 if its up
				urlUp.With(labels).Set(1)
			} else {
				urlUp.With(labels).Set(0)
			}
		} else {
			// if we are calling https://httpstat.us/503
			if resp.StatusCode == 503 {
				//set value is 1 if its up
				urlUp.With(labels).Set(1)
			} else {
				urlUp.With(labels).Set(0)
			}
		}
	} else {
		// if we are calling other than https://httpstat.us/503 and https://httpstat.us/200
		if resp.StatusCode == 200 {
			urlUp.With(labels).Set(1)
		} else {
			urlUp.With(labels).Set(0)
		}

	}
	//read the response
	_, _ = ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	//calculate the end of the time
	elapsed := float64(time.Since(start).Milliseconds())
	//set the time to metics
	responseTime.With(labels).Set(elapsed)
	// wait for 10 sec
	time.Sleep(time.Second * 10)
	return resp.StatusCode
}
