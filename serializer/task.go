package serializer

import (
	"github.com/jinzhu/gorm"
	"todo_list/model"
)

type Task struct {
	gorm.Model
	//User      User   `gorm:"ForeignKey:Uid"`
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	Content   string `json:"content"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

func BuildTask(task model.Task) Task {
	return Task{
		//Id: gorm.Model{ID}
		ID:        task.ID,
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
