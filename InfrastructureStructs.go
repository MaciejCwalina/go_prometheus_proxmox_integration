package main

type VirtualMachine struct {
	CPU       float64 `json:"cpu"`
	Disk      int64   `json:"disk"`
	DiskRead  int64   `json:"diskread"`
	DiskWrite int64   `json:"diskwrite"`
	ID        string  `json:"id"`
	MaxCPU    int     `json:"maxcpu"`
	MaxDisk   int64   `json:"maxdisk"`
	MaxMem    int64   `json:"maxmem"`
	Mem       int64   `json:"mem"`
	Name      string  `json:"name"`
	NetIn     int64   `json:"netin"`
	NetOut    int64   `json:"netout"`
	Node      string  `json:"node"`
	Status    string  `json:"status"`
	Template  int     `json:"template"`
	Type      string  `json:"type"`
	Uptime    int64   `json:"uptime"`
	VMID      int     `json:"vmid"`
}

type Node struct {
	CPU     float64 `json:"cpu"`
	Disk    int64   `json:"disk"`
	ID      string  `json:"id"`
	Level   string  `json:"level"`
	MaxCPU  int     `json:"maxcpu"`
	MaxDisk int64   `json:"maxdisk"`
	MaxMem  int64   `json:"maxmem"`
	Mem     int64   `json:"mem"`
	Node    string  `json:"node"`
	Status  string  `json:"status"`
	Type    string  `json:"type"`
	Uptime  int64   `json:"uptime"`
}

type Storage struct {
	Content    string `json:"content"`
	Disk       int64  `json:"disk"`
	ID         string `json:"id"`
	MaxDisk    int64  `json:"maxdisk"`
	Node       string `json:"node"`
	PluginType string `json:"plugintype"`
	Shared     int    `json:"shared"`
	Status     string `json:"status"`
	Storage    string `json:"storage"`
	Type       string `json:"type"`
}
