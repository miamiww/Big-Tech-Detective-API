package Postgres

import (
	// "encoding/json"
	// "log"
	// "net/http"
	"fmt"
	// "github.com/gorilla/mux"
	// "github.com/rs/cors"
	"github.com/jackc/pgx/v4"
	"os"
	"context"
)


// var CIDR_db *gorm.DB //making Session accessible outside of the package
var Conn *pgx.Conn
func init() {
  var err error
  Conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

//   CIDR_db, err = gorm.Open( "postgres", "host=db port=5432 user=alden dbname=cidr_blocks sslmode=disable password=")

  if err != nil {

    panic("failed to connect database")
  }
  defer Conn.Close(context.Background())

  fmt.Println("postgres init done")
}