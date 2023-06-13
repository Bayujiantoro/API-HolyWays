package handlers

import (
	dto "holyways/dto/result"
	userdto "holyways/dto/user"
	"holyways/models"
	"holyways/repository"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	UserRepo repository.UserRepository
}

func HandlerUser(UserRepo repository.UserRepository) *userHandler {
	return &userHandler{UserRepo}
}

func (h *userHandler) FindUsers(c echo.Context) error {
	users, err := h.UserRepo.FindUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: users,
	})
}

func (h *userHandler) CreateUser(c echo.Context) error {
	request := new(userdto.CreateUser)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Phone:    request.Phone,
	}

	data, err := h.UserRepo.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponse(data),
	})
}

func convertResponse(user models.User) userdto.UserResponse {
	return userdto.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
	}
}
