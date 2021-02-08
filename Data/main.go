package Data
import (
	
  "net"
  "github.com/miamiww/cidranger"
  "fmt"
	"github.com/jackc/pgx/v4"
	"os"
  "context"
)

var BlockRanger cidranger.Ranger //making Ranger accessible outside of the package
func init() {

  // setting up the connection to the postgres database, will work locally and remote because DATABASE_URL is set up on both
  var err error
  conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

  if err != nil {
    panic("failed to connect database")
  }
  defer conn.Close(context.Background())

  fmt.Println("postgres connect init done")
  

  //Setting up the trie build
  BlockRanger = cidranger.NewPCTrieRanger()

  // this is taking the whole table of cidr blocks from the database 
  // iterable is all the rows that need to be iterated through
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
    // parse cidr block notation into useable form by the trie ranger
	  _, network, _ := net.ParseCIDR(cidr)
	  BlockRanger.Insert(cidranger.NewBasicRangerEntry(*network,company))
  }
  fmt.Println("data trie init done")
}