package Data
import (
	
  "net"
  "github.com/miamiww/cidranger"
  // "github.com/miamiww/Blocker-API/Postgres"  
  "fmt"
	"github.com/jackc/pgx/v4"
	"os"
  "context"
)

var BlockRanger cidranger.Ranger //making Ranger accessible outside of the package
func init() {

  var err error
  conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

//   CIDR_db, err = gorm.Open( "postgres", "host=db port=5432 user=alden dbname=cidr_blocks sslmode=disable password=")

  if err != nil {
    panic("failed to connect database")
  }
  defer conn.Close(context.Background())

  fmt.Println("postgres connect init done")
  
  BlockRanger = cidranger.NewPCTrieRanger()

  // m := map[string]interface{}{}
  // var (
  //   company string
  //   cidr string
  // )

  // query := "SELECT Company,CIDR FROM test;"
  iterable, err := conn.Query(context.Background(),"SELECT Company, CIDR FROM test;")
  // fmt.Println(iterable)

  if err != nil {
		fmt.Println(err)
	}
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
    // if err := iterable.Scan(&cidr); err != nil {
		// 	// Check for a scan error.
		// 	// Query rows will be closed with defer.
		// 	fmt.Println(err)
		// }
	  _, network, _ := net.ParseCIDR(cidr)
	  BlockRanger.Insert(cidranger.NewBasicRangerEntry(*network,company))
  }
  fmt.Println("data trie init done")
}