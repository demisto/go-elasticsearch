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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/multi-match-query.asciidoc#L472>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//    "multi_match" : {
//       "query":      "Jon",
//       "type":       "cross_fields",
//       "analyzer":   "standard", <1>
//       "fields":     [ "first", "last", "*.edge" ]
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_multi_match_query_179f0a3e84ff4bbac18787a018eabf89(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:179f0a3e84ff4bbac18787a018eabf89[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "multi_match": {
		      "query": "Jon",
		      "type": "cross_fields",
		      "analyzer": "standard",
		      "fields": [
		        "first",
		        "last",
		        "*.edge"
		      ]
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
	// end:179f0a3e84ff4bbac18787a018eabf89[]
}
