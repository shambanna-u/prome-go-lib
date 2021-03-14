package main

import (
	"fmt"
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
	dynamicLabels = []string{"url"}
	responseTime  = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sample_external_url_response_ms",
			Help: "Services http response.",
		},
		dynamicLabels,
	)
	urlUp = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sample_external_url_up",
			Help: "Services http response.",
		},
		dynamicLabels,
	)
)

func main() {
	registry := prometheus.NewRegistry()
	_ = registry.Register(responseTime)
	_ = registry.Register(urlUp)
	go runAlways()
	gwHandler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	http.Handle("/metrics", gwHandler)
	log.Fatal(http.ListenAndServe(":2222", nil))
}

func runAlways() {
	for {
		urls := []string{"https://httpstat.us/503", "https://httpstat.us/200"}
		for _, url := range urls {
			instrumentedHandler(url)
		}
	}
}

func instrumentedHandler(url string) int {

	fmt.Println("hello " + url)
	start := time.Now()
	labels := prometheus.Labels{"url": url}
	resp, _ := http.Get(url)
	fmt.Println(resp.StatusCode)
	respCode, _ := strconv.Atoi(url[strings.LastIndex(url, "/")+1:])
	if respCode != 0 {
		fmt.Println(respCode)
		if respCode == 200 {
			if resp.StatusCode == 200 {
				urlUp.With(labels).Set(1)
			} else {
				urlUp.With(labels).Set(0)
			}
		} else {
			if resp.StatusCode == 503 {
				urlUp.With(labels).Set(1)
			} else {
				urlUp.With(labels).Set(0)
			}
		}
	} else {
		if resp.StatusCode == 200 {
			urlUp.With(labels).Set(1)
		} else {
			urlUp.With(labels).Set(0)
		}

	}
	_, _ = ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	elapsed := float64(time.Since(start).Milliseconds())

	responseTime.With(labels).Set(elapsed)
	time.Sleep(time.Second * 10)
	return resp.StatusCode
}
