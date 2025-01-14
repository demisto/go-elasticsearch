// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/demisto/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/index_.asciidoc#L547>
//
// --------------------------------------------------------------------------------
// PUT twitter/_doc/1?op_type=create
// {
//   "user" : "kimchy",
//   "post_date" : "2009-11-15T14:12:12",
//   "message" : "trying out Elasticsearch"
// }
// --------------------------------------------------------------------------------

func Test_docs_index__7e2e9f23daa324318b834bfaf67b798c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:7e2e9f23daa324318b834bfaf67b798c[]
	res, err := es.Index(
		"twitter",
		strings.NewReader(`{
		  "user": "kimchy",
		  "post_date": "2009-11-15T14:12:12",
		  "message": "trying out Elasticsearch"
		}`),
		es.Index.WithDocumentID("1"),
		es.Index.WithOpType("create"),
		es.Index.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:7e2e9f23daa324318b834bfaf67b798c[]
}
