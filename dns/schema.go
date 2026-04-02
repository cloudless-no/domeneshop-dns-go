package dns

import (
	"strconv"

	"github.com/cloudless-no/domeneshop-dns-go/dns/schema"
)

// DomainFromSchema converts schema.Domain to a Domain.
func DomainFromSchema(s schema.Domain) *Domain {
	domain := &Domain{
		ID:             strconv.Itoa(s.ID),
		Registered:     s.Registered,
		ExpiryDate:     s.ExpiryDate,
		Domain:         s.Domain,
		NS:             s.NS,
		Registrant:     s.Registrant,
		Renew:          s.Renew,
		Status:         s.Status,
		Services: 		ServicesFromSchema(s.Services),
	}

	return domain
}

// TxtVerificationFromSchema converts schema.TxtVerification to TxtVerification
func ServicesFromSchema(s schema.Services) *Services {
	return &Services{
		Registrar:  s.Registrar,
		DNS:        s.DNS,
		Email:      s.Email,
		WebHotel:   s.WebHotel,
	}
}

// RecordFromSchema convers a schema.Record to Record
func RecordFromSchema(s schema.Record) *Record {
	return &Record{
		Type:   RecordType(s.Type),
		ID:     strconv.Itoa(s.ID),
		Host:   s.Host,
		Data:   s.Data,
		Ttl:    s.Ttl,
		Domain: &Domain{ID: strconv.Itoa(s.DomainID)},
	}
}


