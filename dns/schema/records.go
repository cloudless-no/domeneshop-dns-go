package schema

// Record represents a record in Domeneshop DNS.
type Record struct {
	Type     string   `json:"type"`
	ID       string   `json:"id"`
	Host     string   `json:"host"`
	Data     string   `json:"data"`
	Ttl      int      `json:"ttl"`
	DomainID   string  `json:"domain_id"`
}

// RecordListResponse defines the schema of the response when
// listing zones.
type RecordListResponse struct {
	Records []Record `json:"records"`
}

// RecordResponse defines the schema of the response when
// listing zones.
type RecordResponse struct {
	Record Record `json:"record"`
}

// RecordCreateRequest defines a schema for the request to
// create a record.
type RecordCreateRequest struct {
	Host   string `json:"host"`
	Ttl    *int   `json:"ttl"`
	Type   string `json:"type"`
	Data  string `json:"data"`
	DomainID   string  `json:"domain_id"`
}

// RecordUpdateRequest defines a schema for the request to
// update a record.
type RecordUpdateRequest struct {
	Host   string `json:"host"`
	Ttl    *int   `json:"ttl"`
	Type   string `json:"type"`
	Data  string `json:"data"`
	DomainID   string  `json:"domain_id"`
}
