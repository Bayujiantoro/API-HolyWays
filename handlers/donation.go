package handlers

import (
	"fmt"
	donationdto "holyways/dto/donation"
	dto "holyways/dto/result"
	"holyways/models"
	"holyways/repository"
	"net/http"
	"time"

	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerDonation struct {
	DonationRepo repository.DonationRepo
}

func DonationHandler(Donation repository.DonationRepo) *handlerDonation {
	return &handlerDonation{Donation}
}

func (h *handlerDonation) CreateDonation(c echo.Context) error  {
	Money, _ := strconv.Atoi(c.FormValue("Money"))
	FundID, _ := strconv.Atoi(c.FormValue("FundID"))


	userLogin := c.Get("userLogin")
	fmt.Println("userLogin : ", userLogin)
	userID := userLogin.(jwt.MapClaims)["id"].(float64)
	user, _ := h.DonationRepo.GetUserById(int(userID))

	Fund, _ := h.DonationRepo.GetFundById(FundID)
	fmt.Println(user)

	request := models.Donation{
		Date: time.Now(),
		Money: Money,
		FundID: FundID,
		UserID: int(userID),
		User: user,
		Fund: Fund,
	}
	fmt.Println("ini request : ", request)

	data, err := h.DonationRepo.CreateDonation(request)
	fmt.Println("ini data : ", data)
	

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseDonation(data),
	})
}

func (h *handlerDonation) UpdateDonation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Money, _ := strconv.Atoi(c.FormValue("Money"))
	request := models.Donation{
		Date: time.Now(),
		Money: Money,
	}
	fmt.Println("ini request : ", request)
	donation, err := h.DonationRepo.GetDonation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	
	if request.Money > 0 {
		donation.Date = request.Date
	}
	data, err := h.DonationRepo.UpdateDonation(donation)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseDonation(data),
	})
}

func (h *handlerDonation) FindDonation(c echo.Context) error {
	donation, err := h.DonationRepo.FindDonation() 

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: donation,
	})
}

func (h *handlerDonation) GetDonation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := h.DonationRepo.GetDonation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseDonation(data),
	})
}

func (h *handlerDonation) DeleteDonation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	donation, err := h.DonationRepo.GetDonation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	data, err := h.DonationRepo.DeleteDonation(donation, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseDonation(data),
	})
}


func convertResponseDonation(donation models.Donation) donationdto.DonationResponse{
	return donationdto.DonationResponse{
		ID: donation.ID,
		Date: donation.Date,
		Money: donation.Money,
		UserID: donation.UserID,
		User: donation.User,
		FundID: donation.FundID,
		Fund: donation.Fund,
	}
}