package Handler 

import(
    "net/http"
	"todo-api/internal/api"
	
)

type TodoHandler struct {
	
}

func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request, params api.ListTodosParams){

}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request, params api.CreateTodoParams){

}