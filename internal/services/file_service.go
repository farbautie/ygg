package services

import (
	"errors"
	"io"

	"github.com/farbautie/ygg/pkg/storage"
)

type FileService struct {
	storage storage.Storage
}

func NewFileService(storage storage.Storage) *FileService {
	return &FileService{
		storage: storage,
	}
}

func (s *FileService) Save(path string, data io.Reader) error {
	if path == "" {
		return errors.New("path cannot be empty")
	}

	return s.storage.Save(path, data)
}

func (s *FileService) Read(path string) (io.ReadCloser, error) {
	if path == "" {
		return nil, errors.New("path cannot be empty")
	}

	return s.storage.Read(path)
}

func (s *FileService) Delete(path string) error {
	if path == "" {
		return errors.New("path cannot be empty")
	}

	return s.storage.Delete(path)
}

func (s *FileService) List(path string) ([]string, error) {
	if path == "" {
		return nil, errors.New("path cannot be empty")
	}

	return s.storage.List(path)
}
