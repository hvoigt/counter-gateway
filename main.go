package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strings"
)

var counters map[string]*prometheus.CounterVec

func init() {
	counters = make(map[string]*prometheus.CounterVec)
}

func ensureCounterExists(counterName string, labels map[string]string) *prometheus.CounterVec {
	if counter, ok := counters[counterName]; ok {
		return counter
	}

	labelsArray := []string{}
	for key := range labels {
		labelsArray = append(labelsArray, key)
	}

	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: counterName,
		Help: "A counter metric that increments when /increment is called",
	}, labelsArray)
	prometheus.MustRegister(counter)
	counters[counterName] = counter
	return counter
}

func parseLabels(label_parameters []string) map[string]string {
	labels := make(map[string]string)
	for _, label := range label_parameters {
		parts := strings.Split(label, "=")
		if len(parts) != 2 {
			continue
		}
		labels[parts[0]] = parts[1]
	}
	return labels
}

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	counterName := r.URL.Query().Get("counter")
	labels := parseLabels(r.URL.Query()["label"])

	counter := ensureCounterExists(counterName, labels)

	counter.With(labels).Inc()

	fmt.Fprintf(w, "Ok")
}

func main() {
	http.HandleFunc("/increment", incrementHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
