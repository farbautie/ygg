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
