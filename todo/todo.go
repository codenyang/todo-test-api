package todo

type ToDo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Check bool   `json:"check"`
}

var (
	todos []ToDo
	index = 0
)

func AddTodo(todo ToDo) {
	todo.ID = index
	index++
	todos = append(todos, todo)
}

func CheckTodo(todo ToDo) bool {
	for i, v := range todos {
		if v.ID == todo.ID {
			todos[i].Check = !todos[i].Check
			return true
		}
	}
	return false
}

func DelTodo(todo ToDo) bool {
	for i, v := range todos {
		if v.ID == todo.ID {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}

func GetTodos() []ToDo {
	return todos
}
