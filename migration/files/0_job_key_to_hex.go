package files

import (
	logging "github.com/ipfs/go-log"
	"github.com/syndtr/goleveldb/leveldb"
)

var log = logging.Logger("migrate-files")

// RunMigration0 Job key to hex
func RunMigration0(db *leveldb.DB) error {
	err := db.Put([]byte("perico"), []byte("sample"), nil)
	if err != nil {
		return err
	}
	log.Infof("Migration 0 Run successfully")
	return nil
}

