package handlers

import (
	"fmt"
	funddto "holyways/dto/fund"
	dto "holyways/dto/result"
	"holyways/models"
	"holyways/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handlerFund struct {
	FundRepo repository.FundRepo
}

func FundHandler(Fund repository.FundRepo) *handlerFund {
	return &handlerFund{Fund}
}

func (h *handlerFund) CreateFund(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	
	GoalsMoney , _ := strconv.Atoi(c.FormValue("GoalsMoney"))
	UserID , _ := strconv.Atoi(c.FormValue("UserID"))

	
	request := models.Fund{
		Title: c.FormValue("Title"),
		Description: c.FormValue("Description"),
		GoalsMoney: GoalsMoney,
		GoalsDay: c.FormValue("GoalsDay"),
		Image: dataFile,
		UserID: UserID,
	}
	
	User, _ := h.FundRepo.GetUserById(request.UserID)
	
	fund := models.Fund {
		Title: request.Title,
		Description: request.Description,
		GoalsMoney: request.GoalsMoney,
		GoalsDay: request.GoalsDay,
		Image: dataFile,
		UserID: request.UserID,
		User: User,
	}
	fmt.Println("ini fund : ", fund)
	data, err := h.FundRepo.CreateFund(fund)
	

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseFund(data),
	})
}



func (h *handlerFund) UpdateFund(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	id, _ := strconv.Atoi(c.Param("id"))

	GoalsMoney , _ := strconv.Atoi(c.FormValue("GoalsMoney"))
	UserID , _ := strconv.Atoi(c.FormValue("UserID"))

	request := models.Fund{
		Title: c.FormValue("Title"),
		Description: c.FormValue("Description"),
		GoalsMoney: GoalsMoney,
		GoalsDay: c.FormValue("GoalsDay"),
		Image: dataFile,
		UserID: UserID,
	}

	fund, err := h.FundRepo.GetFund(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if request.Title != "" {
		fund.Title = request.Title
	} 
	if request.Description != "" {
		fund.Description = request.Description
	}
	if request.GoalsMoney > 0 {
		fund.GoalsMoney = request.GoalsMoney
	} 
	if request.GoalsDay != "" {
		fund.GoalsDay = request.GoalsDay
	}
	if request.Image != "" {
		fund.Image = request.Image
	} 
	if request.UserID > 0 {
		fund.UserID = request.UserID
	}

	data, err := h.FundRepo.UpdateFund(fund)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseFund(data),
	})
}

func (h *handlerFund) FindFund(c echo.Context) error {
	fund, err := h.FundRepo.FindFund()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: fund,
	})
}

func (h *handlerFund) GetFund(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fund, err := h.FundRepo.GetFund(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseFund(fund),
	})
}

func (h *handlerFund) DeleteFund(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fund , err := h.FundRepo.GetFund(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	data , err := h.FundRepo.DeleteFund(fund, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseFund(data),
	})
}

func convertResponseFund(fund models.Fund) funddto.FundResponse{
	return funddto.FundResponse{
		ID: fund.ID,
		Title: fund.Title,
		Description: fund.Description,
		GoalsMoney: fund.GoalsMoney,
		GoalsDay: fund.GoalsDay,
		Image: fund.Image,
		UserID: fund.UserID,
		User: fund.User,

	}
}