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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/count.asciidoc#L7>
//
// --------------------------------------------------------------------------------
// GET /twitter/_count?q=user:kimchy
// --------------------------------------------------------------------------------

func Test_search_count_1b542e3ea87a742f95641d64dcfb1bdb(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1b542e3ea87a742f95641d64dcfb1bdb[]
	res, err := es.Count(
		es.Count.WithIndex("?q=user:kimchy"),
		es.Count.WithQuery("user:kimchy"),
		es.Count.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:1b542e3ea87a742f95641d64dcfb1bdb[]
}
