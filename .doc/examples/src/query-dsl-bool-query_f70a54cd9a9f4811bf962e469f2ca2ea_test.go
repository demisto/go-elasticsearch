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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/bool-query.asciidoc#L88>
//
// --------------------------------------------------------------------------------
// GET _search
// {
//   "query": {
//     "bool": {
//       "filter": {
//         "term": {
//           "status": "active"
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_bool_query_f70a54cd9a9f4811bf962e469f2ca2ea(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:f70a54cd9a9f4811bf962e469f2ca2ea[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "bool": {
		      "filter": {
		        "term": {
		          "status": "active"
		        }
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
	// end:f70a54cd9a9f4811bf962e469f2ca2ea[]
}
