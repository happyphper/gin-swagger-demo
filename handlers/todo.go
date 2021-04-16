package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/happyphper/gin-swagger-demo/models"
	"github.com/happyphper/gin-swagger-demo/response"
	"net/http"
	"strconv"
	"time"
)

type TodoIndexQuery struct {
	Page int `form:"page"`
}
// TodoList
// @Summary 待办事项列表
// @Description 用于待办事项的列表展示，支持分页、筛选等
// @Tags Todo
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1) minimum(1) maximum(100)
// @Param ids query []string false "ID数组" collectionFormat(multi)
// @Success 200 {object} response.Response{code=int,data=[]models.Todo} "desc"
// @Router /todos [get]
func TodoIndex(ctx *gin.Context) {
	var query TodoIndexQuery
	if err := ctx.BindQuery(&query); err != nil {
		ctx.JSON(http.StatusOK, response.ValidationRes(err.Error()))
		return
	}

	var data []models.Todo

	for i := 0; i < 10; i++ {
		data = append(data, models.Todo{
			ID:        i,
			Title:     fmt.Sprintf("Title %d", i),
			Status:    0,
			CreatedAt: time.Now().Format(`2006-01-02 15:04:05`),
			UpdatedAt: time.Now().Format(`2006-01-02 15:04:05`),
		})
	}

	ctx.JSON(http.StatusOK, response.Success(data))
}
// TotoShow
// @Summary 待办事项详情
// @Description 用于待办事项的列表展示，支持分页、筛选等
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path string true "Todo id"
// @Success 200 {object} response.Response{code=int,data=models.Todo} "desc"
// @Router /todos/{id} [get]
func TodoShow(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ValidationRes(err.Error()))
		return
	}

	todo := models.Todo{
		ID:        id,
		Title:     fmt.Sprintf("Title %d", id),
		Status:    0,
		CreatedAt: time.Now().Format(`2006-01-02 15:04:05`),
		UpdatedAt: time.Now().Format(`2006-01-02 15:04:05`),
	}

	ctx.JSON(http.StatusOK, response.Success(todo))
}

type TodoForm struct {
	Title string `form:"title" validate:"required,min=2"`
}
// TotoStore
// @Summary 待办事项添加
// @Description 用于添加待办事项
// @Tags Todo
// @Accept json
// @Produce json
// @Param body body handlers.TodoForm true "Todo FormData"
// @Success 200 {object} response.Response{code=int,data=models.Todo} "desc"
// @Router /todos [post]
func TodoStore(ctx *gin.Context) {
	var form TodoForm
	if err := ctx.BindJSON(&form); err != nil {
		ctx.JSON(http.StatusOK, response.ValidationRes(err.Error()))
		return
	}

	todo := models.Todo{
		ID:        1,
		Title:     form.Title,
		Status:    0,
		CreatedAt: time.Now().Format(`2006-01-02 15:04:05`),
		UpdatedAt: time.Now().Format(`2006-01-02 15:04:05`),
	}

	ctx.JSON(http.StatusOK, response.Success(todo))
}
// TodoUpdate
// @Summary 待办事项更新
// @Description 用于待办事项更新
// @Tags Todo
// @Accept json
// @Produce json
// @param id path int true "Todo id"
// @Param body body handlers.TodoForm true "Todo FormData"
// @Success 200 {object} response.Response{code=int,data=models.Todo} "desc"
// @Router /todos/{id} [put]
func TodoUpdate(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ValidationRes(err.Error()))
		return
	}

	var form TodoForm
	if err := ctx.BindJSON(&form); err != nil {
		ctx.JSON(http.StatusOK, response.ValidationRes(err.Error()))
		return
	}

	todo := models.Todo{
		ID:        id,
		Title:     form.Title,
		Status:    0,
		CreatedAt: time.Now().Format(`2006-01-02 15:04:05`),
		UpdatedAt: time.Now().Format(`2006-01-02 15:04:05`),
	}

	ctx.JSON(http.StatusOK, response.Success(todo))
}
// TodoDelete
// @Summary 待办事项删除
// @Description 用于待办事项删除
// @Tags Todo
// @Accept json
// @Produce json
// @param id path int true "Todo id"
// @Success 200 {object} response.Response{} "desc"
// @Router /todos/{id} [delete]
func TodoDestroy(ctx *gin.Context)  {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ValidationRes(err.Error()))
		return
	}

	fmt.Printf("%d 已删除\n", id)

	ctx.JSON(http.StatusOK, response.Success(nil))
}