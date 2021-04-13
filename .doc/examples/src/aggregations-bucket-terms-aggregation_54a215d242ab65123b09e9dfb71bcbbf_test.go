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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/terms-aggregation.asciidoc#L224>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "aggs": {
//     "genres": {
//       "terms": {
//         "field": "genre",
//         "order": { "_key": "asc" }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_terms_aggregation_54a215d242ab65123b09e9dfb71bcbbf(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:54a215d242ab65123b09e9dfb71bcbbf[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "genres": {
		      "terms": {
		        "field": "genre",
		        "order": {
		          "_key": "asc"
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
	// end:54a215d242ab65123b09e9dfb71bcbbf[]
}
