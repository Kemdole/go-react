package todo

type Service struct {
	repository RepositoryIface
}

func NewService(rep RepositoryIface) *Service {
	return &Service{
		repository: rep,
	}
}

func (s *Service) CreateTodo() {

}

func (s *Service) GetTodo() {

}

func (s *Service) GetTodoList() {

}

func (s *Service) UpdateTodo() {

}

func (s *Service) DeleteTodo() {

}
