package todo

type RepositoryIface interface {
}

type Repository struct {
}

func NewRepository() RepositoryIface {
	return &Repository{}
}
