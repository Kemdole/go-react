package todo

import "go-react/common"

type Service struct {
	repository RepositoryIface
}

func NewService(rep RepositoryIface) *Service {
	return &Service{
		repository: rep,
	}
}

func (s *Service) CreateTodo(req CreateTodoRequest) (*CreateTodoResponse, *common.Error) {
	return &CreateTodoResponse{}, nil
}

func (s *Service) GetTodo() (*GetTodoResponse, *common.Error) {
	return &GetTodoResponse{}, nil
}

func (s *Service) GetTodoList() (*ListTodoResponse, *common.Error) {
	return &ListTodoResponse{}, nil
}

func (s *Service) UpdateTodo(req UpdateTodoRequest) (*UpdateTodoResponse, *common.Error) {
	return &UpdateTodoResponse{}, nil
}

func (s *Service) DeleteTodo() *common.Error {
	return nil
}
