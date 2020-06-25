package main

import (
    "encoding/csv"
    "encoding/json"
	"fmt"
	"net"
    "os"
	"github.com/miamiww/cidranger"
)


type CsvLine struct {
    Column1 string
    Column2 string
}

var ip IP

func main() {
	ranger := cidranger.NewPCTrieRanger()

    lines, err := ReadCsv("ip_ranges_asn_only.csv")
    if err != nil {
        panic(err)
    }

    // Loop through lines & turn into object
    for _, line := range lines {
        data := CsvLine{
            Column1: line[0],
            Column2: line[1],
        }
		fmt.Println(data.Column1 + " " + data.Column2)
		_, network, _ := net.ParseCIDR(data.Column2)
        ranger.Insert(cidranger.NewBasicRangerEntry(*network,data.Column1))
    }
    ip_id := "107.178.255.0"
	containingNetworks, err := ranger.ContainingNetworks(net.ParseIP(ip_id))
    fmt.Printf("%+v\n",containingNetworks)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, network := range containingNetworks {
        connected := network.Network()
        netjson, _ := json.Marshal(network)
        fmt.Println(string(netjson))
        fmt.Printf("%+v\n",network)
        // fmt.Printf(network.ipNet)
        fmt.Printf("Containing networks: %s\n", connected.String())
        fmt.Println(network.Getcompany())

        ip = IP {
            IP_Address: ip_id,
            Company:    network.Getcompany(),
        }
        fmt.Printf("%+v\n",ip)
	}
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

    // Open CSV file
    f, err := os.Open(filename)
    if err != nil {
        return [][]string{}, err
    }
    defer f.Close()

    // Read File into a Variable
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        return [][]string{}, err
    }

    return lines, nil
}