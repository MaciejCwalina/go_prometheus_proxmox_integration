package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	guage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxmox_status_info",
			Help: "Gets information about proxmox",
		},
		[]string{"name", "id", "status"},
	)
)

func RecordMetrics(systemManager *SystemManager) {
	go func() {
		for {
			infraData, err := systemManager.GetInfrastructureData()
			if err != nil {
				log.Fatal("Failed to Get Infrastructure Information", err.Error())
			}

			guage.Reset()
			systemManager.PopulateSlices(infraData)
			for _, vm := range systemManager.vms {
				var isRunning float64
				if vm.Status == "running" {
					isRunning = 1
				} else {
					isRunning = 0
				}

				guage.WithLabelValues(vm.Name, vm.ID, vm.Status).Set(isRunning)
			}

			time.Sleep(5 * time.Second)
		}
	}()
}

func main() {
	systemManager := SystemManager{}
	RecordMetrics(&systemManager)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":2112", nil))
}
