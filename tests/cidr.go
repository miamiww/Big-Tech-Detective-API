package main

import (
	"fmt"
	"net"
    "encoding/csv"
    "os"
)

func main() {
	ranger := cidranger.NewPCTrieRanger()

	_, network1, _ := net.ParseCIDR("192.168.1.0/24")
	_, network2, _ := net.ParseCIDR("128.168.1.0/24")
	_, network3, _ := net.ParseCIDR("128.0.0.0/5")
	ranger.Insert(cidranger.NewBasicRangerEntry(*network1))
	ranger.Insert(cidranger.NewBasicRangerEntry(*network2))
	ranger.Insert(cidranger.NewBasicRangerEntry(*network3))

	containingNetworks, err := ranger.ContainingNetworks(net.ParseIP("128.168.1.0"))

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, network := range containingNetworks {
		connected := network.Network()
		fmt.Printf("Containing networks: %s\n", connected.String())
	}

	// Prints out
	// λ go run main.go
	// Containing networks: 128.0.0.0/5
	// Containing networks: 128.168.1.0/24

}
