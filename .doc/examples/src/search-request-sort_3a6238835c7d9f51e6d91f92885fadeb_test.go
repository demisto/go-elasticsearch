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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L11>
//
// --------------------------------------------------------------------------------
// PUT /my-index-000001
// {
//   "mappings": {
//     "properties": {
//       "post_date": { "type": "date" },
//       "user": {
//         "type": "keyword"
//       },
//       "name": {
//         "type": "keyword"
//       },
//       "age": { "type": "integer" }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_3a6238835c7d9f51e6d91f92885fadeb(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:3a6238835c7d9f51e6d91f92885fadeb[]
	res, err := es.Indices.Create(
		"my-index-000001",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "post_date": {
		        "type": "date"
		      },
		      "user": {
		        "type": "keyword"
		      },
		      "name": {
		        "type": "keyword"
		      },
		      "age": {
		        "type": "integer"
		      }
		    }
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:3a6238835c7d9f51e6d91f92885fadeb[]
}
