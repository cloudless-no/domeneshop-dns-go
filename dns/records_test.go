package dns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cloudless-no/domeneshop-dns-go/dns/schema"
)

func TestRecordList(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	as := newAssert(t)

	env.Mux.HandleFunc(pathRecords, func(w http.ResponseWriter, r *http.Request) {
		resp := schema.RecordListResponse{
			Records: []schema.Record{
				{
					ID:     "1",
					Type:   "A",
					Host:   "domeneshop.cloud",
					DomainID: "1",
				},
				{
					ID:     "2",
					Type:   "A",
					Host:   "domeneshop.com",
					DomainID: "2",
				},
				{
					ID:     "3",
					Type:   "A",
					Host:   "dns.domeneshop.com",
					DomainID: "2",
				},
			},
		}

		switch r.URL.Query().Get("domain_id") {
		case "1":
			resp.Records = []schema.Record{resp.Records[0]}
		case "3":
			resp.Records = []schema.Record{}
		}

		json.NewEncoder(w).Encode(resp) // nolint: errcheck
	})

	opts := RecordListOpts{}
	domains, _, err := env.Client.Record.List(env.Context, opts)
	as.NoError(err)
	as.EqInt(3, len(domains))

	opts.DomainID = "1"
	domains, _, err = env.Client.Record.List(env.Context, opts)
	as.NoError(err)
	as.EqInt(1, len(domains))

	opts.DomainID = "3"
	domains, _, err = env.Client.Record.List(env.Context, opts)
	as.NoError(err)
	as.EqInt(0, len(domains))
}

func TestRecordGetByID(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	as := newAssert(t)

	env.Mux.HandleFunc(fmt.Sprintf("%s/1", pathRecords), func(w http.ResponseWriter, r *http.Request) {
		var resp schema.RecordResponse
		resp.Record = schema.Record{
			ID:   "1",
			Host: "domeneshop.com",
		}

		json.NewEncoder(w).Encode(resp) // nolint: errcheck
	})

	id := "0"
	_, resp, err := env.Client.Record.GetByID(env.Context, id, RecordListOpts{DomainID: "1"})
	if !as.EqInt(http.StatusNotFound, resp.StatusCode) {
		as.NoError(err)
	}

	id = "1"
	rec, _, err := env.Client.Record.GetByID(env.Context, id, RecordListOpts{DomainID: "1"})
	if as.NoError(err) {
		as.EqStr(id, rec.ID)
		as.EqStr("domeneshop.com", rec.Host)
	}
}

func TestRecordCreate(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	as := newAssert(t)

	env.Mux.HandleFunc(pathRecords, func(w http.ResponseWriter, r *http.Request) {
		var body schema.RecordCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		var resp schema.RecordResponse
		resp.Record = schema.Record{
			ID:   "1",
			Host: body.Host,
		}

		json.NewEncoder(w).Encode(resp) // nolint: errcheck
	})

	opts := RecordCreateOpts{
		Host:  "",
		Type:  RecordTypeA,
		Data: "10.0.0.0",
		Domain:  &Domain{ID: "1"},
	}

	_, _, err := env.Client.Record.Create(env.Context, opts)
	if as.Error(err) {
		as.EqStr("host required", err.Error())
	}

	opts.Host = "domeneshop.com"
	rec, _, err := env.Client.Record.Create(env.Context, opts)
	if as.NoError(err) {
		as.EqStr("1", rec.ID)
		as.EqStr(opts.Host, rec.Host)
	}
}

func TestRecordUpdate(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	as := newAssert(t)

	env.Mux.HandleFunc(fmt.Sprintf("%s/1", pathRecords), func(w http.ResponseWriter, r *http.Request) {
		var body schema.RecordCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		var resp schema.RecordResponse
		resp.Record = schema.Record{
			ID:   "1",
			Host: body.Host,
		}

		json.NewEncoder(w).Encode(resp) // nolint: errcheck
	})

	rc := &Record{ID: "1"}
	opts := RecordUpdateOpts{
		Host:  "",
		Type:  RecordTypeA,
		Data: "10.0.0.0",
		Domain:  &Domain{ID: "1"},
	}

	_, _, err := env.Client.Record.Update(env.Context, rc, opts)
	if as.Error(err) {
		as.EqStr("host required", err.Error())
	}

	rc = &Record{ID: "0"}
	opts.Host = "dns.domeneshop.com"
	_, resp, err := env.Client.Record.Update(env.Context, rc, opts)
	if !as.EqInt(http.StatusNotFound, resp.StatusCode) {
		as.NoError(err)
	}

	rc = &Record{ID: "1"}
	opts.Host = "dns.domeneshop.com"
	recNew, _, err := env.Client.Record.Update(env.Context, rc, opts)
	if as.NoError(err) {
		as.EqStr(rc.ID, recNew.ID)
		as.EqStr(opts.Host, recNew.Host)
	}
}

func TestRecordDelete(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	as := newAssert(t)

	env.Mux.HandleFunc(fmt.Sprintf("%s/1", pathRecords), func(w http.ResponseWriter, r *http.Request) {})

	rec := &Record{ID: "0"}
	resp, err := env.Client.Record.Delete(env.Context, rec, RecordUpdateOpts{Domain: &Domain{ID: "1"}})
	if !as.EqInt(http.StatusNotFound, resp.StatusCode) {
		as.NoError(err)
	}

	rec.ID = "1"
	_, err = env.Client.Record.Delete(env.Context, rec, RecordUpdateOpts{Domain: &Domain{ID: "1"}})
	as.NoError(err)
}
