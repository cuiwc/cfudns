package cf

import (
	"context"
	"errors"
	"log"

	"github.com/cloudflare/cloudflare-go"
)

var (
	ErrDNSNameInvalid = errors.New("invalid DNS name")
)

type Client struct {
	api *cloudflare.API
}

func New(token string) (*Client, error) {
	api, err := cloudflare.NewWithAPIToken(token)
	if err != nil {
		return nil, err
	}
	return &Client{
		api: api,
	}, nil
}

func (c *Client) UpdateRecord(zoneName, dnsName, dnsContent string) error {
	zoneID, err := c.api.ZoneIDByName(zoneName)
	if err != nil {
		return err
	}

	dnsRecords, err := c.api.DNSRecords(context.Background(), zoneID, cloudflare.DNSRecord{
		Name: dnsName,
	})
	if err != nil {
		return err
	}

	if len(dnsRecords) != 1 {
		return ErrDNSNameInvalid
	}

	dnsRecord := dnsRecords[0]
	log.Println("Current DNS content is: ", dnsRecord.Content)

	dnsRecord.Content = dnsContent
	err = c.api.UpdateDNSRecord(context.Background(), zoneID, dnsRecord.ID, dnsRecord)
	if err != nil {
		return err
	}

	return nil
}
