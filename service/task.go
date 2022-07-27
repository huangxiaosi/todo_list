package service

import (
	"time"
	"todo_list/model"
	"todo_list/serializer"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}

type ShowTaskService struct {
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}

//新增一条
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	var code = 200
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "备忘录创建错误。",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "创建成功。",
	}
}

//展示一条
func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败。",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    "查询成功。",
	}
}

//放回用户所有
func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	//return serializer.Response{
	//	Status: 200,
	//	Data:   serializer.BuildTasks(tasks),
	return serializer.BuildListPresponse(serializer.BuildTasks(tasks), uint(count))
}

//更新
func (service UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	model.DB.First(&task, tid)
	task.Title = service.Title
	task.Content = service.Content
	task.Status = service.Status
	model.DB.Save(&task)
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
		Msg:    "更新完成。",
	}
}
