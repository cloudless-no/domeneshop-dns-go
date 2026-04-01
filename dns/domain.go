package dns

import (
	"context"
	"fmt"

	"github.com/cloudless-no/domeneshop-dns-go/dns/schema"
)

// type DomainStatus string

// const (
// 	DomainStatusActive 		DomainStatus = `active`
// 	DomainStatusInactive   	DomainStatus = `inactive`
// 	DomainStatusPending  	DomainStatus = `pending`
// )

// Domain represents a domain in Domeneshop DNS.
type Domain struct {
	ID              string          
	Registered      schema.DnsTime  
	ExpiryDate      schema.DnsTime  
	Domain          string     
	NS              []string      
	Registrant      string       
	Renew           bool          
	Status          string     
	Services 		*Services
}

// TxtVerification represents the text verification of a domain.
type Services struct {
	Registrar 	string 
	DNS 		bool
	Email 		bool
	WebHotel 	bool
}

// DomainClient is a client for domains API.
type DomainClient struct {
	client *Client
}



// List returns all domains with the given parameters.
func (c DomainClient) List(ctx context.Context) ([]*Domain, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("%s?%s", pathDomains), nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.DomainListResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	domains := make([]*Domain, 0, len(body.Domains))
	for _, z := range body.Domains {
		domains = append(domains, DomainFromSchema(z))
	}

	return domains, resp, nil
}

// GetByID returns the domain with the given id.
func (c DomainClient) GetByID(ctx context.Context, id string) (*Domain, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s", pathDomains, id), nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.DomainResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return DomainFromSchema(body.Domain), resp, nil
}
