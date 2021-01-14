package main
import (
  "net/http"
  "os"
  "log"
  "fmt"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "github.com/miamiww/Blocker-API/IPs"
  "github.com/miamiww/Blocker-API/Data"

)

type heartbeatResponse struct {
  Status string `json:"status"`
  Code int `json:"code"`
}

type messageResponse struct{
  Update bool `json:"update"`
  Message string `json:"message"`
}

func main() {

  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }

  CIDRs := Data.BlockRanger
  fmt.Println(CIDRs)

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", heartbeat)
  // router.HandleFunc("/ips/new/", IPs.Post)
  router.HandleFunc("/ip/", IPs.Post)
  router.HandleFunc("/ips/", IPs.Get)
  router.HandleFunc("/ips/{ipv4}",IPs.GetOne)
  router.HandleFunc("/message/",message)
  fmt.Println("server started")
  log.Fatal(http.ListenAndServe(addr,handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}

func heartbeat(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}

func message(w http.ResponseWriter, r *htp.Request){
  jsoon.NewEncoder(w).Encode(messageResponse{Update:false, Message:""})
}

// for Heroku
func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}