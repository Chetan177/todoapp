package rest

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"todo/pkg/model"
)

func (r *Rest) health(c echo.Context) error {
	res := &model.ApiReturn{Message: "server is up and running"}

	return c.JSON(http.StatusOK, res)
}

func (r *Rest) createTask(c echo.Context) error {
	res := &model.ApiReturn{Message: "success"}
	t := &model.Task{}

	err := c.Bind(t)
	if err != nil {
		res.Message = "validation failed"
		return c.JSON(http.StatusBadRequest, res)
	}
	log.Println("got request createTask:: ", t)
	tid, _ := uuid.NewUUID()
	t.ID = tid.String()
	id, err := r.db.CreateTask(t)
	if err != nil {
		res.Message = "unable to insert into db " + err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	res.Id = id
	return c.JSON(http.StatusOK, res)
}

func (r *Rest) deleteTask(c echo.Context) error {
	res := &model.ApiReturn{Message: "success"}
	id := c.Param("id")
	if id == "" {
		res.Message = "validation failed id is empty"
		return c.JSON(http.StatusBadRequest, res)
	}

	log.Println("got request deleteTask:: ", id)

	count, err := r.db.DeleteTask(id)
	if err != nil {
		res.Message = "unable to delete " + err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	if count > 0 {
		return c.JSON(http.StatusOK, res)
	}
	res.Message = "no item to delete"
	return c.JSON(http.StatusOK, res)
}

func (r *Rest) getTask(c echo.Context) error {
	res := &model.ApiReturn{Message: "success"}
	id := c.Param("id")
	if id == "" {
		res.Message = "validation failed id is empty"
		return c.JSON(http.StatusBadRequest, res)
	}

	log.Println("got request getTask:: ", id)

	task, err := r.db.GetTask(id)
	if err != nil {
		res.Message = "unable to get " + err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	return c.JSON(http.StatusOK, task)
}

func (r *Rest) getAllTask(c echo.Context) error {
	res := &model.ApiReturn{Message: "success"}
	owner := c.Param("owner")
	if owner == "" {
		res.Message = "validation failed id is empty"
		return c.JSON(http.StatusBadRequest, res)
	}

	log.Println("got request getAllTask:: ", owner)

	task, err := r.db.GetAllTask(owner)
	if err != nil {
		res.Message = "unable to get " + err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	return c.JSON(http.StatusOK, task)
}

func (r *Rest) markDone(c echo.Context) error {
	res := &model.ApiReturn{Message: "success"}
	id := c.Param("id")
	done := c.Param("done")
	b, _ := strconv.ParseBool(done)
	if id == "" {
		res.Message = "validation failed id is empty"
		return c.JSON(http.StatusBadRequest, res)
	}

	log.Println("got request markDone:: ", id, " ", done)

	err := r.db.MarkTaskDone(id, b)
	if err != nil {
		res.Message = "unable to get " + err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	res.Id = id
	return c.JSON(http.StatusOK, res)
}
