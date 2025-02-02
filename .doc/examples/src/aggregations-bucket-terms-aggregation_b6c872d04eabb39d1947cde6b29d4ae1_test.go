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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/terms-aggregation.asciidoc#L369>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "aggs": {
//     "tags": {
//       "terms": {
//         "field": "tags",
//         "min_doc_count": 10
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_terms_aggregation_b6c872d04eabb39d1947cde6b29d4ae1(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b6c872d04eabb39d1947cde6b29d4ae1[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "tags": {
		      "terms": {
		        "field": "tags",
		        "min_doc_count": 10
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
	// end:b6c872d04eabb39d1947cde6b29d4ae1[]
}
