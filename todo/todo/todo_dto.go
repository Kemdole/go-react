package todo

type TodoDto struct {
	ID string
}

type CreateTodoRequest struct {
}

type CreateTodoResponse struct {
	TodoDto
}

type GetTodoResponse struct {
	TodoDto
}

type ListTodoResponse struct {
	Todos []TodoDto
}

type UpdateTodoRequest struct {
}

type UpdateTodoResponse struct {
	TodoDto
}
