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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete.asciidoc#L172>
//
// --------------------------------------------------------------------------------
// DELETE /twitter/_doc/1
// --------------------------------------------------------------------------------

func Test_docs_delete_c5e5873783246c7b1c01d8464fed72c4(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:c5e5873783246c7b1c01d8464fed72c4[]
	res, err := es.Delete("twitter", "1", es.Delete.WithPretty())
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:c5e5873783246c7b1c01d8464fed72c4[]
}
