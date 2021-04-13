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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/exists-query.asciidoc#L56>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "bool": {
//       "must_not": {
//         "exists": {
//           "field": "user"
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_exists_query_0c7b637b6b9f18c2149a9f9f51e3f09a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:0c7b637b6b9f18c2149a9f9f51e3f09a[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "bool": {
		      "must_not": {
		        "exists": {
		          "field": "user"
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
	// end:0c7b637b6b9f18c2149a9f9f51e3f09a[]
}
