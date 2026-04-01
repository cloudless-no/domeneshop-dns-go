package schema

// DomainListResponse defines the schema of the response when
// listing domains.
type DomainListResponse struct {
	Domains []Domain `json:"domains"`
}

// DomainGetResponse defines the schema of the response when
// listing domains.
type DomainResponse struct {
	Domain Domain `json:"domain"`
}

// Domain represents a domain in Domeneshop DNS.
type Domain struct {
	ID              string          `json:"id"`
	Registered      DnsTime        	`json:"registered_date"`
	ExpiryDate      DnsTime        	`json:"expiry_date"`
	Domain          string          `json:"domain"`
	NS              []string        `json:"nameservers"`
	Registrant      string          `json:"registrant"`
	Renew           bool            `json:"renew"`
	Status          string          `json:"status"`
	Services 		Services 		`json:"services"`
}

// TxtVerification represents the text verification of a domain.
type Services struct {
	Registrar 	string 	`json:"registrar"`
	DNS 		bool 	`json:"dns"`
	Email 		bool 	`json:"email"`
	WebHotel 	bool 	`json:"web_hotel"`
}
