// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/demisto/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/api-conventions.asciidoc#L231>
//
// --------------------------------------------------------------------------------
// GET /_search?q=elasticsearch&filter_path=took,hits.hits._id,hits.hits._score
// --------------------------------------------------------------------------------

func Test_api_conventions_09dbd90c5e22ea4a17b4cf9aa72e08ae(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:09dbd90c5e22ea4a17b4cf9aa72e08ae[]
	res, err := es.Search(
		es.Search.WithFilterPath("took,hits.hits._id,hits.hits._score"),
		es.Search.WithQuery("elasticsearch"),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:09dbd90c5e22ea4a17b4cf9aa72e08ae[]
}
