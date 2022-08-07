package storageI

type StorageI interface {
	LoadIn([]byte) error
	Update([]byte) error
	Delete(string) error
}
