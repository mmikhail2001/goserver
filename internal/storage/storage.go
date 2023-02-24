package storage

type Auth interface {
	CreateUser(user mode.User) (int, error)
}

type storage struct {
	Auth
}

func NewRepository() *storage {
	return &storage{}
}
