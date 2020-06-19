package IPs

import (
	// "os"
	"net"
	"net/http"
	// "github.com/gocql/gocql"
	"fmt"
	"encoding/json"
	// "encoding/csv"
	"github.com/miamiww/cassandraAPI/Cassandra"
	"github.com/gorilla/mux"
	"github.com/miamiww/cidranger"
)



// Get -- handles GET request to /ips/ to fetch all ips
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params (unused here)
func Get(w http.ResponseWriter, r *http.Request) {
	var ipList []CIDRS
	m := map[string]interface{}{}

	query := "SELECT Company,CIDR FROM ipblocks"
	iterable := Cassandra.Session.Query(query).Iter()
	for iterable.MapScan(m) {
		ipList = append(ipList, CIDRS{
			CIDR:      m["CIDR"].(string),
			Company:   m["Company"].(string),
		})
		m = map[string]interface{}{}
	}

	json.NewEncoder(w).Encode(AllIPsResponse{CIDRs: ipList})
}

// GetOne -- handles GET request to /ips/{ipv4} to fetch one ip
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params
func GetOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting one")
	var ip IP
	var errs []string
	var found bool = false

	vars := mux.Vars(r)
	ip_id := vars["ipv4"]
	ip_address_checked := net.ParseIP(ip_id)
	if ip_address_checked == nil{
		errs = append(errs, "not a valid IP address")
	} else{
		fmt.Println("making query to database to build trie")

		ranger := cidranger.NewPCTrieRanger()

		m := map[string]interface{}{}

		query := "SELECT Company,CIDR FROM ipdatabase.ipblocks"
		iterable := Cassandra.Session.Query(query).Iter()
		for iterable.MapScan(m) {
			fmt.Println("adding to ranger")

			_, network, _ := net.ParseCIDR(m["CIDR"].(string))
			fmt.Printf("%+v\n",network)
			ranger.Insert(cidranger.NewBasicRangerEntry(*network,m["Company"].(string)))
			m = map[string]interface{}{}
		}

		found, err := ranger.Contains(ip_address_checked)
		if err != nil {
			errs = append(errs, "Trie failure")
			return
		}
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


// func ReadCsv(filename string) ([][]string, error) {

//     // Open CSV file
//     f, err := os.Open(filename)
//     if err != nil {
//         return [][]string{}, err
//     }
//     defer f.Close()

//     // Read File into a Variable
//     lines, err := csv.NewReader(f).ReadAll()
//     if err != nil {
//         return [][]string{}, err
//     }
