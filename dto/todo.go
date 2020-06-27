package dto

// 数据传输对象。Data Transfer Object
// 将数据库的对象转换为前端需要的字段。
import "todo/models"

type TodoDTO struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
}

func ToTodoDTO(todo models.Todo) TodoDTO {
	return TodoDTO{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
	}
}

func ToTodoDTOs(todos []models.Todo) []TodoDTO {
	todoDTOs := make([]TodoDTO, len(todos))
	for i, item := range todos {
		todoDTOs[i] = ToTodoDTO(item)
	}
	return todoDTOs
}
