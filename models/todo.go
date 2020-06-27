package models

import (
	"time"
	"todo/config"

	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"type:char(18);not null"`
	Description string    `json:"desc"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:null ON UPDATE CURRENT_TIMESTAMP"` // TODO: 不知道为什么，gorm 一定会一个当前时间。
	DeletedAt   time.Time `json:"deleted_at" gorm:"default:null"`                             // 这里不给 default:null 也会报错
}

func (b *Todo) TableName() string {
	return "todo"
}

func GetAllTodos() (todo []Todo, err error) {
	if err = config.DB.Find(&todo).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

// create a todo
func CreateATodo(title string) (todo Todo, err error) {
	todo = Todo{
		Title: title,
	}
	if err = config.DB.Create(&todo).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

func GetATodo(todo *Todo, id string) (err error) {
	if err = config.DB.First(todo, id).Error; err != nil {
		return err
	}
	return nil
}

// delete a todo（soft delete）
func DeleteATodo(id string) (rowsAffected int64, err error) {
	// config.DB.First(&produ)
	var todo Todo
	// RowsAffected
	db := config.DB.Where("id = ?", id).Delete(&todo)
	err = db.Error
	return db.RowsAffected, err
}

func UpdateATodo(id uint, title string) (err error) {
	todo := Todo{
		ID: id,
	}
	return config.DB.Model(&todo).Update("title", title).Error
}
