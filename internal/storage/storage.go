package storage

type OperationStorager interface {
	Set(URL string) (string, error)
	Get(shortURL string) (string, error)
}
