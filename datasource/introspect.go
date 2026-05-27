package datasource

import (
	"github.com/araddon/qlbridge/schema"
)

var (
	// IntrospectCount is default number of rows to evaluate for introspection
	// based schema discovery.
	IntrospectCount = 20
)

// IntrospectSchema discover schema from contents of row introspection.
func IntrospectSchema(s *schema.Schema, name string, iter schema.Iterator) error {
	_ = "STUB: not implemented"
	return nil
}

// IntrospectTable accepts a table and schema Iterator and will
// read a representative sample of rows, introspecting the results
// to create a schema.  Generally used for CSV, Json files to
// create strongly typed schemas.
func IntrospectTable(tbl *schema.Table, iter schema.Iterator) error {
	_ = "STUB: not implemented"
	return nil
}

//u.Infof("s:%s INTROSPECT SCHEMA name %q", s, name)

//u.Debugf("msg %#v", msg)

//u.Debugf("i:%v k:%s  v: %T %v", i, k, v, v)

//fld := tbl.FieldMap[k]
//u.Debugf("add field? %+v", fld)
//u.Debugf("%s = %v   type: %T   vt:%s new? %v", k, val, val, valType, !exists)

// if k == "" {
// 	for k2, ki := range mt.ColIndex {
// 		if ki == i {
// 			k = k2
// 			break
// 		}
// 	}
// }

//u.Debugf("i:%v k:%s  v: %T %v", i, k, v, v)

//fld := tbl.FieldMap[k]
//u.Debugf("add field? %+v", fld)
//u.Debugf("%s = %v   type: %T   vt:%s new? %v", k, val, val, valType, !exists)

// hm.....

//u.Debugf("%+v", f)

//u.Debugf("%s: %v", tbl.Name, tbl.Columns())
