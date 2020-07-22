package utils

import (
	"os"
	"path"
)

func IsFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func CreateFile(filename string) (*os.File, error) {
	if IsFileExists(filename) { // 存在的话，删除
		err := os.Remove(filename)
		if err != nil {
			return nil, err
		}
	}

	err := os.MkdirAll(path.Dir(filename), os.ModePerm)
	if err != nil {
		return nil, err
	}

	return os.Create(filename)
}

