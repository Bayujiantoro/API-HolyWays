package handlers

import (
	dto "holyways/dto/result"
	userdto "holyways/dto/user"
	"holyways/models"
	"holyways/repository"
	"net/http"
	"strconv"

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
	users, err := h.UserRepo.FindUser()

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

func(h *userHandler) UpdateUser(c echo.Context) error {
	request := new(userdto.CreateUser)
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user, err := h.UserRepo.GetUser(id)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})}
	if request.Name != "" {
		user.Name = request.Email
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Password != "" {
		user.Password = request.Password
	}
	if request.Phone != "" {
		user.Phone = request.Phone
	}
	data, err := h.UserRepo.UpdateUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})}
	

	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponse(data),
	})
}

func (h *userHandler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	User, err := h.UserRepo.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponse(User),
	})
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	User , err := h.UserRepo.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	data , err := h.UserRepo.DeleteUser(User, id)
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
		Fund: user.Fund,
		Donation: user.Donation,
	}
}
