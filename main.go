package main

import (
	"flag"
	"log"
	"os"

	"github.com/cuiwc/cfudns/internal/cf"
)

func main() {
	var tokenFlag = flag.String("token", os.Getenv("CF_TOKEN"), "Cloudflare API token")
	var zoneFlag = flag.String("zone", "", "Zone name")
	var dnsNameFlag = flag.String("dns-name", "", "DNS name")
	var dnsValueFlag = flag.String("dns-value", "", "DNS value")

	flag.Parse()

	if *tokenFlag == "" || *zoneFlag == "" || *dnsNameFlag == "" || *dnsValueFlag == "" {
		flag.Usage()
		os.Exit(-1)
	}

	c, err := cf.New(*tokenFlag)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	err = c.UpdateRecord(*zoneFlag, *dnsNameFlag, *dnsValueFlag)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	os.Exit(0)
}
