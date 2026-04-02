package dns

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloudless-no/domeneshop-dns-go/dns/schema"
)

type RecordType string

const (
	RecordTypeA     RecordType = "A"
	RecordTypeAAAA  RecordType = "AAAA"
	RecordTypeMX    RecordType = "MX"
	RecordTypeCNAME RecordType = "CNAME"
	RecordTypeTXT   RecordType = "TXT"
	RecordTypeSRV   RecordType = "SRV"
)

// Record represents a record in the Domeneshop DNS.
type Record struct {
	Type     RecordType
	ID       string
	Host     string
	Data     string
	Ttl      int
	Domain   *Domain
	Priority *int
	Weight   *int
	Port     *int
}

// RecordClient is a client for records API.
type RecordClient struct {
	client *Client
}

// RecordListOpts specifies options for listing records
type RecordListOpts struct {
	DomainID string
}



// List returns all records with the given parameters.
func (c RecordClient) List(ctx context.Context, opts RecordListOpts) ([]*Record, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s", pathDomains, opts.DomainID, pathRecords), nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.RecordListResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	records := make([]*Record, 0, len(body))
	for _, r := range body {
		records = append(records, RecordFromSchema(r))
	}

	return records, resp, nil
}

// GetByID returns a record with the given id.
func (c RecordClient) GetByID(ctx context.Context, id string, opts RecordListOpts) (*Record, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s", pathDomains, opts.DomainID, pathRecords, id), nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.RecordResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return RecordFromSchema(body.Record), resp, nil
}

// RecordCreateOpts specifies options for creating a record.
type RecordCreateOpts struct {
	Host  string
	Ttl   *int
	Type  RecordType
	Data string
	Domain *Domain
}

func (o RecordCreateOpts) validate() error {
	if o.Host == "" {
		return errors.New("host required")
	}
	if o.Type == "" {
		return errors.New("type required")
	}
	if o.Data == "" {
		return errors.New("data required")
	}
	if o.Domain == nil {
		return errors.New("domain_id required")
	}

	return nil
}

// Create creates a new record.
func (c RecordClient) Create(ctx context.Context, opts RecordCreateOpts) (*Record, *Response, error) {
	if err := opts.validate(); err != nil {
		return nil, nil, err
	}

	var reqBody schema.RecordCreateRequest
	reqBody.Host = opts.Host
	reqBody.Ttl = opts.Ttl
	reqBody.Type = string(opts.Type)
	reqBody.Data = opts.Data

	reqBodyData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s", pathDomains, opts.Domain.ID, pathRecords), bytes.NewReader(reqBodyData))
	if err != nil {
		return nil, nil, err
	}

	var body schema.RecordResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return RecordFromSchema(body.Record), resp, nil
}

// RecordUpdateOpts specifies options for creating a record.
type RecordUpdateOpts struct {
	Host  string
	Ttl   *int
	Type  RecordType
	Data string
	Domain  *Domain
}

func (o RecordUpdateOpts) validate() error {
	if o.Host == "" {
		return errors.New("host required")
	}
	if o.Type == "" {
		return errors.New("type required")
	}
	if o.Data == "" {
		return errors.New("data required")
	}
	if o.Domain == nil {
		return errors.New("domain_id required")
	}

	return nil
}

// Update updates a record.
func (c RecordClient) Update(ctx context.Context, rec *Record, opts RecordUpdateOpts) (*Record, *Response, error) {
	if err := opts.validate(); err != nil {
		return nil, nil, err
	}

	var reqBody schema.RecordUpdateRequest
	reqBody.Host = opts.Host
	reqBody.Ttl = opts.Ttl
	reqBody.Type = string(opts.Type)
	reqBody.Data = opts.Data

	reqBodyData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest(ctx, "PUT", fmt.Sprintf("%s/%s/%s/%s", pathDomains, opts.Domain.ID, pathRecords, rec.ID), bytes.NewReader(reqBodyData))
	if err != nil {
		return nil, nil, err
	}

	var body schema.RecordResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return RecordFromSchema(body.Record), resp, nil
}

// Delete deletes a record.
func (c RecordClient) Delete(ctx context.Context, rec *Record, opts RecordUpdateOpts) (*Response, error) {
	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("%s/%s/%s/%s", pathDomains, opts.Domain.ID, pathRecords, rec.ID), nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req, nil)
}

// RecordEntry represents a record entry used for verification and updates.
type RecordEntry struct {
	Type   RecordType
	Host   string
	Data  string
	Ttl    *int
}
