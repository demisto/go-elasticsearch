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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/match-query.asciidoc#L18>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "match": {
//       "message": {
//         "query": "this is a test"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_match_query_4c3db8987d7b2d3d3df78ff1e71e7ede(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:4c3db8987d7b2d3d3df78ff1e71e7ede[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match": {
		      "message": {
		        "query": "this is a test"
		      }
		    }
		  }
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:4c3db8987d7b2d3d3df78ff1e71e7ede[]
}
