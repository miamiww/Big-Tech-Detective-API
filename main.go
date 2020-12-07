package main
import (
  "net/http"
  "log"
  "fmt"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  // "github.com/miamiww/Blocker-API/Postgres"
  "github.com/miamiww/Blocker-API/IPs"
  "github.com/miamiww/Blocker-API/Data"
)

type heartbeatResponse struct {
  Status string `json:"status"`
  Code int `json:"code"`
}

func main() {
  // CIDR_connection := Postgres.Conn

  CIDRanger := Data.BlockRanger
  fmt.Println(CIDRanger)

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", heartbeat)
  // router.HandleFunc("/ips/new/", IPs.Post)
  router.HandleFunc("/ips/", IPs.Get)
  router.HandleFunc("/ips/{ipv4}",IPs.GetOne)
  log.Fatal(http.ListenAndServe(":8080",handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
