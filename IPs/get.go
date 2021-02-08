package IPs

import (
	"net"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/miamiww/Blocker-API/Data"
	"github.com/jackc/pgx/v4"
	"context"
	"os"
	"io/ioutil"
)



// Get -- handles GET request to /ips/ to fetch all ips
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params (unused here)
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all")

	var ipList []CIDRS

	var err error
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
  
	if err != nil {
		fmt.Println("failed to connect database")
	}
	defer conn.Close(context.Background())

	iterable, err := conn.Query(context.Background(),"SELECT Company, CIDR FROM cidrs;")

	if err != nil {
		  fmt.Println(err)
	}
	
	// loops through all rows using Next() method, which closes out connection automatically
	for iterable.Next() {
	  var (
		  company string
		  cidr string
		)
	  if err := iterable.Scan(&company, &cidr); err != nil {
			  // Check for a scan error.
			  // Query rows will be closed with defer.
			  fmt.Println(err)
	  }
	  ipList = append(ipList, CIDRS{
				CIDR:      cidr,
				Company:   company,
			})
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
	var err error
	
	vars := mux.Vars(r)
	ip_id := vars["ipv4"]
	ip_address_checked := net.ParseIP(ip_id)
	if ip_address_checked == nil{
		errs = append(errs, "not a valid IP address")
	} else{
		// fetches the trie ranger of cidr blocks from Data module and checks to see if the requested IP is within it
		ranger := Data.BlockRanger

		found, err = ranger.Contains(ip_address_checked)
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

func Post(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		fmt.Println("getting post")
	
		var input PostIPRecieve
		var ip IP
		var errs []string
		var found bool = false
		var err error

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errs = append(errs, "Error reading request body")

			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)

		}

		json.Unmarshal([]byte(body),&input)
		ip_id := input.IP
		ip_address_checked := net.ParseIP(ip_id)
		if ip_address_checked == nil{
			errs = append(errs, "not a valid IP address")
		} else{
			// fetches the trie ranger of cidr blocks from Data module and checks to see if the requested IP is within it
			ranger := Data.BlockRanger

			found, err = ranger.Contains(ip_address_checked)
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

	} else {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}

	
	
}