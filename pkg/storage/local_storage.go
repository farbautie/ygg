package storage

import (
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	Path string
}

func NewLocalStorage(path string) (*LocalStorage, error) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, err
	}

	return &LocalStorage{
		Path: path,
	}, nil
}

func (s *LocalStorage) Save(path string, data io.Reader) error {
	fullPath := filepath.Join(s.Path, path)
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, data)
	return err
}

func (s *LocalStorage) Read(path string) (io.ReadCloser, error) {
	return nil, nil
}

func (s *LocalStorage) Delete(path string) error {
	return nil
}

func (s *LocalStorage) List(path string) ([]string, error) {
	return nil, nil
}
