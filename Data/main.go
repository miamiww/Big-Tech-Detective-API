package Data
import (
	
  "net"
  "github.com/miamiww/cidranger"
  "github.com/miamiww/cassandraAPI/Cassandra"  
  "fmt"
)

var BlockRanger cidranger.Ranger //making Ranger accessible outside of the package
func init() {
  BlockRanger = cidranger.NewPCTrieRanger()

  m := map[string]interface{}{}

  query := "SELECT Company,CIDR FROM ipdatabase.ipblocks"
  iterable := Cassandra.Session.Query(query).Iter()
  for iterable.MapScan(m) {

	  _, network, _ := net.ParseCIDR(m["cidr"].(string))
	  BlockRanger.Insert(cidranger.NewBasicRangerEntry(*network,m["company"].(string)))
	  m = map[string]interface{}{}
  }
  fmt.Println("data trie init done")
}