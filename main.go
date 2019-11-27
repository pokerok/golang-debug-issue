package main

import (
	"fmt"
	"github.com/zonedb/zonedb"
)

func main() {
	var domain string = "google.com"

	z := zonedb.PublicZone(domain)

	fmt.Println(z)
}
