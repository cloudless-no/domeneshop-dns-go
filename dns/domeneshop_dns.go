package dns

// Version of the SDK
const Version = "v0.1.0"

// Endpoint is the base URL of the API.
const Endpoint = "https://api.domeneshop.no/v0"
// https://api.domeneshop.no/v0/domains/${domainId}/dns
const UserAgent = "domeneshop-dns/" + Version

const (
	pathDomains        = "/domains"
	pathRecords        = "dns"
	pathFirstDomainRecords        = "/domains/1/dns"
)
