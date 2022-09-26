package appstorage

import "os"

func createDirIfNotExist(absPath string) error {

	if absPath == "" {
		return nil
	}

	if err := os.MkdirAll(absPath, 0770); err != nil {
		return err
	}

	return nil
}
