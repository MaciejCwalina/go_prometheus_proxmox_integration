package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RecordMetrics(systemManager *SystemManager) {
	go func() {
		for {
			startedTime := time.Now()
			infraData := systemManager.GetInfrastructureData()
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
				proxmoxMaxVMMemory.WithLabelValues(vm.Name).Set(float64(vm.MaxMem))
				proxmoxUsedVMMemory.WithLabelValues(vm.Name).Set(float64(vm.Mem))
				proxmoxNetworkIn.WithLabelValues(vm.Name).Set(float64(vm.NetIn))
				proxmoxNetworkOut.WithLabelValues(vm.Name).Set(float64(vm.NetOut))
				proxmoxUpTime.WithLabelValues(vm.Name).Set(float64(vm.Uptime))
			}

			proxmoxScrapingTime.Set(float64(time.Since(startedTime)))
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
