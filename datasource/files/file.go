package files

import (
	"database/sql/driver"
	"io"
	"time"

	u "github.com/araddon/gou"
	"github.com/lytics/cloudstorage"
)

var (
	// FileColumns are the default columns for the "file" table
	FileColumns = []string{"file", "table", "path", "partialpath", "size", "partition", "updated", "deleted", "filetype"}

	_ = u.EMPTY
)

type FileTable struct {
	Table       string
	PartialPath string
	FileCount   int
}

// FileInfo describes a single file
// Say a folder of "./tables" is the "root path" specified
// then say it has folders for "table names" underneath a redundant "tables"
// ./tables/
//
//	/tables/
//	       /appearances/appearances1.csv
//	       /players/players1.csv
//
// Name = "tables/appearances/appearances1.csv"
// Table = "appearances"
// PartialPath = tables/appearances
type FileInfo struct {
	obj         cloudstorage.Object
	Path        string         // Root path
	Name        string         // Name/Path of file
	PartialPath string         // non-file-name part of path
	Table       string         // Table name this file participates in
	FileType    string         // csv, json, etc
	Partition   int            // which partition
	Size        int            // Content-Length size in bytes
	AppendCols  []driver.Value // Additional Column info extracted from file name/folder path
}

// FileReader file info and access to file to supply to ScannerMakers
type FileReader struct {
	*FileInfo
	F    io.ReadCloser // Actual file reader
	Exit chan bool     // exit channel to shutdown reader
}

func (m *FileInfo) String() string { _ = "STUB: not implemented"; return "" }

// Values as as slice, create a row of values describing this file
// for use in sql listing of files
func (m *FileInfo) Values() []driver.Value { _ = "STUB: not implemented"; return nil }

func (m *FileInfo) updated() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Convert a cloudstorage object to a File.  Interpret the table name
// for given full file path.
func FileInfoFromCloudObject(path string, obj cloudstorage.Object) *FileInfo {
	_ = "STUB: not implemented"
	return nil
}

// Get the part of path as follows
//  /path/partialpath/filename.csv

//u.Debugf("Fi: name=%q table=%q  partial:%q partial2:%q", fi.Name, fi.Table, fi.PartialPath, partialPath)

// Find table name from full path of file, and the path of tables.
//
// There are different "table" naming conventions we use to
// find table names.
//
//  1. Support multiple partitioned files in folder which is name of table
//     rootpath/tables/nameoftable/nameoftable1.csv
//     rootpath/tables/nameoftable/nameoftable2.csv
//
//  2. Suport Table as name of file inside folder
//     rootpath/users.csv
//     rootpath/accounts.csv
func TableFromFileAndPath(path, fileIn string) string { _ = "STUB: not implemented"; return "" }

// the path prefix was supplied as part of config to tell us where
// the files are, so we can safely strip it out

// We are going to strip the tables prefix

//u.Warnf("table not readable from filename %q  path=%q", fileIn, path)
