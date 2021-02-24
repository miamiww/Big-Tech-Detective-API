# API for [Big Tech Detective](https://bigtechdetective.net)

## A RESTful API for querying a Postgres database of CIDR blocks

This is the API for the [Big Tech Detective](https://gitlab.com/big-tech-detective/Big-Tech-Detective) browser extension. CIDR blocks are stored in a Postgres database and pulled into a trie data structure for easy lookup (as described [here](https://github.com/yl2chen/cidranger)). The API is accessible at [https://bigtechdetective.club](https://bigtechdetective.club)

## API Documentation

The API has the following endpoints:

### [https://bigtechdetective.club/](https://bigtechdetective.club/) returns a heartbeat

### [/ips](https://bigtechdetective.club/ips) returns the whole dataset in json

```json
{
    "ips":
        [
            {"CIDR":"54.230.212.0/23", "company":"Amazon"},
            {"CIDR":"54.230.214.0/23", "company":"Amazon"},
            // 'etc'
            {"CIDR":"2620:112:3000::/44", "company":"Microsoft"}

        ]
}
```

### [https://bigtechdetective.club/ips/{ip_address}](https://bigtechdetective.club/ips/{ip_address}) gives returns a response indicating whether or not ip_address is in the dataset

for example, [/ips/54.192.0.0](https://bigtechdetective.club/ips/54.192.0.0) returns

```json
{
    "ip":
        {
            "ipaddress":"54.192.0.0",
            "company":"Amazon"    
        }
}
```

an address that isn't in the database, such as `/ips/192.168.0.1`, returns

```json
{
    "errors": [
        "IP not found"
    ]
}
```

and a request that isn't recognized as an ip address returns, such as `/ips/ipaddress`

```json
{
    "errors": [
        "not a valid IP address",
        "IP not found"
    ]
}

```

## Can I build my own project from the API?

Please do! But keep in mind our server costs 😅 Let us know what you make by emailing us at detective@bigtechdetective.net
