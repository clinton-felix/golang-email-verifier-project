package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}

	// if error is not nil or is equal to an error from the scan
	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasDMARC, hasSPF bool
	var spfRecord, dmarcRecord string

	// looking up the mx of the domain we are parsing as argument into checkDomain
	mxRecord, err := net.LookupMX(domain)

	if err != nil {
		log.Print("Error: %v\n", err)
	}
	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecords, err :=  net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}