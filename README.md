# API for Big Tech Detective

### A RESTful API for querying a Postgres database of CIDR blocks

API endpoints and server functions are defined in main.go
Data from a Postgres database is pulled into a trie data structure in /Data
The functions that respond to API requests and process and define the data returned are stored in IPs/ 