// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build integration

package estransport_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/demisto/go-elasticsearch/v8/estransport"
	"github.com/demisto/go-elasticsearch/v8/esutil"
)

var (
	_ = fmt.Print
)

func TestTransportRetries(t *testing.T) {
	var counter int

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		counter++

		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println("req.Body:", string(body))

		http.Error(w, "FAKE 502", http.StatusBadGateway)
	}))
	serverURL, _ := url.Parse(server.URL)

	transport, _ := estransport.New(estransport.Config{URLs: []*url.URL{serverURL}})

	bodies := []io.Reader{
		strings.NewReader(`FAKE`),
		esutil.NewJSONReader(`FAKE`),
	}

	for _, body := range bodies {
		t.Run(fmt.Sprintf("Reset the %T request body", body), func(t *testing.T) {
			counter = 0

			req, err := http.NewRequest("GET", "/", body)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			res, err := transport.Perform(req)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			body, _ := ioutil.ReadAll(res.Body)

			fmt.Println("> GET", req.URL)
			fmt.Printf("< %s (tries: %d)\n", bytes.TrimSpace(body), counter)

			if counter != 4 {
				t.Errorf("Unexpected number of attempts, want=4, got=%d", counter)
			}
		})
	}
}

func TestTransportHeaders(t *testing.T) {
	u, _ := url.Parse("http://localhost:9200")

	hdr := http.Header{}
	hdr.Set("Accept", "application/yaml")

	tp, _ := estransport.New(estransport.Config{
		URLs:   []*url.URL{u},
		Header: hdr,
	})

	req, _ := http.NewRequest("GET", "/", nil)
	res, err := tp.Perform(req)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if !bytes.HasPrefix(body, []byte("---")) {
		t.Errorf("Unexpected response body:\n%s", body)
	}
}
