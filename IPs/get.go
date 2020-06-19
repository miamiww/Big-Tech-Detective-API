package IPs

import (
	"os"
	"net"
	"net/http"
	// "github.com/gocql/gocql"
	"encoding/json"
	"encoding/csv"
	// "github.com/miamiww/cassandraAPI/Cassandra"
	"github.com/gorilla/mux"
	"github.com/miamiww/cidranger"
)



// Get -- handles GET request to /ips/ to fetch all ips
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params (unused here)
// func Get(w http.ResponseWriter, r *http.Request) {
// 	var ipList []IP
// 	m := map[string]interface{}{}

// 	query := "SELECT id,ipv4,company FROM ips"
// 	iterable := Cassandra.Session.Query(query).Iter()
// 	for iterable.MapScan(m) {
// 		ipList = append(ipList, IP{
// 			ID:        m["id"].(gocql.UUID),
// 			IPV4:      m["ipv4"].(string),
// 			Company:   m["company"].(string),
// 		})
// 		m = map[string]interface{}{}
// 	}

// 	json.NewEncoder(w).Encode(AllIPsResponse{IPs: ipList})
// }

// GetOne -- handles GET request to /ips/{ipv4} to fetch one ip
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params
func GetOne(w http.ResponseWriter, r *http.Request) {
	var ip IP
	var errs []string
	var found bool = false

	vars := mux.Vars(r)
	ip_id := vars["ipv4"]
	ip_address_checked := net.ParseIP(ip_id)
	if ip_address_checked == nil{
		errs = append(errs, "not a valid IP address")
	} else{
		type CsvLine struct {
			Column1 string
			Column2 string
		}
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
			_, network, _ := net.ParseCIDR(data.Column2)
			ranger.Insert(cidranger.NewBasicRangerEntry(*network,data.Column1))
		}
		found, err = ranger.Contains(ip_address_checked)
		if found {
			containingNetworks, err := ranger.ContainingNetworks(ip_address_checked)

		
			if err != nil {
				errs = append(errs, "Trie failure")
				return
			}
	
			for _, network := range containingNetworks {
				ip = IP {
					IP_Address: ip_id,
					Company:    network.Getcompany(),
				}
			}
		}

	}

	if !found {
		errs = append(errs, "IP not found")
	}


    if found {
		json.NewEncoder(w).Encode(GetIPResponse{IP: ip})
	} else {
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}

// Enrich -- turns an array of ip UUIDs into a map of {uuid: "firstname lastname"}
// params:
// uuids - array of ip UUIDs to fetch
// returns:
// a map[string]string of {uuid: "firstname lastname"}
// func Enrich(uuids []gocql.UUID) map[string]string {
// 	if len(uuids) > 0 {
// 		fmt.Println("---\nfetching names", uuids)
// 		names := map[string]string{}
// 		m := map[string]interface{}{}

// 		query := "SELECT id,firstname,lastname FROM ips WHERE id IN ?"
// 		iterable := Cassandra.Session.Query(query, uuids).Iter()
// 		for iterable.MapScan(m) {
// 			fmt.Println("m", m)
// 			ipID := m["id"].(gocql.UUID)
// 			fmt.Println("ipID", ipID.String())
// 			names[ipID.String()] = fmt.Sprintf("%s %s", m["firstname"].(string), m["lastname"].(string))
// 			m = map[string]interface{}{}
// 		}
// 		fmt.Println("names", names)
// 		return names
// 	}
// 	return map[string]string{}
// }


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