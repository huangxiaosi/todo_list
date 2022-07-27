package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"todo_list/pkg/utils"
	"todo_list/service"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	//claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(c.Param("tid"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func ListTasks(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	//claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateTask); err == nil {
		res := updateTask.Update(c.Param("tid"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask); err == nil {
		res := searchTask.Search(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	//claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(c.Param("tid"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}
