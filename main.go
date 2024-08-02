package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	proxmoxStatusInfoGuage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxmox_status_info",
			Help: "Gets information about proxmox",
		},

		[]string{"name", "id", "status"},
	)

	proxmoxVMCpuInfo = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxmox_vm_cpu_info",
			Help: "Gets the information about the CPU usage in vms",
		},

		[]string{"name", "cpu_usage", "max_cores"},
	)

	proxmoxDiskRead = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxmox_disk_read",
			Help: "Gets the disk read of vms",
		},

		[]string{"name"},
	)

	proxmoxDiskWrite = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxmox_disk_write",
			Help: "Gets the disk write of vms",
		},

		[]string{"name"},
	)
)

func RecordMetrics(systemManager *SystemManager) {
	go func() {
		for {
			infraData, err := systemManager.GetInfrastructureData()
			if err != nil {
				log.Fatal("Failed to Get Infrastructure Information", err.Error())
			}

			proxmoxStatusInfoGuage.Reset()
			proxmoxVMCpuInfo.Reset()
			systemManager.PopulateSlices(infraData)
			for _, vm := range systemManager.vms {
				var isRunning float64
				if vm.Status == "running" {
					isRunning = 1
				} else {
					isRunning = 0
				}

				proxmoxStatusInfoGuage.WithLabelValues(vm.Name, vm.ID, vm.Status).Set(isRunning)
				proxmoxVMCpuInfo.WithLabelValues(vm.Name, strconv.FormatFloat(vm.CPU, 'E', -1, 64), strconv.Itoa(vm.MaxCPU)).Set(vm.CPU)
				proxmoxDiskRead.WithLabelValues(vm.Name).Set(float64(vm.DiskRead))
				proxmoxDiskWrite.WithLabelValues(vm.Name).Set(float64(vm.DiskWrite))
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
