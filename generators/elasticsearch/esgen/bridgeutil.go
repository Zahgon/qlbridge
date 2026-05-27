package esgen

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/generators/elasticsearch/gentypes"
	"github.com/araddon/qlbridge/lex"
)

var _ = u.EMPTY

type floatval interface {
	Float() float64
}

// makeRange returns a range filter for Elasticsearch given the 3 nodes that
// make up a comparison.
func makeRange(lhs *gentypes.FieldType, op lex.TokenType, rhs expr.Node) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert scalars to correct type

// TODO:  we might need to change the operator???
//  given lh identity "purchase_count" = int = 10
//  right hand side = float 9.7

// rhsval can be converted to a float!

/*
	"nested": {
		"query": {
		    "term": {
		        "map_actioncounts.k": "Web hit"
		    }
		},
		"path": "map_actioncounts"
	}

	"nested": {
		"query": {
		    "bool": {
		      "must": [
		          {
		              "term": {
		                  "mapvals_fields.k": "has_data"
		              }
		          },
		          {
		              "term": {
		                  "mapvals_fields.b": true
		              }
		          }
		      ]
		    }
		},
		"path": "mapvals_fields"
	}
	"nested": {
		"query": {
			"bool": {
				"must": [
					{
						"term": {
							"k": "open"
						}
					},
					{
						"range": {
							"f": {"gte": 7}
						}
					}
				]
			}
		},
		"path": "map_events"
	}
	q = esMap{"nested": esMap{"path": parent, "filter": esMap{"and": []esMap{
				{"term": esMap{parent + ".k": child}},
				{"range": esMap{parent + valuePath: esMap{esRangeOps[seg.SegType]: rhsNum}}},
			}}}}
*/

// makeBetween returns a range filter for Elasticsearch given the 3 nodes that
// make up a comparison.
func makeBetween(lhs *gentypes.FieldType, lower, upper interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	/*
		"nested": {
			"query": {
				"bool": {
					"must": [
						{
							"term": {
								"k": "open"
							}
						},
						{
							"range": {
								"f": {"gt": 7}
							}
						},
						{
							"range": {
								"f": {"lt": 15}
							}
						}
					]
				}
			},
			"path": "map_events"
		}

		"must": [
		    {
		        "range": {
		            "f": {"gt": 7}
		        }
		    },
		    {
		        "range": {
		            "f": {"lt": 15}
		        }
		    }
		]
	*/return nil, nil
}

// makeWildcard returns a wildcard/like query
//
//	{"wildcard": {field: value}}
func makeWildcard(lhs *gentypes.FieldType, value string) (interface{}, error) {
	_ = "STUB: not implemented"
	/*
		"nested": {
			"query": {
				"bool": {
					"must": [
						{
							"term": { "map_events.k": "open" }
						},
						{
							"wildcard": {"map_events.v": "hel"}
						}
					]
				}
			},
			"path": "map_events"
		}

		{"wildcard": {field: value}}
	*/return nil, nil
}

// makeTimeWindowQuery maps the provided threshold and window arguments to the indexed time buckets
func makeTimeWindowQuery(lhs *gentypes.FieldType, threshold, window, ts int64) (interface{}, error) {
	_ = "STUB: not implemented"
	/*
		"nested": {
			"query": {
			  "bool":{
				"must": [
					{
						"term": { "timebucket_visits.threshold": 1 }
					},
					{
						"term": { "timebucket_visits.window": 3 }
					},
					{
						"range": {
							"timebucket_visits.enter: { "lte": 16916 }
						}
					},
					{
						"range": {
							"timebucket_visits.exit: { "gte": 16916 }
						}
					},
				]
			  }
			}
			"path": "timebucket_visits"
		}
	*/return nil, nil
}
