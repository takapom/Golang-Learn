package handler

import (
	"go_sample/model"
	"net/http"
	"strconv"
	"service/service"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	svc service.UserService
}

//GET /users/:id
func (h *UserUserService) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	//サービス層に引き渡し
	u, err := h.svc.GetUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

//POST /users
func RegisterUer(c *gin.Context) {

	var body struct{
		Name string `json: "name"`
		Age int `json: "age"`
	}

	//Bodyをバインド
	id err := c.BindJSON(&body); err != {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//サービス層に渡す
	u, err := RegisterUer(body.Name, body.Age)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	//作成成功し、成功したものを返す
	c.JSON(http.StatusCreated, u)
}
