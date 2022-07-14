package controller

import (
	"net/http"
	"strconv"
	"usermsg/repository"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type UserController struct {
	repo repository.UserRepo
}

func NewUserController(repo repository.UserRepo) *UserController {
	return &UserController{
		repo: repo,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user repository.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	u.repo.Insert(&user)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":   user.Id,
		"name": user.Name,
	})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var user repository.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	u.repo.Update(id, user)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": user.Id,
	})
}

func (u *UserController) GetUserByName(c *gin.Context) {
	var name string = c.Param("name")

	user, err := u.repo.GetByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	u.repo.Delete(id)

	c.String(http.StatusOK, "employee deleted")
}
