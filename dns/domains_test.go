package dns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cloudless-no/domeneshop-dns-go/dns/schema"
)

func TestDomainList(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	env.Mux.HandleFunc(pathDomains, func(w http.ResponseWriter, r *http.Request) {
		res := schema.DomainListResponse{
			{ID: "1", Domain: "domeneshop.com"},
			{ID: "2", Domain: "domeneshop.cloud"},
		}

		if r.URL.Query().Get("domain") != "" {
			res = schema.DomainListResponse{res[0]}
		}

		json.NewEncoder(w).Encode(res) // nolint: errcheck
	})

	domains, _, err := env.Client.Domain.List(env.Context)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(domains) != 2 {
		t.Errorf("expected %d but got %d", 2, len(domains))
	}
}

func TestDomainGetByID(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	as := newAssert(t)

	env.Mux.HandleFunc(fmt.Sprintf("%s/1", pathDomains), func(w http.ResponseWriter, r *http.Request) {
		resp := schema.Domain{
			ID: "1",
		}

		json.NewEncoder(w).Encode(schema.DomainResponse{ // nolint: errcheck
			Domain: resp,
		})
	})

	id := "0"
	_, resp, err := env.Client.Domain.GetByID(env.Context, id)
	if as.EqInt(http.StatusNotFound, resp.StatusCode) {
		as.Error(err)
	}

	id = "1"
	domain, _, err := env.Client.Domain.GetByID(env.Context, id)
	if as.NoError(err) && as.NotNil(domain) {
		as.EqStr(id, domain.ID)
	}
}
