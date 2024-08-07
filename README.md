# Proxmox Monitoring with Prometheus

This project is a Go application that collects metrics from a Proxmox server and exposes them to Prometheus for monitoring. It uses the Prometheus client library to create and update metrics, and the `pvesh` command to retrieve data from Proxmox.

## Features

- Collects metrics for virtual machines (VMs), nodes, and storage
- Exposes metrics in the Prometheus format
- Supports multiple VMs, nodes, and storage types

## Metrics

The following metrics are collected and exposed:

- `proxmox_status_info`: Information about the status of VMs, nodes, and storage
- `proxmox_vm_cpu_info`: CPU usage and maximum cores for VMs
- `proxmox_disk_read`: Disk read for VMs
- `proxmox_disk_write`: Disk write for VMs
- `proxmox_vm_max_memory`: Maximum memory for VMs
- `proxmox_vm_used_memory`: Used memory for VMs
- `proxmox_vm_net_out`: Network output for VMs
- `proxmox_vm_net_in`: Network input for VMs
- `proxmox_scrape_time`: Time taken to scrape data from Proxmox
- `proxmox_vm_uptime`: Uptime for VMs
 ```bash
   go build -o proxmox-exporter
```
Run the application:
```bash
./proxmox-exporter
```
Access the metrics at http://<server_ip>:2112/metrics.
Configure Prometheus to scrape the metrics from the application.


## Configuration
The application uses the pvesh command to retrieve data from Proxmox. Make sure that pvesh is available in your system's PATH.
You can configure the scrape interval by modifying the time.Sleep() function in the RecordMetrics() function.
## Contributing
If you find any issues or have suggestions for improvement, feel free to open an issue or submit a pull request.
