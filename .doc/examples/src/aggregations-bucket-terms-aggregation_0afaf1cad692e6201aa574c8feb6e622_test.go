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

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/terms-aggregation.asciidoc#L626>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "aggs" : {
//         "tags" : {
//             "terms" : {
//                 "field" : "tags",
//                 "include" : ".*sport.*",
//                 "exclude" : "water_.*"
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_terms_aggregation_0afaf1cad692e6201aa574c8feb6e622(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:0afaf1cad692e6201aa574c8feb6e622[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "tags": {
		      "terms": {
		        "field": "tags",
		        "include": ".*sport.*",
		        "exclude": "water_.*"
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
	// end:0afaf1cad692e6201aa574c8feb6e622[]
}
