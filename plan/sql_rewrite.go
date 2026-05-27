package plan

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/rel"
	"github.com/araddon/qlbridge/value"
)

var _ = u.EMPTY

var fr = expr.NewFuncRegistry()

func init() {
	fr.Add("typewriter", &defaultTypeWriter{})
}

type defaultTypeWriter struct{}

// defaultTypeWriter Convert a qlbridge value type to qlbridge value type
func (m *defaultTypeWriter) Eval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

func (m *defaultTypeWriter) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func (m *defaultTypeWriter) IsAgg() bool { _ = "STUB: not implemented"; return false }
func (m *defaultTypeWriter) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *

	// RewriteShowAsSelect Rewrite Schema SHOW Statements AS SELECT statements
	// so we only need a Select Planner, not separate planner for show statements
	new(value.ValueType)
}

func RewriteShowAsSelect(stmt *rel.SqlShow, ctx *Context) (*rel.SqlSelect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SHOW FULL TABLES;    = select name, table_type from tables;
// TODO:  note the stupid "_in_mysql", assuming i don't have to implement
/*
   mysql> show full tables;
   +---------------------------+------------+
   | Tables_in_mysql           | Table_type |
   +---------------------------+------------+
   | columns_priv              | BASE TABLE |

*/

// show tables;

// SHOW CREATE {TABLE | DATABASE | EVENT | VIEW }

// SHOW databases;  ->  select Database from databases;

/*
	mysql> show full columns from user;
	+------------------------+-----------------------------------+-----------------+------+-----+-----------------------+-------+---------------------------------+---------+
	| Field                  | Type                              | Collation       | Null | Key | Default               | Extra | Privileges                      | Comment |

*/

/*
	mysql> show columns from user;
	+------------------------+-----------------------------------+------+-----+-----------------------+-------+
	| Field                  | Type                              | Null | Key | Default               | Extra |
	+------------------------+-----------------------------------+------+-----+-----------------------+-------+
*/

/*
	mysql> show keys from `user` from `mysql`;
	+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
	| Table | Non_unique | Key_name | Seq_in_index | Column_name | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment |
	+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
	| user  |          0 | PRIMARY  |            1 | Host        | A         |        NULL |     NULL | NULL   |      | BTREE      |         |               |
	| user  |          0 | PRIMARY  |            2 | User        | A         |           3 |     NULL | NULL   |      | BTREE      |         |               |
	+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+

	mysql> show indexes from `user` from `mysql`;
	+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
	| Table | Non_unique | Key_name | Seq_in_index | Column_name | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment |
	+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
	| user  |          0 | PRIMARY  |            1 | Host        | A         |        NULL |     NULL | NULL   |      | BTREE      |         |               |
	| user  |          0 | PRIMARY  |            2 | User        | A         |           3 |     NULL | NULL   |      | BTREE      |         |               |
	+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+

*/

// SHOW [GLOBAL | SESSION] VARIABLES [like_or_where]

/*
   mysql> show variables LIKE 'version';
   +---------------+----------+
   | Variable_name | Value    |
   +---------------+----------+
   | version       | 5.7.10-3 |
   +---------------+----------+
*/

// Status is a subset of just some variables
// http://dev.mysql.com/doc/refman/5.7/en/server-status-variables.html

// SHOW [GLOBAL | SESSION | SLAVE ] STATUS [like_or_where]

/*
	mysql> show global status;
	+--------------------------------+-----------------+
	| Variable_name                  | Value
	+--------------------------------+------------------
	| Aborted_clients                | 0
	| Aborted_connects               | 0
	| Binlog_snapshot_file           |
	| Binlog_snapshot_position       | 0
*/

/*
	show engines;
	mysql> show engines;
	+--------------------+---------+----------------------------------------------------------------------------+--------------+------+------------+
	| Engine             | Support | Comment                                                                    | Transactions | XA   | Savepoints |
	+--------------------+---------+----------------------------------------------------------------------------+--------------+------+------------+
	| InnoDB             | DEFAULT | Percona-XtraDB, Supports transactions, row-level locking, and foreign keys | YES          | YES  | YES        |
	| CSV                | YES     | CSV storage engine                                                         | NO           | NO   | NO         |
	| MyISAM             | YES     | MyISAM storage engine                                                      | NO           | NO   | NO         |
	| BLACKHOLE          | YES     | /dev/null storage engine (anything you write to it disappears)             | NO           | NO   | NO         |
	| PERFORMANCE_SCHEMA | YES     | Performance Schema                                                         | NO           | NO   | NO         |
	| MEMORY             | YES     | Hash based, stored in memory, useful for temporary tables                  | NO           | NO   | NO         |
	| ARCHIVE            | YES     | Archive storage engine                                                     | NO           | NO   | NO         |
	| MRG_MYISAM         | YES     | Collection of identical MyISAM tables                                      | NO           | NO   | NO         |
	| FEDERATED          | NO      | Federated MySQL storage engine                                             | NULL         | NULL | NULL       |
	+--------------------+---------+----------------------------------------------------------------------------+--------------+------+------------+
*/

/*
	show procuedure status;
	show function status;

		| Db  | Name | Type | Definer | Modified | Created | Security_type | Comment| character_set_client | collation_connection | Database Collation |
*/

// We are going to ReWrite LIKE clause to WHERE clause

// See if the Like Clause has wildcard matching, if so
// our internal vm uses * not %

//u.Debugf("add where: %s", stmt.Where)

func RewriteDescribeAsSelect(stmt *rel.SqlDescribe, ctx *Context) (*rel.SqlSelect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
