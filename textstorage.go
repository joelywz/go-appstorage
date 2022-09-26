package appstorage

import (
	"fmt"
	"os"
	"path"
)

type textStorage struct {
	baseDir string
}

func NewAppStorage(appName string) (*textStorage, error) {

	dir, err := os.UserConfigDir()

	if err != nil {
		return nil, err
	}

	return &textStorage{
		baseDir: fmt.Sprintf("%s/%s", dir, appName),
	}, nil
}

func (storage *textStorage) Save(relPath string, data string) error {

	absPath := storage.getAbsolutePath(relPath)

	createDirIfNotExist(path.Join(absPath, ".."))

	file, err := os.Create(absPath)

	if err != nil {
		return err
	}

	defer func() {
		file.Close()
	}()

	_, err = file.WriteString(data)

	if err != nil {
		return err
	}

	return nil
}

func (stroage *textStorage) Read(relPath string) (string, error) {
	absPath := stroage.getAbsolutePath(relPath)

	b, err := os.ReadFile(absPath)

	if err != nil {
		return "", err
	}

	return string(b), err
}

func (storage *textStorage) getAbsolutePath(relPath string) string {
	return path.Join(storage.baseDir, relPath)
}
