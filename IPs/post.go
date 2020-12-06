package IPs

// import (
// 	"net/http"
// 	"github.com/gocql/gocql"
// 	"encoding/json"
// 	"github.com/miamiww/Blocker-API/Cassandra"
// 	"fmt"
// )

// // Post -- handles POST request to /ips/new to create new ip
// // params:
// // w - response writer for building JSON payload response
// // r - request reader to fetch form data or url params
// func Post(w http.ResponseWriter, r *http.Request) {
// 	var errs []string
// 	var gocqlUUID gocql.UUID

// 	ip, errs := FormToIP(r)

// 	var created bool = false
// 	if len(errs) == 0 {
// 		fmt.Println("creating a new ip")
// 		gocqlUUID = gocql.TimeUUID()
// 		if err := Cassandra.Session.Query(`
// 		INSERT INTO ips (id, ipv4, company) VALUES (?, ?, ?)`,
// 		gocqlUUID, ip.IPV4, ip.Company).Exec(); err != nil {
// 			errs = append(errs, err.Error())
// 		} else {
// 			created = true
// 		}
// 	}

// 	if created {
// 		fmt.Println("ip_id", gocqlUUID)
// 		json.NewEncoder(w).Encode(NewIPResponse{ID: gocqlUUID})
// 	} else {
// 		fmt.Println("errors", errs)
// 		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
// 	}
// }
