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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/datehistogram-aggregation.asciidoc#L138>
//
// --------------------------------------------------------------------------------
// POST /sales/_search?size=0
// {
//   "aggs": {
//     "sales_over_time": {
//       "date_histogram": {
//         "field": "date",
//         "calendar_interval": "2d"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_datehistogram_aggregation_f5815d573cee0447910c9668003887b8(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:f5815d573cee0447910c9668003887b8[]
	res, err := es.Search(
		es.Search.WithIndex("sales"),
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "sales_over_time": {
		      "date_histogram": {
		        "field": "date",
		        "calendar_interval": "2d"
		      }
		    }
		  }
		}`)),
		es.Search.WithSize(0),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:f5815d573cee0447910c9668003887b8[]
}
