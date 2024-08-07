package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RecordMetrics(systemManager *SystemManager) {
	go func() {
		for {
			startedTime := time.Now()
			infraData := systemManager.GetInfrastructureData()
			ResetGuages()
			systemManager.PopulateSlices(infraData)
			for _, vm := range systemManager.vms {
				SetGuages(&vm)
			}

			proxmoxScrapingTime.Set(float64(time.Since(startedTime)))
			time.Sleep(5 * time.Second)
		}
	}()
}

func main() {
	RecordMetrics(&SystemManager{})
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":2112", nil))
}
