package appstorage

import (
	"os"
	"path"

	"github.com/syndtr/goleveldb/leveldb"
)

type levelDbStorage struct {
	db *leveldb.DB
}

func NewLevelDbStorage(appName string, dbName string) (*levelDbStorage, error) {

	dir, err := os.UserConfigDir()

	if err != nil {
		return nil, err
	}

	dbPath := path.Join(dir, appName, dbName)

	db, err := leveldb.OpenFile(dbPath, nil)

	if err != nil {
		return nil, err
	}

	return &levelDbStorage{
		db: db,
	}, nil
}

func (storage *levelDbStorage) Get(key string) (string, error) {
	val, err := storage.db.Get([]byte(key), nil)

	if err != nil {
		return "", err
	}

	return string(val), nil
}

func (storage *levelDbStorage) Put(key string, data string) error {
	return storage.db.Put([]byte(key), []byte(data), nil)
}

func (storage *levelDbStorage) Delete(key string) error {
	return storage.db.Delete([]byte(key), nil)
}

func (storage *levelDbStorage) Close() error {
	return storage.db.Close()
}
