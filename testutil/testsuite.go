package testutil

import (
	"sync"
)

var runInitOnce = sync.Once{}

func runInit() { _ = "STUB: not implemented"; return }

//panic("wtf")

// don't re-setup

// make sure logging is always non-nil

// load our mock data sources "users", "orders"

// RunDDLTests run the DDL (CREATE SCHEMA, TABLE, alter) test harness suite.
func RunDDLTests(t TestingT) {
	_ = "STUB: not implemented"

	// DDL
	return
}

// RunTestSuite run the normal DML SQL test suite.
func RunTestSuite(t TestingT) {
	_ = "STUB: not implemented"

	// Literal Queries
	return
}

// Slightly different test method allows source

// - yy func evaluates
// - projection (user_id, email)

// - ensure we can evaluate against "NULL"
// - extra paren in where
// - `db`.`col` syntax

// Mixed *, literal, fields

// - user_id != NULL (on string column)
// - as well as count(*)

// Aliasing columns in group by

// nested functions in aggregations
// - note also lack of group-by ie determine Is Agg query in rel ast

// aaron@email.combob@email.comnot_an_email_2 = 42 characters / 3 = 14

// Distinct keyword

// This is an error because we have schema on this table, and this column
// doesn't exist.

/*
	// TODO: #56 DISTINCT inside count()
	testutil.TestSelect(t, "SELECT COUNT(DISTINCT(`users.user_id`)) AS cd FROM users",
		[][]driver.Value{{int64(3)}},
	)

	// TODO: #56 this doesn't work because ordering is non-deterministic coming out of group by currently
	//  which technically don't think there is any sql expectation of ordering, but there is for this test harness
	testutil.TestSelect(t, "select `users`.`user_id` AS userids FROM users GROUP BY `users`.`user_id`;",
		[][]driver.Value{{"hT2impsabc345c"}, {"9Ip1aKbeZe2njCDM"}, {"hT2impsOPUREcVPc"}},
	)
*/

// RunSimpleSuite run the normal DML SQL test suite.
func RunSimpleSuite(t TestingT) {
	_ = "STUB: not implemented"

	// // Function in select projected columns that needs to be late evaluated.
	// // "select json.jmespath(body,\"name\") AS name FROM article WHERE `author` = \"aaron\";",
	// TestSelect(t, "select json.jmespath(json_data,\"name\") AS name FROM users WHERE `email` = \"aaron@email.com\";",
	//
	//	[][]driver.Value{{"aaron"}},
	//
	// )
	// return
	return
}

// Literal Queries

// - ensure we can evaluate against "NULL"
// - extra paren in where
// - `db`.`col` syntax

// - user_id != NULL (on string column)
// - as well as count(*)

// Aliasing columns in group by

// nested functions in aggregations
// - note also lack of group-by ie determine Is Agg query in rel ast

// aaron@email.combob@email.comnot_an_email_2 = 42 characters / 3 = 14

// Distinct keyword

// Function in select projected columns that needs to be late evaluated.
// "select json.jmespath(body,\"name\") AS name FROM article WHERE `author` = \"aaron\";",

// This is an error because we have schema on this table, and this column
// doesn't exist.

/*
	// TODO: #56 DISTINCT inside count()
	testutil.TestSelect(t, "SELECT COUNT(DISTINCT(`users.user_id`)) AS cd FROM users",
		[][]driver.Value{{int64(3)}},
	)

	// TODO: #56 this doesn't work because ordering is non-deterministic coming out of group by currently
	//  which technically don't think there is any sql expectation of ordering, but there is for this test harness
	testutil.TestSelect(t, "select `users`.`user_id` AS userids FROM users GROUP BY `users`.`user_id`;",
		[][]driver.Value{{"hT2impsabc345c"}, {"9Ip1aKbeZe2njCDM"}, {"hT2impsOPUREcVPc"}},
	)
*/
