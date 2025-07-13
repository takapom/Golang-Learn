// ここはサービス層に渡してあげるのが基本
package handler

import (
	"Todo_layered/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewTodoHandler(svc service.TodoService) *TodoHandler {
	return &TodoHandler{service: svc}
}

// サービス層のどの型定義を使うか
type TodoHandler struct {
	service service.TodoService
}

// ルーディング設定
func (h *TodoHandler) RegisterRoutes(r *gin.Engine) {
	todos := r.Group("/todos")
	todos.POST("", h.CreateTodo)
	todos.GET("", h.GetTodos)
	todos.GET("/:id", h.GetTodo)
	todos.PUT("/:id", h.UpdateTodo)
	todos.DELETE("/:id", h.DeleteTodo)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var input struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := h.service.CreateTodo(input.Title, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.service.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.service.GetTodo(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := h.service.UpdateTodo(uint(id), input.Title, input.Description, input.Completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteTodo(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
