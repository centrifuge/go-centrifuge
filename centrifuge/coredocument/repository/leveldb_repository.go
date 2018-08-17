package coredocumentrepository

import (
	"sync"

	"github.com/CentrifugeInc/centrifuge-protobufs/gen/go/coredocument"
	"github.com/golang/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb"
)

var once sync.Once

type LevelDBCoreDocumentRepository struct {
	LevelDB *leveldb.DB
}

func NewLevelDBCoreDocumentRepository(cdr CoreDocumentRepository) {
	once.Do(func() {
		coreDocumentRepository = cdr
	})
	return
}

func (repo *LevelDBCoreDocumentRepository) GetKey(id []byte) []byte {
	return append([]byte("coredoc"), id...)
}

func (repo *LevelDBCoreDocumentRepository) FindById(id []byte) (doc *coredocumentpb.CoreDocument, err error) {
	docBytes, err := repo.LevelDB.Get(repo.GetKey(id), nil)
	if err != nil {
		return nil, err
	}

	doc = &coredocumentpb.CoreDocument{}
	err = proto.Unmarshal(docBytes, doc)
	if err != nil {
		return nil, err
	}
	return
}

func (repo *LevelDBCoreDocumentRepository) CreateOrUpdate(doc *coredocumentpb.CoreDocument) (err error) {
	key := repo.GetKey(doc.DocumentIdentifier)
	data, err := proto.Marshal(doc)

	if err != nil {
		return
	}
	err = repo.LevelDB.Put(key, data, nil)
	return
}
