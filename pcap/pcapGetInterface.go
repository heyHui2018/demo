package main

//获取当前设备所有接口信息

import (
	"fmt"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}

	// Print device information
	for _, device := range devices {
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices addresses: ", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
}
