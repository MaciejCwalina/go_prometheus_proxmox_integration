package main

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	// Import the pcap package to capture packets.

	"github.com/prometheus/client_golang/prometheus/promhttp"
	// Import the gopacket package to decode packets.
	// Import the layers package to access the various network layers.
)

func RecordMetrics(systemManager *SystemManager, packetsChannel chan []Packet) {
	go func() {
		for {
			startedTime := time.Now()
			SetNetworkGuages(<-packetsChannel)
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

func StartHttpServer() net.Listener {
	listener, err := net.Listen("tcp", ":5145")
	if err != nil {
		log.Fatal(err.Error())
	}

	return listener
}

func HandleIncomingConnections(listener net.Listener) net.Conn {
	conn, err := listener.Accept() //thread blocking
	if err != nil {
		log.Fatal("Failed to accept connection")
	}

	return conn
}

func ReadFileBytes(validDataChannel chan []byte, conn net.Conn) {
	go func() {
		for {
			byteArr := make([]byte, 1024)
			readLength, _ := conn.Read(byteArr)
			validDataChannel <- byteArr[:readLength]
		}
	}()
}

type Packet struct {
	Dest string
	Src  string
	Size int
}

// Very rough implementation.
// I have to polish this up
func ParseData(validDataChannel chan []byte, packetChannel chan []Packet) {
	go func() {
		for {
			packetInfo := string(<-validDataChannel)
			packetInfoSplit := strings.Split(packetInfo, "\n")
			packets := []Packet{}
			for _, packet := range packetInfoSplit {
				packetSplit := strings.Split(packet, " ")
				lengthPacketSplit := len(packetSplit)

				if lengthPacketSplit < 8 {
					continue
				}

				packetSize, err := strconv.Atoi(packetSplit[7])
				if err != nil {
					log.Println("Cannot parse PacketSize!")
				}

				dest := packetSplit[2]
				src := packetSplit[4]
				destSplit := strings.Split(dest, ".")
				srcSplit := strings.Split(src, ".")
				dest = ""
				src = ""
				for i := 0; i < len(destSplit); i++ {
					if i == len(destSplit)-1 {
						break
					}

					dest += destSplit[i] + "."
				}

				for i := 0; i < len(srcSplit); i++ {
					if i == len(srcSplit)-1 {
						break
					}

					src += srcSplit[i] + "."
				}

				packet := Packet{
					dest,
					src,
					packetSize,
				}

				packets = append(packets, packet)
			}

			packetChannel <- packets
		}
	}()
}

// Add something that displays how many bytes were sent and where and the country of origin if for example we sent 12TB of data to Russia raise and alert immieditly.
func main() {
	listener := StartHttpServer()
	conn := HandleIncomingConnections(listener)
	validDataChannel := make(chan []byte)
	packetChannel := make(chan []Packet)
	ReadFileBytes(validDataChannel, conn)
	ParseData(validDataChannel, packetChannel)
	RecordMetrics(&SystemManager{}, packetChannel)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":2112", nil))
}
