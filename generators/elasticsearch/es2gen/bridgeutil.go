package es2gen

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

// scalar returns a JSONable representation of a scalar node type for use in ES
// filters.
//
// Does not support Null.
func scalar(node expr.Node) (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// ES supports string encoded ints

// Make sure this is a scalar value node

// makeRange returns a range filter for Elasticsearch given the 3 nodes that
// make up a comparison.
func makeRange(lhs *gentypes.FieldType, op lex.TokenType, rhs expr.Node) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert scalars from strings to floats if lhs is numeric and rhs is a
// float (ES handles ints as strings just fine).

// rhsval can be converted to a float!

/*
	"nested": {
		"filter": {
		    "term": {
		        "map_actioncounts.k": "Web hit"
		    }
		},
		"path": "map_actioncounts"
	}

	"nested": {
		"filter": {
		    "and": [
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
		},
		"path": "mapvals_fields"
	}
	"nested": {
		"filter": {
		    "and": [
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
			"filter": {
			    "and": [
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
			},
			"path": "map_events"
		}

		"and": [
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
//	{"query": {"wildcard": {field: value}}}
func makeWildcard(lhs *gentypes.FieldType, value string) (interface{}, error) {
	_ = "STUB: not implemented"
	/*
		"nested": {
			"filter": {
			    "and": [
			        {
			            "term": { "map_events.k": "open" }
			        },
			        { "wildcard": {"map_events.v": "hel"}
			        }
			    ]
			},
			"path": "map_events"
		}

		{"query": {"wildcard": {field: value}}}
	*/return nil, nil
}

// makeTimeWindowQuery maps the provided threshold and window arguments to the indexed time buckets
func makeTimeWindowQuery(lhs *gentypes.FieldType, threshold, window, ts int64) (interface{}, error) {
	_ = "STUB: not implemented"
	/*
		"nested": {
			"filter": {
				"and": [
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
			"path": "timebucket_visits"
		}
	*/return nil, nil
}
