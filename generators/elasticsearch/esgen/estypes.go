package esgen

import (
	"encoding/json"

	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/generators/elasticsearch/gentypes"
)

/*
Native go data types that map to the Elasticsearch
Search DSL
*/
var _ = u.EMPTY
var _ = json.Marshal

type BoolFilter struct {
	Occurs         BoolOccurrence `json:"bool"`
	MinShouldMatch int            `json:"minimum_should_match,omitempty"`
}

type BoolOccurrence struct {
	Filter  []interface{} `json:"filter,omitempty"`
	Should  []interface{} `json:"should,omitempty"`
	MustNot interface{}   `json:"must_not,omitempty"`
}

func AndFilter(v []interface{}) *BoolFilter { _ = "STUB: not implemented"; return nil }
func OrFilter(v []interface{}) *BoolFilter  { _ = "STUB: not implemented"; return nil }
func NotFilter(v interface{}) *BoolFilter   { _ = "STUB: not implemented"; return nil }

// Filter structs

type exists struct {
	Exists map[string]string `json:"exists"`
}

// Exists creates a new Elasticsearch filter {"exists": {"field": field}}
func Exists(field *gentypes.FieldType) interface{} {
	_ = "STUB: not implemented"
	// u.Debugf("exists?  nested?%v  for %s", field.Nested(), field.String())
	return nil
}

/*
	"nested": {
		"query": {
		    "term": {
		        "map_actioncounts.k": "Web hit"
		    }
		},
		"path": "map_actioncounts"
	}
*/

//Nested(field.Path, &term{map[string][]string{"k": field.Field}})

//	type and struct {
//		Filters []interface{} `json:"and"`
//	}
type boolean struct {
	Bool interface{} `json:"bool"`
}
type must struct {
	Filters []interface{} `json:"must"`
}

type in struct {
	Terms map[string][]interface{} `json:"terms"`
}

// In creates a new Elasticsearch terms filter
//
// {"terms": {field: values}}
//
//	{ "nested": {
//	     "query": {
//	        "bool" : {
//	           "must" :[
//	              {"term": { "k":fieldName}},
//	              filter,
//	           ]
//	     } ,
//	     "path":"path_to_obj"
//	 }}
func In(field *gentypes.FieldType, values []interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Nested creates a new Elasticsearch nested filter
//
//	{ "nested": {
//	     "query": {
//	        "bool" : {
//	           "must" :[
//	              {"term": { "k":fieldName}},
//	              filter,
//	           ]
//	     } ,
//	     "path":"path_to_obj"
//	 }}
func Nested(field *gentypes.FieldType, filter interface{}) *nested {
	_ = "STUB: not implemented"

	// Hm.  Elasticsearch doc seems to insinuate we don't need
	// this path + ".k" but unit tests say otherwise
	return nil
}

// by, _ := json.MarshalIndent(n, "", "  ")
// u.Infof("NESTED4:  \n%s", string(by))

type nested struct {
	Nested *NestedQuery `json:"nested,omitempty"`
}

type NestedQuery struct {
	Query interface{} `json:"query"`
	Path  string      `json:"path"`
}

type RangeQry struct {
	GTE interface{} `json:"gte,omitempty"`
	LTE interface{} `json:"lte,omitempty"`
	GT  interface{} `json:"gt,omitempty"`
	LT  interface{} `json:"lt,omitempty"`
}

type RangeFilter struct {
	Range map[string]RangeQry `json:"range"`
}

type term struct {
	Term map[string]interface{} `json:"term"`
}

// Term creates a new Elasticsearch term filter {"term": {field: value}}
func Term(fieldName string, value interface{}) *term { _ = "STUB: not implemented"; return nil }

type matchall struct {
	MatchAll *struct{} `json:"match_all"`
}

// MatchAll maps to the Elasticsearch "match_all" filter
var MatchAll = &matchall{&struct{}{}}

// MatchNone matches no documents.
var MatchNone = NotFilter(MatchAll)

type wildcard struct {
	Wildcard map[string]string `json:"wildcard"`
}

func wcFunc(val string) string { _ = "STUB: not implemented"; return "" }

// Wilcard creates a new Elasticserach wildcard query
//
//	{"wildcard": {field: value}}
//
// nested
//
//	{"nested": {
//	   "filter" : { "and" : [
//	           {"wildcard": {"v": value}},
//	           {"term":{"k": field_key}}
//	   "path": path
//	  }
//	}
func Wildcard(field, value string) *wildcard { _ = "STUB: not implemented"; return nil }
