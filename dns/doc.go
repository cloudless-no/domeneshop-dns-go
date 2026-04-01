// Package dns an SDK for the Domeneshop DNS API.
//
// Read the API docs over on https://dns.domeneshop.com/api-docs.
//
// Example:
//
//	package main
//
//	import (
//		"context"
//		"fmt"
//		"log"
//
//		"github.com/cloudless-no/domeneshop-dns-go/dns"
//	)
//
//	func main() {
//		client := dns.NewClient(dns.WithToken("token"))
//
//		record, _, err := client.Record.GetByID(context.Background(), "randomid")
//		if err != nil {
//			log.Fatalf("error retrieving record: %v\n", err)
//		}
//
//		fmt.Printf("record of type: '%s' found with value: %s", record.Type, record.Value)
//	}
package dns
