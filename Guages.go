package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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

	proxmoxMaxVMMemory = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxmox_vm_max_memory",
			Help: "Gets the memory information from a VM",
		},

		[]string{"name"},
	)

	proxmoxUsedVMMemory = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxmox_vm_used_memory",
			Help: "Gets the memory information from a VM",
		},

		[]string{"name"},
	)

	proxmoxNetworkOut = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxmox_vm_net_out",
			Help: "Gets the network output of a all the VMs",
		},

		[]string{"name"},
	)

	proxmoxNetworkIn = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "proxmox_vm_net_in",
			Help: "Gets the network input of a all the VMs",
		},

		[]string{"name"},
	)

	proxmoxScrapingTime = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "proxmox_scrape_time",
			Help: "Sends the time it took to scrape data from proxmox and the vms",
		},
	)
)
