package todo

import "go-react/common/db"

type RepositoryIface interface {
}

type Repository struct {
	dao db.Dao
}

func NewRepository(dao db.Dao) RepositoryIface {
	return &Repository{
		dao: dao,
	}
}
