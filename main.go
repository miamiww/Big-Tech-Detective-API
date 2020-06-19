package main
import (
  "net/http"
  "log"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  // "github.com/miamiww/cassandraAPI/Cassandra"
  "github.com/miamiww/cassandraAPI/IPs"
)

type heartbeatResponse struct {
  Status string `json:"status"`
  Code int `json:"code"`
}

func main() {
  // CassandraSession := Cassandra.Session
  // defer CassandraSession.Close()

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", heartbeat)
  // router.HandleFunc("/ips/new/", IPs.Post)
  // router.HandleFunc("/ips/", IPs.Get)
  router.HandleFunc("/ips/{ipv4}",IPs.GetOne)
  log.Fatal(http.ListenAndServeTLS(":8080", "/etc/letsencrypt/live/thegreatest.website/fullchain.pem","/etc/letsencrypt/live/thegreatest.website/privkey.pem",handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
