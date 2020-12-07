package Data
import (
	
  "net"
  "github.com/miamiww/cidranger"
  "github.com/miamiww/Blocker-API/Postgres"  
  "fmt"
  "context"
)

var BlockRanger cidranger.Ranger //making Ranger accessible outside of the package
func init() {
  BlockRanger = cidranger.NewPCTrieRanger()

  // m := map[string]interface{}{}

  // query := "SELECT Company,CIDR FROM test;"
  iterable, err := Postgres.Conn.Query(context.Background(),"SELECT Company, CIDR FROM test;")
  if err != nil {
		fmt.Println(err)
	}
  for iterable.Next() {
    var company string
    var cidr string
    if err := iterable.Scan(&company); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			fmt.Println(err)
    }
    if err := iterable.Scan(&cidr); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			fmt.Println(err)
		}
	  _, network, _ := net.ParseCIDR(cidr)
	  BlockRanger.Insert(cidranger.NewBasicRangerEntry(*network,company))
  }
  fmt.Println("data trie init done")
}