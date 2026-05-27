package files

import (
	"sync"

	"github.com/lytics/cloudstorage"
	"github.com/lytics/cloudstorage/google"
	"github.com/lytics/cloudstorage/localfs"

	"github.com/araddon/qlbridge/schema"
)

var (
	// TODO:   move to test files
	localFilesConfig = cloudstorage.Config{
		Type:       localfs.StoreType,
		AuthMethod: localfs.AuthFileSystem,
		LocalFS:    "./tables",
		TmpDir:     "/tmp/localcache",
	}

	// TODO:   complete manufacture this from config
	gcsConfig = cloudstorage.Config{
		Type:       google.StoreType,
		AuthMethod: google.AuthGCEDefaultOAuthToken,
		Project:    "lytics-dev",
		Bucket:     "lytics-dataux-tests",
		TmpDir:     "/tmp/localcache",
	}
)

var (
	// the global filestore registry mutex
	fileStoreMu sync.Mutex
	fileStores  = make(map[string]FileStoreCreator)
)

func init() {
	RegisterFileStore("gcs", createGCSFileStore)
	RegisterFileStore("localfs", createLocalFileStore)
}

// FileStoreLoader defines the interface for loading files
func FileStoreLoader(ss *schema.Schema) (cloudstorage.StoreReader, error) {
	_ = "STUB: not implemented"
	return *new(cloudstorage.StoreReader), nil
}

//u.Debugf("json conf:\n%s", ss.Conf.Settings.PrettyJson())

// FileStoreCreator defines a Factory type for creating FileStore
type FileStoreCreator func(*schema.Schema) (FileStore, error)

// FileStore Defines handler for reading Files, understanding
// folders and how to create scanners/formatters for files.
// Created by FileStoreCreator
//
// FileStoreCreator(schema) -> FileStore
//
//	FileStore.Objects() -> File
//	         FileHandler(File) -> FileScanner
//	              FileScanner.Next() ->  Row
type FileStore interface {
	cloudstorage.StoreReader
}

// RegisterFileStore global registry for Registering
// implementations of FileStore factories of the provided @storeType
func RegisterFileStore(storeType string, fs FileStoreCreator) { _ = "STUB: not implemented"; return }

func createGCSFileStore(ss *schema.Schema) (FileStore, error) {
	_ = "STUB: not implemented"
	return *new(FileStore), nil
}

// We don't actually need the gs:// because cloudstore does it

func createLocalFileStore(ss *schema.Schema) (FileStore, error) {
	_ = "STUB: not implemented"
	return *new(FileStore), nil
}
