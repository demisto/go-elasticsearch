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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/setup/logging-config.asciidoc#L155>
//
// --------------------------------------------------------------------------------
// PUT /_cluster/settings
// {
//   "transient": {
//     "logger.org.elasticsearch.transport": "trace"
//   }
// }
// --------------------------------------------------------------------------------

func Test_setup_logging_config_8e6bfb4441ffa15c86d5dc20fa083571(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:8e6bfb4441ffa15c86d5dc20fa083571[]
	res, err := es.Cluster.PutSettings(
		strings.NewReader(`{
		  "transient": {
		    "logger.org.elasticsearch.transport": "trace"
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:8e6bfb4441ffa15c86d5dc20fa083571[]
}
