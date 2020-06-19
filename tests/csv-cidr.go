package main

import (
    "encoding/csv"
    "encoding/json"
	"fmt"
	"net"
    "os"
	"github.com/yl2chen/cidranger"
)


type CsvLine struct {
    Column1 string
    Column2 string
}

type companyRangerEntry interface {
    Network() net.IPNet
    Getcompany() string
}

type RangerEntry interface {
    Network() net.IPNet
    Getcompany() string
}


type basicCompanyRangerEntry struct {
    IpNet net.IPNet
    Company string
}

func (b *basicCompanyRangerEntry) Network() net.IPNet{
    return b.IpNet 
}

func (b *basicCompanyRangerEntry) Getcompany() string {
    return b.Company
}

// NewBasicRangerEntry returns a basic RangerEntry that only stores the network
// itself.
func NewCompanyRangerEntry(ipNet net.IPNet, company string) companyRangerEntry {
	return &basicCompanyRangerEntry{
        IpNet: ipNet,
        Company: company,
	}
}

type companyRanger interface {
	Insert(entry companyRangerEntry) error
	Remove(network net.IPNet) (companyRangerEntry, error)
	Contains(ip net.IP) (bool, error)
	ContainingNetworks(ip net.IP) ([]companyRangerEntry, error)
	CoveredNetworks(network net.IPNet) ([]companyRangerEntry, error)
	Len() int
}

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
		ranger.Insert(NewCompanyRangerEntry(*network,data.Column1))
	}
	containingNetworks, err := ranger.ContainingNetworks(net.ParseIP("107.178.255.0"))
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