package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"strings"
	"time"
)

var (
	device string = "eth0"
	snapshotLen int32 = 1024
	promiscuous bool = false
	err error
	timeout time.Duration = 30 * time.Second //若为负数，表示不缓存，直接输出
	handle      *pcap.Handle
)

func main() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	var filter string = "tcp" //BPF过滤语法
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		printPacketInfo(packet)
		time.Sleep(1 * time.Second)
	}
}

//常用的包层
//packet.LinkLayer() 以太网
//packet.NetworkLayer() 网络层，通常 也就是 IP 层
//packet.TransportLayer() 传输层，比如 TCP/UDP
//packet.ApplicationLayer() 应用层，比如 HTTP 层。
//packet.ErrorLayer() ……出错了

func printPacketInfo(packet gopacket.Packet) {
	// an ethernet packet
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("Ethernet layer detected.")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Println("BaseLayer Contents: ", ethernetPacket.BaseLayer.Contents)
		fmt.Println("BaseLayer Payload: ", ethernetPacket.BaseLayer.Payload)
		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		// Ethernet type is typically IPv4 but could be ARP or other
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
		fmt.Println()
	}

	//  IP
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 layer detected.")
		ip, _ := ipLayer.(*layers.IPv4)
		// IP layer variables:
		fmt.Println("BaseLayer Contents: ", ip.BaseLayer.Contents)
		fmt.Println("BaseLayer Payload: ", ip.BaseLayer.Payload)
		// Version (Either 4 or 6)
		fmt.Println("Version: ", ip.Version)
		// IHL (IP Header Length in 32-bit words)
		fmt.Println("IHL: ", ip.IHL)
		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
		fmt.Println("TOS: ", ip.TOS)
		fmt.Println("Length: ", ip.Length)
		fmt.Println("Id: ", ip.Id)
		fmt.Println("Flags: ", ip.Flags)
		fmt.Println("FragOffset: ", ip.FragOffset)
		fmt.Println("TTL: ", ip.TTL)
		fmt.Println("Protocol: ", ip.Protocol)
		// Checksum, SrcIP, DstIP
		fmt.Println("Checksum: ", ip.Checksum)
		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Options: ", ip.Options)
		fmt.Println("Padding: ", ip.Padding)
		fmt.Println()
	}

	//  TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected.")
		tcp, _ := tcpLayer.(*layers.TCP)
		// TCP layer variables:
		fmt.Println("BaseLayer Contents: ", tcp.BaseLayer.Contents)
		fmt.Println("BaseLayer Payload: ", tcp.BaseLayer.Payload)
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence number: ", tcp.Seq)
		fmt.Println("Ack: ", tcp.Ack)
		fmt.Println("DataOffset: ", tcp.DataOffset)
		fmt.Println("Window: ", tcp.Window)
		fmt.Println("Checksum: ", tcp.Checksum)
		fmt.Println("Urgent: ", tcp.Urgent)
		fmt.Println("Options: ", tcp.Options)
		fmt.Println()
	}

	// Iterate over all layers, printing out each layer type
	fmt.Println("All packet layers:")
	for _, layer := range packet.Layers() {
		fmt.Println("- ", layer.LayerType())
	}
	fmt.Println()

	// When iterating through packet.Layers() above,
	// if it lists Payload layer then that is the same as
	// this applicationLayer. applicationLayer contains the payload
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("Application layer/Payload found.")
		fmt.Println(string(applicationLayer.Payload()))
		// Search for a string inside the payload
		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
			fmt.Println("HTTP found!")
		}
	}

	// Check for errors
	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet:", err)
	} else {
		fmt.Println("NO Error")
	}
}