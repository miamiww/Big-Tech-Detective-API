package IPs

import (
	"github.com/gocql/gocql"
)

// IP struct to hold profile data for our ip
type IP struct {
	ID        gocql.UUID `json:"id"`
	IPV4      string `json:"ipv4"`
	Company   string `json:"company"`
}

// GetIPResponse to form payload returning a single IP struct
type GetIPResponse struct {
	IP IP `json:"ip"`
}

// AllIPsResponse to form payload of an array of IP structs
type AllIPsResponse struct {
	IPs []IP `json:"ips"`
}

// NewIPResponse builds a payload of new ip resource ID
type NewIPResponse struct {
	ID gocql.UUID `json:"id"`
}

// ErrorResponse returns an array of error strings if appropriate
type ErrorResponse struct {
	Errors []string `json:"errors"`
}
