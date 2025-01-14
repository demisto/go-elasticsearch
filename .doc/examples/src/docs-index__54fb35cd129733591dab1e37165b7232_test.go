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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/index_.asciidoc#L443>
//
// --------------------------------------------------------------------------------
// PUT twitter/_doc/1?version=2&version_type=external
// {
//   "message" : "elasticsearch now has versioning support, double cool!"
// }
// --------------------------------------------------------------------------------

func Test_docs_index__54fb35cd129733591dab1e37165b7232(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:54fb35cd129733591dab1e37165b7232[]
	res, err := es.Index(
		"twitter",
		strings.NewReader(`{
		  "message": "elasticsearch now has versioning support, double cool!"
		}`),
		es.Index.WithDocumentID("1"),
		es.Index.WithVersion(2),
		es.Index.WithVersionType("external"),
		es.Index.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:54fb35cd129733591dab1e37165b7232[]
}
