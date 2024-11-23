package storage

import "io"

type Storage interface {
	Save(path string, data io.Reader) error
	Read(path string) (io.ReadCloser, error)
	Delete(path string) error
	List(path string) ([]string, error)
}
